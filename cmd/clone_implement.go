package cmd

import (
	"github.com/spf13/cobra"
)

// cloneCmd represents the clone command
var cloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "Clone a repository using an identity",
	Long: `NOT IMPLEMENTED
	
	Usage: git-id clone <id slug> [url and stuff]`, // can we say to cobra: please let's fwd all flags?
}

// replace github.com with hijacked thing
// git config execs (call git-id use)

func init() {
	rootCmd.AddCommand(cloneCmd)
}

// NOMVP: is there anything similar to clone we need to cover?
