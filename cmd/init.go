package cmd

import (
	"io/fs"
	"os"
	"path"

	"github.com/jtagcat/git-id/pkg"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// git id init
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize git-id",
	Long: `Ran once per user account.
Moves ~/.ssh/config to ~/.ssh/global.conf.
This enables default identities, and is currently the only supported setup.`,
	Run: func(_ *cobra.Command, args []string) {
		sshConfig_path := path.Join(flSSHConfigDir, "config")
		gitidConfig_path := path.Join(flSSHConfigDir, flGIConfig_name)

		// init git-id.conf
		if _, err := os.Stat(gitidConfig_path); err == fs.ErrNotExist {
			// write only if doesn't exist
			if err := os.WriteFile(gitidConfig_path, []byte(gitidHeaderInfo+"\n"), 0o600); err != nil {
				log.Error().Err(err).Msgf("Failed creating %q", gitidConfig_path)
			}
		} else if err != nil {
			log.Fatal().Err(err).Msgf("Failed stating %q", gitidConfig_path)
		} else if err == nil {
			log.Info().Str("path", gitidConfig_path).Msg("git-id config already exists")
		}

		// TODO: is already included or not?
		// include git-id.conf
		if err := pkg.FileAppend(sshConfig_path, []byte("Include '"+flGIConfig_name+"'")); err != nil {
			log.Error().Err(err).Msgf("Failed adding Include to %q", sshConfig_path)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
