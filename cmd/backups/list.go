package backups

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tifye/hosts-file-editor-cli/cmd/cli"
)

func newListCommand(hostsCli cli.Cli) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all backups",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("backups/list called")
		},
	}
	return cmd
}
