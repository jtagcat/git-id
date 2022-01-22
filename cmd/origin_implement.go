package cmd

import (
	"github.com/spf13/cobra"
)

// git id remote
var remoteCmd = &cobra.Command{
	Use:   "remote",
	Short: "Manage remotes",
}

func init() {
	rootCmd.AddCommand(remoteCmd)
	remoteCmd.AddCommand(addRemoteCmd)
	remoteCmd.AddCommand(rmRemoteCmd)
}

// git id remote add
var addRemoteCmd = &cobra.Command{
	Use:   "add",
	Short: "Add an remote",
	Long: `NOT IMPLEMENTED
	
	Usage: git-id remote add <remote slug> <actual host>`,
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
//Host <remote slug>
//  HostName github.com
//  IdentitiesOnly yes

// git id remote rm: NOMVP
var rmRemoteCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove an remote",
	Long:  `NOT IMPLEMENTED`,
	// ssh config fallback: alias deleted to just domain
}

// git id remote set: NOVMP
// username
// email

// git id remote rename: VERY NOMVP
