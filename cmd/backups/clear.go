package backups

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
	"github.com/tifye/hosts-file-editor-cli/cmd/cli"
	"github.com/tifye/hosts-file-editor-cli/pkg"
)

func newClearCommand(hostsCli cli.Cli) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "clear",
		Short: "Clear all backup files. User will be prompted for confirmation",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			var didConfirm bool
			err := huh.NewConfirm().
				Title("Are you sure you want to clear all backups? This action cannot be backed up").
				Value(&didConfirm).
				WithAccessible(hostsCli.AccessibleMode()).
				Run()
			if err != nil {
				log.Fatalln(err)
			}

			if !didConfirm {
				return
			}

			backupsDir, err := pkg.GetBackupsDirPath()
			if err != nil {
				log.Fatalln(err)
			}
			backups, err := pkg.GetListOfBackups()
			if err != nil {
				log.Fatalln(err)
			}

			for _, backup := range backups {
				err := os.Remove(filepath.Join(backupsDir, backup.Filepath))
				if err != nil {
					fmt.Printf("Failed to remove backup file %s", backup.Filepath)
				}
			}

			backups, err = pkg.GetListOfBackups()
			if err != nil {
				log.Fatalln(err)
			}

			renderBackupsList(backups)
		},
	}

	return cmd
}
