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
	Long: `Needs to be ran once per user account.
Moves ~/.ssh/config to ~/.ssh/global.conf.
This enables default identities, and is currently the only supported setup.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: flagify args
		//TODO: headers broken, other stuff borken, use ssh_config instead
		sshConfig_path := path.Join(sshConfig_parentdir, "config")
		baseConfig_path := path.Join(sshConfig_parentdir, baseConfig_name)

		// config → base.conf
		if err := pkg.RenameNoOverwrite(sshConfig_path, baseConfig_path); err != nil {
			// if !errors.Is(err, fs.ErrExist) { //TODO FT: better messaging? interactive overwrite?
			log.Fatal().Err(err).Msgf("Failed moving %q to %q", sshConfig_path, baseConfig_path)
		}

		// config: Include base.conf git-id_defaults.conf
		if err := os.WriteFile(sshConfig_path, []byte("Include \""+gitidDefaultsConfig_name+"\"\n"+
			"Include \""+baseConfig_name+"\"\n"), 0600); err != nil {
			log.Fatal().Err(err).Msgf("Failed creating new %q; old config at %q", sshConfig_path, baseConfig_path)
		}

		// init managed files
		for _, o := range [][]string{
			{gitidConfig_name, "Include \"" + baseConfig_name + "\"\n" + gitidHeaderInfo + "\n"},
			{gitidDefaultsConfig_name, gitidHeaderInfo + "\n"}} {
			if err := os.WriteFile(path.Join(sshConfig_parentdir, o[0]), []byte(o[1]), 0600); err != nil {
				log.Error().Err(err).Msgf("Failed creating %q", o[0])
			}
		}
		// import git-id.conf
		if err := pkg.FileAppend(baseConfig_path, []byte("Include \""+gitidConfig_name+"\"")); err != nil {
			log.Error().Err(err).Msgf("Failed adding Include to %q", baseConfig_path)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
