package main

import (
	"os"

	"github.com/puppetlabs/pe-cli/log"
	cmd "github.com/puppetlabs/pe-cli/puppet-access"

	"github.com/puppetlabs/pe-sdk-go/token/filetoken"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var force bool

var deleteCmd = &cobra.Command{
	Use:   "delete-token-file",
	Short: "Removes the local token, but does not expire the token on the server",
	Run:   executeDeleteCommand,
}

func init() {
	cmd.RootCmd.AddCommand(deleteCmd)
}

func executeDeleteCommand(cmd *cobra.Command, args []string) {
	//FIXME add --force flag
	force := false
	fileToken := filetoken.NewFileToken(viper.GetString("token-file"))
	err := fileToken.Delete(force)
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
}
