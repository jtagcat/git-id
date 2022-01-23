package cmd

import (
	"strconv"

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
		if len(args) != 1 {
			log.Fatal().Str("flagcount", strconv.Itoa(len(args))).Msg("expected exactly one argument")
		}
		if !flChRemote && !flChWho {
			log.Fatal().Msg("nothing to do")
		}
		remote := "origin" // NOMVP

		r := pkg.GitOpen(flPath)
		if flChRemote {
			remoteToIdentZL(r, remote, args[0])
		}
		if flChWho {
			// for _, confopt := range [][]string{{"user.name", ident[name]}, {"user.email", ident[email]}, {"user.signingKey", ident[sigkey]}} {
			// 	if confopt[1] != "" { // NOMVP: set to empty?
			// 		r.ConfigSet(confopt[0], confopt[1])
			// 	}
			// }
			log.Warn().Msg("git config not implemented")
		}

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

func remoteToIdentZL(r *git.Repository, remote, ident string) {
	remotes, err := r.RemoteURLGet(remote)
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}
	if len(remotes) != 1 { // NOMVP
		log.Fatal().Msg("MVP needs exactly 1 remote URL for origin")
	}

	// ##### change remote url ##### (not same as clone)
	// parse; check current git-ssh domain
	//   fail if not ssh
	// does domain end with gitidTLD
	//   magic
	// else is domain registered in git-id
	//   fail
	// is ident present under domain
	//   fail
	// parsed url: replace domain (shared code with clone)
	// set-url
	// git config core
}

var (
	flChRemote bool
	flChWho    bool
)

func init() {
	rootCmd.AddCommand(useCmd)
	useCmd.LocalFlags().BoolVar(&flChRemote, "remote", true, "Change remote (SSH) identity")
	useCmd.LocalFlags().BoolVar(&flChWho, "who", true, "Change git identity (name,email)")
}
