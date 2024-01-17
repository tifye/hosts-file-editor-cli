package cmd

import (
	"fmt"
	"log"
	"os"

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
			file, err := os.OpenFile("C:\\windows\\system32\\drivers\\etc\\hosts", os.O_RDWR|os.O_TRUNC, 0666)
			if err != nil {
				log.Fatalf("Failed to open hosts file %s", err)
			}
			offset, err := file.Seek(0, 0)
			if err != nil {
				log.Fatalf("Failed to touch hosts file")
			}
			fmt.Printf("current offset %v", offset)
			defer file.Close()

			cli.Editor.Writer = file
			cli.Editor.Reader = file

			filtered, err := cli.Editor.FilterOutEntries(opts.hostname, opts.ip)
			if err != nil {
				log.Fatalf("Failed to remove entry %s", err)
			}

			file.Seek(0, 0)
			cli.Editor.ReplaceWith(filtered)
		},
	}

	cmd.Flags().StringVarP(&opts.hostname, "hostname", "n", "", "Hostname")
	cmd.Flags().StringVar(&opts.ip, "ip", "", "IP")
	cmd.MarkFlagsOneRequired("ip", "hostname")

	return cmd
}
