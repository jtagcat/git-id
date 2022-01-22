package cmd

import (
	"os"

	"github.com/jtagcat/git-id/pkg"
	"github.com/spf13/cobra"
)

// rootCmd is the base command, 'git id'
var rootCmd = &cobra.Command{
	Use:   "git-id",
	Short: "Dumb git identity management",
	Long: `git-id speeds up setting up and managing multiple identities with git.
Configuration is only applied — after setup, git-id is not needed.`,

	//TODO: output similar to 'git status': is git-id used here, what is used, and does it look legit?
	// Run: func(cmd *cobra.Command, args []string) { },
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
