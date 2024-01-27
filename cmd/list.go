package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/spf13/cobra"
	"github.com/tifye/hosts-file-editor-cli/cmd/cli"
	"github.com/tifye/hosts-file-editor-cli/pkg"
)

func NewListCommand(hostsCli cli.Cli) *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List all entries in the hosts file",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			renderHeader(hostsCli.HostsFile().Header)
			renderList(hostsCli.HostsFile().Entries)
		},
	}
}

func renderList(entries []pkg.HostEntry) {
	re := lipgloss.NewRenderer(os.Stdout)

	var (
		HeaderStyle  = re.NewStyle().Foreground(lipgloss.Color("#f43f5e")).Bold(true).Align(lipgloss.Center).Padding(0, 1)
		CellStyle    = re.NewStyle().Padding(0, 1)
		OddRowStyle  = CellStyle.Copy().Foreground(lipgloss.Color("#9ca3af"))
		EvenRowStyle = CellStyle.Copy().Foreground(lipgloss.Color("#d1d5db"))
		BorderStyle  = re.NewStyle().Foreground(lipgloss.Color("#f43f5e"))
	)

	t := table.New().
		Headers("NR", "HOSTNAME", "IP", "COMMENTS").
		Border(lipgloss.ThickBorder()).
		BorderStyle(BorderStyle).
		StyleFunc(func(row, col int) lipgloss.Style {
			switch {
			case row == 0:
				return HeaderStyle
			case row%2 == 0:
				return EvenRowStyle
			default:
				return OddRowStyle
			}
		})

	for i, entry := range entries {
		t.Row(fmt.Sprint(i), entry.Hostname, entry.IP, strings.TrimSpace(entry.Comment))
	}

	fmt.Println(t)
}
