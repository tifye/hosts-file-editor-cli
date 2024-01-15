/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/tifye/hosts-file-editor-cli/core"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		file, err := os.Open("C:\\windows\\system32\\drivers\\etc\\hosts")
		if err != nil {
			log.Fatalf("failed opening file: %s", err)
		}
		defer file.Close()

		entries, err := core.ParseHostsFile(file)
		if err != nil {
			log.Fatalf("failed parsing file: %s", err)
		}

		if len(entries) == 0 {
			log.Println("No entries found")
		}

		for _, entry := range entries {
			fmt.Printf("%s %s\n", entry.IP, entry.Hostname)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
