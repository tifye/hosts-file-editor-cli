package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/tifye/hosts-file-editor-cli/pkg"
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
			filtered := pkg.FilterOut(cli.HostsFile.Entries, opts.hostname, opts.ip)
			cli.HostsFile.Entries = filtered

			err := pkg.SaveToFile(cli.HostsFile, "C:\\windows\\system32\\drivers\\etc\\hosts")
			if err != nil {
				log.Fatal(err)
			}

			newListCommand(cli).Run(cmd, args)
		},
	}

	cmd.Flags().StringVarP(&opts.hostname, "hostname", "n", "", "Hostname")
	cmd.Flags().StringVar(&opts.ip, "ip", "", "IP")
	cmd.MarkFlagsOneRequired("ip", "hostname")

	return cmd
}
