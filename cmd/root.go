package cmd

import (
	"os"

	"github.com/jtagcat/git-id/pkg"
	"github.com/spf13/cobra"
)

var (
	gitidHeaderInfo       = "# This file is managed by git-id"
	gitidHeaderRemotes    = "\n# Remotes"
	gitidHeaderIdentities = "\n# Identities"
)

// rootCmd is the base command, 'git id'
var rootCmd = &cobra.Command{
	Use:   "git-id",
	Short: "Dumb git identity management",
	Long: `git-id speeds up setting up and managing multiple identities with git.
Configuration is only applied — after setup, git-id is not needed.

	'git-id' aliases to 'git-id status'`,
	Run: func(cmd *cobra.Command, args []string) {
		statusCmd.Run(cmd, args)
	},
    //NOTMVP: git branch, ncdu-style, whatever arrow keys / fzf / quick switcher
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	pkg.ZerologLevelStringint(os.Getenv("LOGLEVEL")) //TODO: parse -vvv and --verbose=5 / --verbose=info
}

// NOTMVP: custom core.sshCommand additions
// very NOMVP: allow hiding/deprecating an id/remote instead of rm
// TODO: rm/change/deprecate: can we use ssh_config things to print something / execute git-id hidden command?
