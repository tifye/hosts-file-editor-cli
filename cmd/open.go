package cmd

import (
	"log"
	"os/exec"

	"github.com/spf13/cobra"
)

var openCmd = &cobra.Command{
	Use:   "open",
	Short: "Opens hosts file in default text editor for viewing only",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		execCmd := exec.Command("notepad", "C:\\windows\\system32\\drivers\\etc\\hosts")
		err := execCmd.Start()
		if err, ok := err.(*exec.ExitError); ok {
			log.Fatalf("Failed to open hosts file in editor %s", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(openCmd)
}
