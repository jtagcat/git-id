package cmd

import (
	"os"
	"path"
	"strconv"

	"github.com/jtagcat/git-id/pkg"
	"github.com/jtagcat/git-id/pkg/encoding/ssh_config"
	"github.com/jtagcat/go-shared"
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
		newSlug := args[0]

		if !flChRemote && !flChWho {
			log.Fatal().Msg("nothing to do")
		}

		r := pkg.GitOpen(flPath)
		url, err := pkg.GitRemoteURLGet0(r, remote) // o: NMVP
		if err != nil {
			log.Fatal().Err(err).Str("remote", remote).Msg("getting git remote url")
		}

		cfgPath := path.Join(flSSHConfigDir, flGIConfig_name)
		f, err := shared.OpenFileExisting(cfgPath, os.O_RDONLY)
		if err != nil {
			log.Fatal().Err(err).Str("path", cfgPath).Msg("opening ssh config")
		}
		defer f.Close()

		cfg, err := ssh_config.DecodeToRawXKeys(f, gitidSSHConfigRootXKeys, gitidSSHConfigSubXKeys)
		if err != nil {
			log.Fatal().Err(err).Msg("decoding ssh config")
		}

		remotes, err := giRemotesFromAnyHost(cfg, url.Host)
		//if err
		//TODO: len(remotes) == 0
		var matches int
		for _, r := range remotes {
			// get newUser.r.git-id
			if true {
				matches++
			}
		}
		if matches == 0 {
			log.Fatal().Str("identity", newSlug).Str("remote", url.Host).Msg("no git-id identity-remote pair matching current host found")
		}

		// we have

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

		if flChRemote {
			// url.Host =
			r.RemoteURLSetFirst(remote, url.String())
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

//TODO: allow specifying remote to switch to
// ssh config fallback? is it possible to fallback, not parallely use ~/.ssh/config or sth?

var (
	flChRemote bool
	flChWho    bool
)

func init() {
	rootCmd.AddCommand(useCmd)
	useCmd.LocalFlags().BoolVar(&flChRemote, "remote", true, "Change remote (SSH) identity")
	useCmd.LocalFlags().BoolVar(&flChWho, "who", true, "Change git identity (name,email)")
}

type stringWithIndex struct {
	str   string
	index int
}
