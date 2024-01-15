package cmd

import (
	"fmt"
	"log"

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
			fmt.Printf("%s %s %s", hostname, ip, comment)
			entry := &core.HostEntry{
				Hostname: hostname,
				IP:       ip,
				Comment:  comment,
			}

			err := cli.Editor.AddEntry(*entry)
			if err != nil {
				log.Fatalf("Failed to add entry %s", err)
			}

			fmt.Println("Entry added")
		},
	}

	cmd.Flags().StringVarP(&hostname, "hostname", "n", "", "Hostname")
	cmd.MarkFlagRequired("hostname")

	cmd.Flags().StringVar(&ip, "ip", "", "IP")
	cmd.MarkFlagRequired("ip")

	cmd.Flags().StringVarP(&comment, "comment", "c", "", "Comment")

	return cmd
}
