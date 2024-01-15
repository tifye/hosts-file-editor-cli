package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var cli *Cli

var rootCmd = &cobra.Command{
	Use:   "hfe",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		newListCommand().Execute()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cli = NewCli()

	rootCmd.AddCommand(newAddCommand(cli))
	rootCmd.AddCommand(newListCommand())
}
