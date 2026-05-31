// Package logger implements the app wide logger
package logger

import (
	"log/slog"
	"os"
	"strings"
)

func InitLogger(env string, logLevel string) {
	var programLevel slog.Level

	// Parse log level string
	switch strings.ToLower(logLevel) {
	case "debug":
		programLevel = slog.LevelDebug
	case "warn":
		programLevel = slog.LevelWarn
	case "error":
		programLevel = slog.LevelError
	default:
		programLevel = slog.LevelDebug
	}
	opts := &slog.HandlerOptions{
		Level: programLevel,
	}

	var handler slog.Handler

	// Choose format based on environment
	if strings.ToLower(env) == "production" {
		handler = slog.NewJSONHandler(os.Stderr, opts)
	} else {
		handler = slog.NewTextHandler(os.Stdout, opts)
	}
	globalLogger := slog.New(handler)
	slog.SetDefault(globalLogger)
}
