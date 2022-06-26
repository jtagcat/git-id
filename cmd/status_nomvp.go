package cmd

import (
	"github.com/spf13/cobra"
)

// git id | git id status
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Show identity being used",
	Long:  `NOT IMPLEMENTED`,
	// also validate that the identity is up to date, and supposed to be working, maybe --testonline or sth?
}

// func init() {
// 	rootCmd.AddCommand(statusCmd)
// 	// same as git id clone: -C: act on different dir
// }
