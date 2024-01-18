package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/tifye/hosts-file-editor-cli/pkg"
)

var (
	hostname string
	ip       string
	comment  string
)

func newAddCommand(cli *Cli) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add a new entry to the hosts file",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			entry := &pkg.HostEntry{
				Hostname: hostname,
				IP:       ip,
				Comment:  comment,
			}

			cli.HostsFile.AddEntry(*entry)

			err := pkg.SaveToFile(cli.HostsFile, "C:\\windows\\system32\\drivers\\etc\\hosts")
			if err != nil {
				log.Fatal(err)
			}

			newListCommand(cli).Run(cmd, args)
		},
	}

	cmd.Flags().StringVarP(&hostname, "hostname", "n", "", "Hostname")
	cmd.MarkFlagRequired("hostname")

	cmd.Flags().StringVar(&ip, "ip", "", "IP")
	cmd.MarkFlagRequired("ip")

	cmd.Flags().StringVarP(&comment, "comment", "c", "", "Comment")

	return cmd
}
