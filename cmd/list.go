package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/spf13/cobra"
	"github.com/tifye/hosts-file-editor-cli/core"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all entries in the hosts file",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		file, err := os.Open("C:\\windows\\system32\\drivers\\etc\\hosts")
		if err != nil {
			log.Fatalf("failed opening file: %s", err)
		}
		defer file.Close()

		entries, err := core.ParseHostsFile(file)
		if err != nil {
			log.Fatalf("failed parsing file: %s", err)
		}

		if len(entries) == 0 {
			log.Println("No entries found")
			return
		}

		renderList(entries)
	},
}

func renderList(entries []core.HostEntry) {
	re := lipgloss.NewRenderer(os.Stdout)

	var (
		HeaderStyle  = re.NewStyle().Foreground(lipgloss.Color("#f43f5e")).Bold(true).Align(lipgloss.Center).Padding(0, 1)
		CellStyle    = re.NewStyle().Padding(0, 1)
		OddRowStyle  = CellStyle.Copy().Foreground(lipgloss.Color("#9ca3af"))
		EvenRowStyle = CellStyle.Copy().Foreground(lipgloss.Color("#d1d5db"))
		BorderStyle  = re.NewStyle().Foreground(lipgloss.Color("#f43f5e"))
	)

	t := table.New().
		Headers("HOSTNAME", "IP", "COMMENTS").
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

	for _, entry := range entries {
		t.Row(entry.Hostname, entry.IP, strings.TrimSpace(entry.Comment))
	}

	fmt.Println(t)
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
