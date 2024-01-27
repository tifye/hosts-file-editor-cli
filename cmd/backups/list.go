package backups

import (
	"fmt"
	"log"
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
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

			renderBackupsList(backups)
		},
	}
	return cmd
}

func renderBackupsList(backups []pkg.Backup) {
	re := lipgloss.NewRenderer(os.Stdout)

	var (
		headerStyle = re.NewStyle().
				Foreground(lipgloss.Color("#f43f5e")).
				Bold(true).
				Align(lipgloss.Center).
				Padding(0, 1)
		cellStyle    = re.NewStyle().Padding(0, 1)
		oddRowStyle  = cellStyle.Copy().Foreground(lipgloss.Color("#d1d5db"))
		evenRowStyle = cellStyle.Copy().Foreground(lipgloss.Color("#9ca3af"))
		borderStyle  = re.NewStyle().Foreground(lipgloss.Color("#f43f5e"))
	)

	t := table.New().
		Headers("NR", "TIME", "COMMENT").
		Border(lipgloss.ThickBorder()).
		BorderStyle(borderStyle).
		StyleFunc(func(row, col int) lipgloss.Style {
			switch {
			case row == 0:
				return headerStyle
			case row%2 == 0:
				return evenRowStyle
			default:
				return oddRowStyle
			}
		})

	for i, backup := range backups {
		t.Row(fmt.Sprint(i), fmt.Sprintf("%s", backup.Time), backup.Comment)
	}

	fmt.Println(t)
}
