package cmd

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
	"github.com/tifye/hosts-file-editor-cli/cmd/cli"
)

func NewHeaderCommand(hostsCli cli.Cli) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "header",
		Short: "Print the header comments (comments at top of hosts file)",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			renderHeader(hostsCli.HostsFile().Header)
		},
	}
	return cmd
}

func renderHeader(header []string) {
	var sb strings.Builder
	for _, line := range header {
		sb.WriteString(strings.TrimPrefix(line, "#") + "\n")
	}

	headerBox := lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#f43f5e")).
		Padding(0, 2).
		Align(lipgloss.Left).
		Render(sb.String())
	fmt.Println(headerBox)
}
