package cmd

import (
	"log"
	"os"

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

			file, err := os.OpenFile("C:\\windows\\system32\\drivers\\etc\\hosts", os.O_TRUNC|os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
				log.Fatalf("Failed to open hosts file for writing %s", err)
			}

			file.Seek(0, 0)
			cli.HostsFile.SaveTo(file)

			if err = file.Close(); err != nil {
				log.Printf("Failed to close hosts file %s", err)
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
