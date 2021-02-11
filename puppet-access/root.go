package cmd

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"

	"github.com/puppetlabs/pe-cli/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

//Main commmand
var (
	RootCmd = &cobra.Command{
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return validateGlobalFlags(cmd)
		},
	}
)

func init() {
	RootCmd.SetHelpCommand(&cobra.Command{
		Use:    "no-help",
		Hidden: true,
	})

	RootCmd.Flags().BoolP("help", "h", false, "Show this screen.")
	RootCmd.Flags().BoolP("version", "V", false, "Show version.")

	setCmdFlags(RootCmd)
	bindConfigFlags(RootCmd)

}

//Execute represents the function that executes the main command
func Execute(version string) error {
	RootCmd.Version = version
	return RootCmd.Execute()
}

func setCmdFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().BoolP(
		"verbose",
		"",
		false,
		"Set verbose output",
	)
	cmd.PersistentFlags().BoolP(
		"debug",
		"d",
		false,
		"Enable debug logging.",
	)
	cmd.PersistentFlags().StringP(
		"token-file",
		"t",
		getDefaultTokenFile(),
		"Location of the token file.",
	)
	cmd.PersistentFlags().StringP(
		"ca-cert",
		"",
		getDefaultCacert(),
		"CA cert to use to contact token-issuing service.",
	)
	cmd.PersistentFlags().StringP(
		"service-url",
		"",
		"",
		"FQDN, port, and API prefix of server where the token issuing service/server can be contacted \n(the Puppet Enterprise console node).(example: https://<HOSTNAME>:4433/rbac-api)",
	)
	cmd.PersistentFlags().StringP(
		"config-file",
		"c",
		getDefaultConfigFile(),
		" Path to configuration file.",
	)

}
func initConfig(cfgFile string) {
	viper.SetConfigType("json")
	err := readConfigFile(cfgFile)
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
}

func readConfigFile(cfgFile string) error {
	err := readGlobalConfigFile()
	if err != nil {
		return err
	}

	return mergeUserConfigFile(cfgFile)
}

func getGlobalConfigFile() (string, error) {
	puppetLabsDir, err := PuppetLabsDir()
	if err != nil {
		return puppetLabsDir, err
	}

	globalConfigFile := filepath.Join(puppetLabsDir, "client-tools", "puppet-access.conf")
	return globalConfigFile, nil
}

func readGlobalConfigFile() error {
	globalConfigFile, err := getGlobalConfigFile()
	if err != nil {
		return err
	}

	_, err = os.Stat(globalConfigFile)
	if err != nil {
		log.Debug(fmt.Sprintf("Failed reading global config file: %s", err.Error()))
		return nil
	}
	viper.SetConfigFile(globalConfigFile)
	return viper.ReadInConfig()
}

func getDefaultConfigFile() string {
	usr, err := user.Current()
	if err != nil {
		log.Error(err.Error())
		return ""
	}

	configFile := filepath.Join(usr.HomeDir, ".puppetlabs", "client-tools", "puppet-access.conf")
	return configFile
}

func mergeUserConfigFile(cfgFile string) error {
	_, err := os.Stat(cfgFile)
	if err != nil {
		if cfgFile == getDefaultConfigFile() {
			log.Debug(fmt.Sprintf("Failed reading default config file: %s", err.Error()))
			return nil
		}
		log.Error(fmt.Sprintf("Failed reading CLI config file: %s", err.Error()))
		return err
	}
	viper.SetConfigFile(cfgFile)
	return viper.MergeInConfig()
}

func validateGlobalFlags(cmd *cobra.Command) error {
	cfgFile, err := cmd.Flags().GetString("config-file")
	if err != nil {
		return err
	}
	initConfig(cfgFile)

	return nil
}

func bindConfigFlags(cmd *cobra.Command) {
	viper.BindPFlag("service-url", cmd.PersistentFlags().Lookup("service-url"))
	viper.BindPFlag("certificate-file", cmd.PersistentFlags().Lookup("ca-cert"))
	viper.BindPFlag("token-file", cmd.PersistentFlags().Lookup("token-file"))
}

func getDefaultCacert() string {
	puppetLabsDir, err := PuppetLabsDir()
	if err != nil {
		log.Error(err.Error())
		return ""
	}

	return filepath.Join(puppetLabsDir, "puppet", "ssl", "certs", "ca.pem")
}

func getDefaultTokenFile() string {
	usr, err := user.Current()
	if err != nil {
		log.Error(err.Error())
		return ""
	}

	return filepath.Join(usr.HomeDir, ".puppetlabs", "token")
}
