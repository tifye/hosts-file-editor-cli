package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/tifye/hosts-file-editor-cli/cmd/cli"
	"github.com/tifye/hosts-file-editor-cli/pkg"
)

type addOptions struct {
	hostname string
	ip       string
	comment  string
}

func NewAddCommand(hostsCli cli.Cli) *cobra.Command {
	opts := &addOptions{}

	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add a new entry to the hosts file",
		Long:  ``,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return pkg.CreateBackupFile(hostsCli.HostsFile(), "add")
		},
		Run: func(cmd *cobra.Command, args []string) {
			entry := &pkg.HostEntry{
				Hostname: opts.hostname,
				IP:       opts.ip,
				Comment:  opts.comment,
			}
			hostsCli.HostsFile().AddEntry(*entry)

			err := pkg.SaveToFile(hostsCli.HostsFile(), "C:\\windows\\system32\\drivers\\etc\\hosts")
			if err != nil {
				log.Fatal(err)
			}

			NewListCommand(hostsCli).Run(cmd, args)
		},
	}

	cmd.Flags().StringVarP(&opts.hostname, "hostname", "h", "", "Hostname")
	cmd.MarkFlagRequired("hostname")

	cmd.Flags().StringVar(&opts.ip, "ip", "", "IP")
	cmd.MarkFlagRequired("ip")

	cmd.Flags().StringVarP(&opts.comment, "comment", "c", "", "Comment")

	return cmd
}
