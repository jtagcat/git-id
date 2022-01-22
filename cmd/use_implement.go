package cmd

import (
	"github.com/spf13/cobra"
)

// git id use
var useCmd = &cobra.Command{
	Use:   "use",
	Short: "Switch to an identity",
	Long: `NOT IMPLEMENTED
	
	Usage: git-id use <id slug>`,
}

// (not shared with clone): change remote url
// git username, email, core.sshCommand
// git config core.sshCommand = "ssh -F ~/.ssh/git-id.conf"
// track what is used where how:
//  - date
//  - git-id user-facing id and remote slugs
//  - actual username, email, config used (config loc + full core.sshCommand)
//  - actual host used (foo.gh.git-id)

// ssh config fallback? is it possible to fallback, not parallely use ~/.ssh/config or sth?

func init() {
	rootCmd.AddCommand(useCmd)
	// TODO: -C: act on different dir
}
