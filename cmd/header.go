package cmd

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

func newHeaderCommand(cli *Cli) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "header",
		Short: "Print the header comments (comments at top of hosts file)",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			var sb strings.Builder
			for _, line := range cli.HostsFile.Header {
				sb.WriteString(strings.TrimPrefix(line, "#") + "\n")
			}

			meep := lipgloss.NewStyle().
				BorderStyle(lipgloss.RoundedBorder()).
				BorderForeground(lipgloss.Color("#f43f5e")).
				Padding(1, 2).
				Align(lipgloss.Left).
				Render(sb.String())
			fmt.Println(meep)
		},
	}
	return cmd
}
