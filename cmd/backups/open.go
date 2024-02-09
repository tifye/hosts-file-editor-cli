package backups

import (
	"log"
	"os/exec"

	"github.com/spf13/cobra"
	"github.com/tifye/hosts-file-editor-cli/cmd/cli"
	"github.com/tifye/hosts-file-editor-cli/pkg"
)

func newOpenCommand(hostsCli cli.Cli) *cobra.Command {
	cmd := &cobra.Command{
		Use:              "open",
		Short:            "Open folder containing backup files",
		Long:             ``,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {},
		Run: func(cmd *cobra.Command, args []string) {
			backupsDir, err := pkg.GetBackupsDirPath()
			if err != nil {
				log.Fatalln(err)
			}

			execCmd := exec.Command("explorer", backupsDir)
			err = execCmd.Start()
			if err, ok := err.(*exec.ExitError); ok {
				log.Fatalf("Failed to open hosts file in editor %s", err)
			}
		},
	}
	return cmd
}
