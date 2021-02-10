package main

import (
	"os"

	"github.com/puppetlabs/pe-cli/log"
	cmd "github.com/puppetlabs/pe-cli/puppet-access"
)

//Version of command
var Version = "0.0.0"

func init() {
	cmd.RootCmd.Use = "puppet-access [command] [flags]"
	cmd.RootCmd.Short = "puppet-access"
}

func main() {
	if err := cmd.Execute(Version); err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
}
