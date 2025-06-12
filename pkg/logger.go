package pkg

import (
	"log/slog"
	"os"
	"strings"
	"sync"

	"github.com/glebbeliaev/purple_dz/config"
)

var (
	logger *slog.Logger
	once   sync.Once
)

func MapToSlogLevel(levelStr string) slog.Level {
	switch strings.ToLower(levelStr) {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}

func NewLogger(cfg *config.LogConfig) *slog.Logger {
	level := MapToSlogLevel(cfg.Level)

	var handler slog.Handler
	switch strings.ToLower(cfg.Format) {
	case "json":
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level:     level,
			AddSource: true,
		})
	default:
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: level,
		})
	}

	logger := slog.New(handler)
	slog.SetDefault(logger)
	return logger
}
