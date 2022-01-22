package cmd

import (
	"github.com/spf13/cobra"
)

// git id origin
var originCmd = &cobra.Command{
	Use:   "origin",
	Short: "Manage origins",
}

func init() {
	rootCmd.AddCommand(originCmd)
	originCmd.AddCommand(addOriginCmd)
	originCmd.AddCommand(rmOriginCmd)
}

// git id origin add
var addOriginCmd = &cobra.Command{
	Use:   "add",
	Short: "Add an origin",
	Long: `NOT IMPLEMENTED
	
	Usage: git-id origin add <origin slug> <actual host>`,
	// log.Info().Msg( it is reccommened to add a default identity bla
	// "this may be uesd by random stuff on your system,
	// system might behave weirdly if this can't be used noninteractively"
}

// ~/.ssh/config:
//Host github.com # default
//  IdentityFile ~/.ssh/gh_rsa
//NOTE: this may be uesd by random stuff on your system,
//NOTE: system might behave weirdly if this can't be used noninteractively

// ~/.ssh/git-id.conf:
//Host <origin slug>
//  HostName github.com
//  IdentitiesOnly yes

// git id origin rm: NOMVP
var rmOriginCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove an origin",
	Long:  `NOT IMPLEMENTED`,
	// ssh config fallback: alias deleted to just domain
}

// git id origin set: NOVMP
// username
// email

// git id origin rename: VERY NOMVP
