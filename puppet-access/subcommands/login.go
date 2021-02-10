package main

import (
	cmd "github.com/puppetlabs/pe-cli/puppet-access"
	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to generate a token in RBAC",
	Run:   executeLoginCommand,
}

func init() {
	cmd.RootCmd.AddCommand(loginCmd)
}

func executeLoginCommand(cmd *cobra.Command, args []string) {

}
