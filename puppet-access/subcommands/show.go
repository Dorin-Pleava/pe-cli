package main

import (
	cmd "github.com/puppetlabs/pe-cli/puppet-access"
	"github.com/spf13/cobra"
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

}
