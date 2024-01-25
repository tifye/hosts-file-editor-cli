package cmd

import (
	"fmt"
	"log"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
	"github.com/tifye/hosts-file-editor-cli/cmd/cli"
	"github.com/tifye/hosts-file-editor-cli/pkg"
)

type removeOptions struct {
	hostname   string
	ip         string
	duplicates bool
}

func NewRemoveCommand(hostsCli cli.Cli) *cobra.Command {
	opts := &removeOptions{}

	cmd := &cobra.Command{
		Use:   "remove",
		Short: "Remove entries either by hostname, ip, or both",
		Long:  ``,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return pkg.CreateBackupFile(hostsCli.HostsFile(), "remove")
		},
		Run: func(cmd *cobra.Command, args []string) {
			defer NewListCommand(hostsCli).Run(cmd, args)

			var filtered []pkg.HostEntry
			if opts.duplicates {
				filtered = pkg.FilterOutDuplicates(hostsCli.HostsFile().Entries, opts.hostname, opts.ip)
				for _, meep := range filtered {
					fmt.Println(meep.String())
				}
			} else {
				filtered = pkg.FilterOut(hostsCli.HostsFile().Entries, opts.hostname, opts.ip)
			}

			numMatchedEntries := len(hostsCli.HostsFile().Entries) - len(filtered)
			if numMatchedEntries <= 0 {
				log.Println("No matching entries found")
				return
			}

			var didConfirm bool
			err := huh.NewConfirm().
				Title(fmt.Sprintf("Matched %d entries, are you sure you want to remove?", numMatchedEntries)).
				Value(&didConfirm).
				WithAccessible(hostsCli.AccessibleMode()).
				Run()
			if err != nil {
				log.Fatal(err)
			}

			if !didConfirm {
				return
			}

			hostsCli.HostsFile().Entries = filtered

			err = pkg.SaveToFile(hostsCli.HostsFile(), "C:\\windows\\system32\\drivers\\etc\\hosts")
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
