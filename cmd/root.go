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

	addCommands(rootCmd, cli)
}

func addCommands(cmd *cobra.Command, cli *Cli) {
	cmd.AddCommand(
		newAddCommand(cli),
		newListCommand(),
		newRemoveCommand(cli),
		newOpenCommand(),
	)
}
