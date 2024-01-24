package cmd

import (
	"fmt"
	"log"

	"github.com/charmbracelet/huh"
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
			defer newListCommand(cli).Run(cmd, args)

			var filtered []pkg.HostEntry
			if opts.duplicates {
				filtered = pkg.FilterOutDuplicates(cli.HostsFile.Entries, opts.hostname, opts.ip)
				for _, meep := range filtered {
					fmt.Println(meep.String())
				}
			} else {
				filtered = pkg.FilterOut(cli.HostsFile.Entries, opts.hostname, opts.ip)
			}

			numMatchedEntries := len(cli.HostsFile.Entries) - len(filtered)
			if numMatchedEntries <= 0 {
				log.Println("No matching entries found")
				return
			}

			var didConfirm bool
			err := huh.NewConfirm().
				Title(fmt.Sprintf("Matched %d entries, are you sure you want to remove?", numMatchedEntries)).
				Value(&didConfirm).
				WithAccessible(cli.AccessibleMode).
				Run()
			if err != nil {
				log.Fatal(err)
			}

			if !didConfirm {
				return
			}

			cli.HostsFile.Entries = filtered

			err = pkg.SaveToFile(cli.HostsFile, "C:\\windows\\system32\\drivers\\etc\\hosts")
			if err != nil {
				log.Fatal(err)
			}
		},
	}

	cmd.Flags().BoolVarP(&opts.duplicates, "duplicates", "d", false, "remove duplicates only")

	cmd.Flags().StringVarP(&opts.hostname, "hostname", "h", "", "remove entries with hostname")
	cmd.Flags().StringVar(&opts.ip, "ip", "", "remove entries with ip")
	return cmd
}
