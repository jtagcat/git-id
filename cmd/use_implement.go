package cmd

import (
	"github.com/gogs/git-module"
	"github.com/jtagcat/git-id/pkg"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// git id use
var useCmd = &cobra.Command{
	Use:   "use",
	Short: "Switch to an identity",
	Long: `NOT IMPLEMENTED
	
	Usage: git-id use <id slug>`,
	Run: func(cmd *cobra.Command, args []string) {
		path, err := pkg.PWDIfEmpty(flPath)
		if err != nil {
			log.Fatal().Err(err).Msg("")
		}
		log.Trace().Str("git_working_directory", path).Msg("")

		r, err := git.Open(path)
		if err != nil {
			log.Fatal().Err(err).Msg("")
		}
		log.Debug().Str("path", path).Msg("repo opened")

		// if err != nil {
		// 	log.Fatal().Err(err)
		// }
		// i, err := r.Head()
		// if err != nil {
		// 	log.Fatal().Err(err).Msg("uhh")
		// }
		// log.Info().Msgf("%v", i)
		//fmt.Print(i)
		// (not shared with clone): change remote url
		// git := pkg.CheckGit(Path)
		// out, _ := git("version")
		// fmt.Print(string(out))
		// git username, email, core.sshCommand
		// git config core.sshCommand = "ssh -F ~/.ssh/git-id.conf"
		// track what is used where how:
		//  - date
		//  - git-id user-facing id and remote slugs
		//  - actual username, email, config used (config loc + full core.sshCommand)
		//  - actual host used (foo.gh.git-id)
	},
}

// ssh config fallback? is it possible to fallback, not parallely use ~/.ssh/config or sth?

func init() {
	rootCmd.AddCommand(useCmd)
}
