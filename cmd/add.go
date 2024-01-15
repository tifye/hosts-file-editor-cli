package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	hostname string
	ip       string
	comment  string
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new entry to the hosts file",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s %s %s", hostname, ip, comment)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringVarP(&hostname, "hostname", "n", "", "Hostname")
	addCmd.MarkFlagRequired("hostname")

	addCmd.Flags().StringVar(&ip, "ip", "", "IP")
	addCmd.MarkFlagRequired("ip")

	addCmd.Flags().StringVarP(&comment, "comment", "c", "", "Comment")
}
