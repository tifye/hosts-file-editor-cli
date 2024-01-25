package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/tifye/hosts-file-editor-cli/cmd/backups"
	"github.com/tifye/hosts-file-editor-cli/cmd/cli"
	"github.com/tifye/hosts-file-editor-cli/pkg"
)

var (
	hostsCli cli.Cli
	rootCmd  *cobra.Command
)

func newRootCommand(hostsCli cli.Cli) *cobra.Command {
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
			if err != nil || hf == nil {
				log.Fatalf("failed parsing file: %s", err)
			}

			hostsCli.SetHostsFile(hf)

			err = file.Close()
			if err != nil {
				log.Fatalf("Failed to close file %s", err)
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			NewListCommand(hostsCli).Execute()
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
	hostsCli = &cli.HostsCli{
		Accessible: os.Getenv("ACCESSIBLE") != "", // Todo: add to config cmd later
	}
	rootCmd = newRootCommand(hostsCli)
	addCommands(rootCmd, hostsCli)
}

func addCommands(cmd *cobra.Command, hostsCli cli.Cli) {
	cmd.AddCommand(
		NewAddCommand(hostsCli),
		NewListCommand(hostsCli),
		NewRemoveCommand(hostsCli),
		NewOpenCommand(),
		NewHeaderCommand(hostsCli),
		backups.NewBackupsCommand(hostsCli),
	)
}
