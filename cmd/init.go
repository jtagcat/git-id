package cmd

import (
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
	Long: `Moves ~/.ssh/config to ~/.ssh/global.conf.
This enables default identities, and is currently the only supported setup.`,
	Run: func(cmd *cobra.Command, args []string) {
		sshConfig_parentdir := "~/.ssh"
		baseConfig_name := "base.conf"
		gitidConfig_name := "git-id.conf"

		sshConfig_path := path.Join(sshConfig_parentdir, "config")
		baseConfig_path := path.Join(sshConfig_parentdir, baseConfig_name)

		if err := pkg.RenameNoOverwrite(sshConfig_path, baseConfig_path); err != nil {
			// if !errors.Is(err, fs.ErrExist) { //TODO FT: better messaging? interactive overwrite?
			log.Fatal().Err(err).Msgf("Failed moving %q to %q", sshConfig_path, baseConfig_path)

		}

		if err := os.WriteFile(sshConfig_path, []byte("Include \""+baseConfig_name+"\""), 0600); err != nil {
			log.Fatal().Err(err).Msgf("Failed creating new %q; old config at %q", sshConfig_path, baseConfig_path)
		}

		if err := os.WriteFile(path.Join(sshConfig_parentdir, gitidConfig_name), []byte(""), 0600); err != nil {
			log.Error().Err(err).Msgf("Failed creating %q", gitidConfig_name)
		}

		// append to file
		if err := pkg.FileAppend(baseConfig_path, []byte("Import \""+baseConfig_name+"\"")); err != nil {
			log.Error().Err(err).Msgf("Failed adding Import to %q", baseConfig_path)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
