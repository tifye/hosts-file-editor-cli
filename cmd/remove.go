package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/tifye/hosts-file-editor-cli/pkg"
)

type removeOptions struct {
	hostname   string
	ip         string
	duplicates bool
}

func newRemoveCommand(cli *Cli) *cobra.Command {
	opts := &removeOptions{}

	cmd := &cobra.Command{
		Use:   "remove",
		Short: "A brief description of your command",
		Long:  `Remove an entry from the hosts file`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return pkg.CreateBackupFile(cli.HostsFile, "remove")
		},
		Run: func(cmd *cobra.Command, args []string) {
			var filtered []pkg.HostEntry

			if opts.duplicates {
				filtered = pkg.FilterOutDuplicates(cli.HostsFile.Entries, opts.hostname, opts.ip)
				for _, meep := range filtered {
					fmt.Println(meep.String())
				}
			} else {
				filtered = pkg.FilterOut(cli.HostsFile.Entries, opts.hostname, opts.ip)
			}

			cli.HostsFile.Entries = filtered

			err := pkg.SaveToFile(cli.HostsFile, "C:\\windows\\system32\\drivers\\etc\\hosts")
			if err != nil {
				log.Fatal(err)
			}

			newListCommand(cli).Run(cmd, args)
		},
	}

	cmd.Flags().BoolVarP(&opts.duplicates, "duplicates", "d", false, "remove duplicates only")

	cmd.Flags().StringVarP(&opts.hostname, "hostname", "h", "", "remove entries with hostname")
	cmd.Flags().StringVar(&opts.ip, "ip", "", "remove entries with ip")
	return cmd
}
