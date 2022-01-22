package cmd

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// git id add
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add an identity",
	Long: `NOT IMPLEMENTED
	
	'default' is a special id.

	Usage: git-id add <remote slug> <id slug>
	Example: git-id add gh foo ~/.ssh/foo_sk --username foobar --email 'user@domain.tld' --description 'only used by git-id, for refreshing memory`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Panic().Msg("command not implemented")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

// TODO: a) use foo.gh.git-id   identifyiable, maybe more machine-manipulative
//       b) use foo.github.com  maybe benefits? *.github.com? anything else?
//       c) use foo.gh          shorter

// git-id.conf:
//Host foo.gh.git-id # foobar
//  IdentityFile ~/.ssh/foobar_sk

//TODO: establish a config file for username + email + desc
//  can do this inside git-id.conf, git-id_default.conf! commented out yaml/json? addition to objects?
