package backups

import (
	"fmt"
	"log"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
	"github.com/tifye/hosts-file-editor-cli/cmd/cli"
	"github.com/tifye/hosts-file-editor-cli/pkg"
)

func newRestoreCommand(hostsCli cli.Cli) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "restore",
		Short: "Restore from a backup",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			backups, err := pkg.GetListOfBackups()
			if err != nil {
				log.Fatalln(err)
			}

			var selectOpts []huh.Option[*pkg.Backup]
			for _, backup := range backups {
				key := fmt.Sprintf("%s %s", backup.Time, backup.Comment)
				selectOpts = append(selectOpts, huh.NewOption(key, &backup))
			}

			var selected *pkg.Backup
			huh.NewSelect[*pkg.Backup]().
				Title("Select a backup").
				Options(selectOpts...).
				Value(&selected).
				Run()

			if selected == nil {
				return
			}

			fmt.Println(selected.Filepath)
		},
	}

	return cmd
}
