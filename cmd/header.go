package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tifye/hosts-file-editor-cli/cmd/cli"
)

func NewHeaderCommand(hostsCli cli.Cli) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "header",
		Short: "Print the header comments",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			cli.RenderHeader(hostsCli.HostsFile().Header)
		},
	}
	return cmd
}
