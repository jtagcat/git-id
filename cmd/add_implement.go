package cmd

import (
	"github.com/spf13/cobra"
)

// git-id add
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add an identity",
	Long: `NOT IMPLEMENTED
	
	'default' is a special id.

	Usage: git-id add <remote slug> <id slug> [ssh key]
	Example: git-id add gh foo ~/.ssh/foo_sk -u foobar -e 'user@domain.tld' -d 'only used by git-id, for refreshing memory`,
}

var (
	flUsername   string
	flEmail      string
	flSigningKey string
)

// func init() {
// 	rootCmd.AddCommand(addCmd)
// 	addCmd.LocalFlags().StringVarP(&flUsername, "username", "u", "", "git user.name")
// 	addCmd.LocalFlags().StringVarP(&flEmail, "email", "e", "", "git user.email")
// 	addCmd.LocalFlags().StringVarP(&flSigningKey, "sigkey", "s", "", "git user.signingKey")
// 	// addCmd.LocalFlags().StringVarP(&flDescription, "description", "d", "", "git-id-only, memory refresher")
// }

// IMPORTANT: NOMVP:
// - username-email should be seperate / child objects we fetch / they are referenced by identities
// - multiple remotes, multiple identities
