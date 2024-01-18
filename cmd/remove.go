package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

type removeOptions struct {
	hostname string
	ip       string
}

func newRemoveCommand(cli *Cli) *cobra.Command {
	opts := &removeOptions{}

	cmd := &cobra.Command{
		Use:   "remove",
		Short: "A brief description of your command",
		Long:  `Remove an entry from the hosts file`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Not implemented")
		},
	}

	cmd.Flags().StringVarP(&opts.hostname, "hostname", "n", "", "Hostname")
	cmd.Flags().StringVar(&opts.ip, "ip", "", "IP")
	cmd.MarkFlagsOneRequired("ip", "hostname")

	return cmd
}
