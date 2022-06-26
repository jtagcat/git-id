package cmd

import (
	"github.com/spf13/cobra"
)

// git id rm
var rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove an identity",
	Long:  `NOT IMPLEMENTED`, // not MVP
}

// ssh config fallback: alias deleted to default

// func init() {
// 	rootCmd.AddCommand(rmCmd)
// }
