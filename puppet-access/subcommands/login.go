package main

import (
	"fmt"
	"os"
	"regexp"

	cmd "github.com/puppetlabs/pe-cli/puppet-access"
	app "github.com/puppetlabs/pe-sdk-go/app/puppet-access"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	print bool
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to generate a token in RBAC",
	Run:   executeLoginCommand,
}

func init() {
	loginCmd.PersistentFlags().StringP(
		"username",
		"",
		"",
		"The username to login (optional)",
	)
	loginCmd.PersistentFlags().StringP(
		"lifetime",
		"l",
		"1h",
		`Specify the duration for which a token will be
valid. The default duration is set by the Puppet
Enterprise administrator. Accepts a value in the
format <integer>[smhdy].`,
	)
	loginCmd.Flags().BoolVarP(
		&print,
		"print",
		"",
		false,
		`Generates a token, but does not save it to disk.
Instead, the token contents are displayed on
STDOUT.`,
	)

	cmd.RootCmd.AddCommand(loginCmd)
}

func executeLoginCommand(cmd *cobra.Command, args []string) {

	fmt.Println("Enter your Puppet Enterprise credentials.")

	username, _ := cmd.PersistentFlags().GetString("username")
	lifetime, _ := cmd.PersistentFlags().GetString("lifetime")

	tries := 3
	for requestTokenWithInput(username, lifetime) != nil {
		tries = tries - 1
		if tries == 0 {
			fmt.Println("Failed to log in. Please try again.")
			os.Exit(1)
		}
		fmt.Println("ERROR: Sorry. We aren't able to log you in.")
		fmt.Println("Please try again.")
	}
}

func requestTokenWithInput(username, lifetime string) error {

	var usernameIn string
	var passwordIn string

	if username == "" {

		fmt.Print("Username: ")
		fmt.Scanln(&usernameIn)
		username = usernameIn
	}

	tries := 3
	for username == "" {
		tries = tries - 1
		if tries < 0 {
			fmt.Println("No username supplied.")
			os.Exit(1)
		}
		fmt.Println("Please enter a username.")
		fmt.Print("Username: ")
		fmt.Scanln(&usernameIn)
		username = usernameIn
	}

	fmt.Print("Password: ")
	fmt.Scanln(&passwordIn)

	validateLifetime(lifetime)
	puppetAccess := app.NewWithMinimalConfig(username, passwordIn, viper.GetString("service-url"), viper.GetString("certificate-file"), viper.GetString("token-file"))

	err := puppetAccess.Login(print)
	return err
}

func validateLifetime(lifetime string) (bool, error) {
	lifetimeExpr := `^\d+[smhdy]$`

	validLifetime, err := regexp.MatchString(lifetimeExpr, lifetime)
	if err != nil {
		return false, err
	}
	return validLifetime, nil
}
