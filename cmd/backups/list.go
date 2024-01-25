package backups

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/tifye/hosts-file-editor-cli/cmd/cli"
	"github.com/tifye/hosts-file-editor-cli/pkg"
)

func newListCommand(hostsCli cli.Cli) *cobra.Command {
	cmd := &cobra.Command{
		Use:              "list",
		Short:            "List all backups",
		Long:             ``,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {},
		Run: func(cmd *cobra.Command, args []string) {
			backups, err := pkg.GetListOfBackups()
			if err != nil {
				log.Fatalln(err)
			}

			for _, backup := range backups {
				fmt.Printf("%s %s\n", backup.Time, backup.Comment)
			}
		},
	}
	return cmd
}
