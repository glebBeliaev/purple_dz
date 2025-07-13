package logger

import (
	"os"

	"github.com/glebbeliaev/purple_dz/config"
	"github.com/rs/zerolog"
)

func NewLogger(config *config.LogConfig) *zerolog.Logger {
	zerolog.SetGlobalLevel(zerolog.Level(config.Level))
	var logger zerolog.Logger
	if config.Format == "json" {
		logger = zerolog.New(os.Stderr).With().Timestamp().Logger()
	} else {
		consoleWriter := zerolog.ConsoleWriter{Out: os.Stderr}
		logger = zerolog.New(consoleWriter).With().Timestamp().Logger()
	}
	return &logger
}
