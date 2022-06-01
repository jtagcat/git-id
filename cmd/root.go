package cmd

import (
	"os"
	"path"

	"github.com/jtagcat/git-id/pkg"
	"github.com/jtagcat/git-id/pkg/encoding/ssh_config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var (
	gitidHeaderInfo         = " This file is managed by git-id"
	gitidHeaderDefaults     = ssh_config.RawTopLevel{Key: "XHeader", Values: []ssh_config.RawValue{{"Defaults", 0}}}
	gitidHeaderIdentities   = ssh_config.RawTopLevel{Key: "XHeader", Values: []ssh_config.RawValue{{"Identities", 0}}}
	gitidHeaderRemotes      = ssh_config.RawTopLevel{Key: "XHeader", Values: []ssh_config.RawValue{{"Remotes", 0}}}
	gitidSSHConfigRootXKeys = map[string]bool{"xheader": false}
	gitidSSHConfigSubXKeys  = []string{"XGitConfig", "XDescription"}

	remote = "origin"
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
	// NOTMVP: git branch, ncdu-style, whatever arrow keys / fzf / quick switcher
}

func Execute() {
	// TODO: DEV
	zerolog.SetGlobalLevel(zerolog.TraceLevel)
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var (
	flSSHConfigDir  string
	flGIConfig_name string
	flGI_TLD        string

	flPath string
)

func init() {
	pkg.ZerologLevelStringint(os.Getenv("LOGLEVEL")) // TODO: parse -vvv and --verbose=5 / --verbose=info

	rootCmd.PersistentFlags().StringVar(&flSSHConfigDir, "ssh-config-dir", "", "empty for ~/.ssh")
	if flSSHConfigDir == "" {
		homedir, err := os.UserHomeDir()
		if err != nil {
			log.Fatal().Err(err).Msg("--ssh-config-dir unset; can't get home directory")
		}
		flSSHConfigDir = path.Join(homedir, ".ssh")
	}
	rootCmd.PersistentFlags().StringVar(&flGIConfig_name, "gitid-config-name", "git-id.conf", "name of git-id managed configuration file")
	rootCmd.PersistentFlags().StringVar(&flGI_TLD, "gitid-tld", "git-id", "ident.remote.git-id ← TLD in remote hijacks")

	rootCmd.PersistentFlags().StringVarP(&flPath, "", "C", "", "Act on `path` instead of working directory.") //**HACK1** bugbug upstream: https://github.com/spf13/pflag/issues/139
}

// NOTMVP: custom core.sshCommand additions
// very NOMVP: allow hiding/deprecating an id/remote instead of rm
// TODO: rm/change/deprecate: can we use ssh_config things to print something / execute git-id hidden command?

// gitidConfig:
// # This file is managed by git-id
//
// #XHeader Defaults
// Match OriginalHost github.com
//   IdentityFile ~/.ssh/id_rsa
//
// #XHeader Identities
// Host jc.gh.git-id
//  IdentityFile ~/.ssh/id_rsa # this is redundant with defaults, IdentityFile is used for matching the default to an ident
//  #XGitConfig user.name jtagcat
//  #XGitConfig user.email blah
//  #XDescription uwu
// Host w.gh.git-id
//  IdentityFile ~/.ssh/work_sk
//
// #XHeader Remotes
// Host *.gh.git-id
//   Hostname github.com
//   #XDescription "iz GitHub"
