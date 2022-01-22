package pkg

import (
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
