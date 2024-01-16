package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/tifye/hosts-file-editor-cli/core"
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
			entry := &core.HostEntry{
				Hostname: hostname,
				IP:       ip,
				Comment:  comment,
			}

			file, err := os.OpenFile("C:\\windows\\system32\\drivers\\etc\\hosts", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
				log.Fatalf("Failed to open hosts file %s", err)
			}
			defer file.Close()

			cli.Editor.Writer = file

			err = cli.Editor.AddEntry(*entry)
			if err != nil {
				log.Fatalf("Failed to add entry %s", err)
			}
		},
	}

	cmd.Flags().StringVarP(&hostname, "hostname", "n", "", "Hostname")
	cmd.MarkFlagRequired("hostname")

	cmd.Flags().StringVar(&ip, "ip", "", "IP")
	cmd.MarkFlagRequired("ip")

	cmd.Flags().StringVarP(&comment, "comment", "c", "", "Comment")

	return cmd
}
