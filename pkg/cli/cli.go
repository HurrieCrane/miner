package cli

import "fmt"

const help = `miner - commandline minecraft server deployment tool

Usage: miner [options] [plan]

miner is a tool that allows you to deploy minecraft servers on
a different platforms with a single platform agnostic plan and
some provider specific configuration.

Some of the options include:
	-help 		A command that displays help information
	-providers 	Lists all of the currently avalible providers

To report issues visit https://github.com
`

const providers = `Some of the avalible providers are:
	terraform-aws 		Deploys to Amazon Web Services using the Terraform tool
`

// PrintHelp prints the help text
func PrintHelp() {
	fmt.Print(help)
}

// PrintProviderList prints the
// complete list of avalible providers
func PrintProviderList() {
	fmt.Print(providers)
}
