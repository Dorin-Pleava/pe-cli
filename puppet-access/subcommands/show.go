package main

import (
	"fmt"
	"os"

	"github.com/puppetlabs/pe-cli/log"
	cmd "github.com/puppetlabs/pe-cli/puppet-access"
	"github.com/puppetlabs/pe-sdk-go/token/filetoken"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Prints the saved token contents to stdout",
	Run:   executeShowCommand,
}

func init() {
	cmd.RootCmd.AddCommand(showCmd)
}

func executeShowCommand(cmd *cobra.Command, args []string) {
	fileToken := filetoken.NewFileToken(viper.GetString("token-file"))
	tokenValue, err := fileToken.Read()
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
	fmt.Println(tokenValue)

}
