package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/tifye/hosts-file-editor-cli/pkg"
)

type Cli struct {
	HostsFile *pkg.HostsFile
}

var (
	cli     *Cli
	rootCmd *cobra.Command
)

func newRootCommand(cli *Cli) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "hfe",
		Short: "",
		Long:  ``,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			file, err := os.Open("C:\\windows\\system32\\drivers\\etc\\hosts")
			if err != nil {
				log.Fatalf("failed opening file: %s", err)
			}

			hf, err := pkg.ParseHostsFile(file)
			if err != nil {
				log.Fatalf("failed parsing file: %s", err)
			}

			cli.HostsFile = hf

			err = file.Close()
			if err != nil {
				log.Fatalf("Failed to close file %s", err)
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			newListCommand(cli).Execute()
		},
	}

	cmd.PersistentFlags().BoolP("help", "", false, "help for this command")

	return cmd
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cli = &Cli{}
	rootCmd = newRootCommand(cli)
	addCommands(rootCmd, cli)
}

func addCommands(cmd *cobra.Command, cli *Cli) {
	cmd.AddCommand(
		newAddCommand(cli),
		newListCommand(cli),
		newRemoveCommand(cli),
		newOpenCommand(),
	)
}
