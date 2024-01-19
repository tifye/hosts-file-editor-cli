package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/tifye/hosts-file-editor-cli/pkg"
)

type addOptions struct {
	hostname string
	ip       string
	comment  string
}

func newAddCommand(cli *Cli) *cobra.Command {
	opts := &addOptions{}

	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add a new entry to the hosts file",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			entry := &pkg.HostEntry{
				Hostname: opts.hostname,
				IP:       opts.ip,
				Comment:  opts.comment,
			}

			cli.HostsFile.AddEntry(*entry)

			err := pkg.SaveToFile(cli.HostsFile, "C:\\windows\\system32\\drivers\\etc\\hosts")
			if err != nil {
				log.Fatal(err)
			}

			newListCommand(cli).Run(cmd, args)
		},
	}

	cmd.Flags().StringVarP(&opts.hostname, "hostname", "h", "", "Hostname")
	cmd.MarkFlagRequired("hostname")

	cmd.Flags().StringVar(&opts.ip, "ip", "", "IP")
	cmd.MarkFlagRequired("ip")

	cmd.Flags().StringVarP(&opts.comment, "comment", "c", "", "Comment")

	return cmd
}
