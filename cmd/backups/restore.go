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

			var selectOpts []huh.Option[pkg.Backup]
			for _, backup := range backups {
				key := fmt.Sprintf("%s %s", backup.Time, backup.Comment)
				selectOpts = append(selectOpts, huh.NewOption(key, backup))
			}

			var selected pkg.Backup
			huh.NewSelect[pkg.Backup]().
				Title("Select a backup").
				Options(selectOpts...).
				Value(&selected).
				Run()

			backupsDir, err := pkg.GetBackupsDirPath()
			if err != nil {
				log.Fatalln(err)
			}

			file, err := os.OpenFile(filepath.Join(backupsDir, selected.Filepath), os.O_RDONLY, 0666)
			if err != nil {
				log.Fatalln(err)
			}

			backupHf, err := pkg.ParseHostsFile(file)
			if err != nil {
				log.Fatalln(err)
			}

			_ = pkg.CreateBackupFile(hostsCli.HostsFile(), "backup")

			err = pkg.SaveToFile(backupHf, "C:\\windows\\system32\\drivers\\etc\\hosts")
			if err != nil {
				log.Fatal(err)
			}

			cli.RenderHeader(backupHf.Header)
			fmt.Println(cli.RenderEntries(backupHf.Entries))
		},
	}

	return cmd
}
