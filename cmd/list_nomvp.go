package cmd

import (
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List configuration",
	Long:  `NOT IMPLEMENTED`,
	Run: func(cmd *cobra.Command, args []string) {
		listIdCmd.Run(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.AddCommand(listIdCmd)
	listCmd.AddCommand(listOriginCmd)
}

var listIdCmd = &cobra.Command{
	Use:   "id",
	Short: "List identities",
	Long: `NOT IMPLEMENTED
	
	List IDs: git-id list id
	List repos using ID: git-id list id <id slug>`,
}

var listOriginCmd = &cobra.Command{
	Use:   "origin",
	Short: "List origins",
	Long:  `NOT IMPLEMENTED`,
}
