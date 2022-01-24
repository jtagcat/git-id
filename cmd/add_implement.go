package cmd

import (
	"github.com/spf13/cobra"
)

// git id add
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add an identity",
	Long: `NOT IMPLEMENTED
	
	'default' is a special id.

	Usage: git-id add <remote slug> <id slug> [ssh key]
	Example: git-id add gh foo ~/.ssh/foo_sk -u foobar -e 'user@domain.tld' -d 'only used by git-id, for refreshing memory`,
}

var (
	flUsername    string
	flEmail       string
	flSigningKey  string
	flDescription string
)

func init() {
	rootCmd.AddCommand(addCmd)
	useCmd.LocalFlags().StringVarP(&flUsername, "username", "u", "", "git user.name")
	useCmd.LocalFlags().StringVarP(&flEmail, "email", "e", "", "git user.email")
	useCmd.LocalFlags().StringVarP(&flSigningKey, "sigkey", "s", "", "git user.signingKey")
	useCmd.LocalFlags().StringVarP(&flDescription, "description", "d", "", "git-id-only, memory refresher")
}

// using foo.gh.git-id

// git-id.conf:
//Host foo.gh.git-id # foobar
//  IdentityFile ~/.ssh/foobar_sk

//TODO: establish a config file for username + email + desc
//  can do this inside git-id.conf, git-id_default.conf! commented out yaml/json? addition to objects?

//IMPORTANT: NOMVP:
// - username-email should be seperate / child objects we fetch / they are referenced by identities
// - multiple remotes, multiple identities
