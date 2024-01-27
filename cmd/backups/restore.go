package backups

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tifye/hosts-file-editor-cli/cmd/cli"
)

func newRestoreCommand(hostsCli cli.Cli) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "restore",
		Short: "Restore from a backup",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("backups/restore called")
		},
	}

	return cmd
}
