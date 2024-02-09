package backups

import (
	"github.com/spf13/cobra"
	"github.com/tifye/hosts-file-editor-cli/cmd/cli"
)

func NewBackupsCommand(hostsCli cli.Cli) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "backups",
		Short: "Subcommands for managing backups",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	cmd.AddCommand(
		newListCommand(hostsCli),
		newOpenCommand(hostsCli),
		newRestoreCommand(hostsCli),
		newClearCommand(hostsCli),
	)

	return cmd
}
