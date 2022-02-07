package pkg

import (
	"github.com/gogs/git-module"
	"github.com/jtagcat/go-shared"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// ZerologLevelStringint(os.Getenv("LOGLEVEL"))
func ZerologLevelStringint(loglevel_string string) (loglevel zerolog.Level) {
	if loglevel_string != "" {
		loglevel, err := zerolog.ParseLevel(loglevel_string)
		if err != nil {
			log.Panic().Err(err).Msg("Failed to parse log level. Should be -1..5 (lower = more verbose)")
		}
		zerolog.SetGlobalLevel(loglevel)
	}
	return zerolog.GlobalLevel()
}
func GitOpen(path string) *git.Repository {
	// can I haz git?
	v, err := git.BinVersion()
	if err != nil {
		log.Fatal().Err(err)
	}
	log.Debug().Str("git_version", v).Msg("")

	// parse path
	path, err = shared.PWDIfEmpty(path)
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}
	log.Trace().Str("git_working_directory", path).Msg("")

	// open repo
	r, err := git.Open(path)
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}
	log.Debug().Str("path", path).Msg("repo opened")
	return r
}
