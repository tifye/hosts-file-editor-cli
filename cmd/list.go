package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tifye/hosts-file-editor-cli/cmd/cli"
)

func NewListCommand(hostsCli cli.Cli) *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List all entries in the hosts file",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			cli.RenderHeader(hostsCli.HostsFile().Header)
			cli.RenderEntries(hostsCli.HostsFile().Entries)
		},
	}
}
