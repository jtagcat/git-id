package cmd

import (
	"os"

	"github.com/jtagcat/git-id/pkg"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
)

var (
	gitidHeaderInfo          = " This file is managed by git-id"
	gitidHeaderRemotes       = "XHeader Remotes"
	gitidHeaderIdentities    = "XHeader Identities"
	gitidTLD                 = "git-id" // foo.gh.git-id
	remote                   = "origin"
	sshConfig_parentdir      = "~/.ssh"
	gitidConfig_name         = "git-id.conf"
	gitidDefaultsConfig_name = "git-id_defaults.conf"
	baseConfig_name          = "base.conf"
)

// rootCmd is the base command, 'git id'
var rootCmd = &cobra.Command{
	Use:   "git-id",
	Short: "Dumb Git identity management",
	Long: `git-id speeds up setting up and managing multiple identities with git.
Configuration is only applied — after setup, git-id is not needed.`,
	// 'git-id' aliases to 'git-id status'`,
	//	Run: func(cmd *cobra.Command, args []string) {
	//		statusCmd.Run(cmd, args)
	//	},
	//NOTMVP: git branch, ncdu-style, whatever arrow keys / fzf / quick switcher
}

func Execute() {
	//TODO: DEV
	zerolog.SetGlobalLevel(zerolog.TraceLevel)
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var (
	flPath string
)

func init() {
	pkg.ZerologLevelStringint(os.Getenv("LOGLEVEL")) //TODO: parse -vvv and --verbose=5 / --verbose=info

	rootCmd.PersistentFlags().StringVarP(&flPath, "", "C", "", "Act on path instead of working directory.") //**HACK1** bugbug upstream: https://github.com/spf13/pflag/issues/139
}

// NOTMVP: custom core.sshCommand additions
// very NOMVP: allow hiding/deprecating an id/remote instead of rm
// TODO: rm/change/deprecate: can we use ssh_config things to print something / execute git-id hidden command?
