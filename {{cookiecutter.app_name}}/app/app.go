package app

import (
	"fmt"
	"log/slog"
	"os"

	"{{ cookiecutter.module_name }}/config"
)

func Run(cfg config.Config) {
	initLogger(cfg.Log.Level, cfg.Log.Caller)
	slog.Info("Hello World!")
}

func initLogger(logLevel string, withSource bool) error {
	var level slog.Level
	if err := level.UnmarshalText([]byte(logLevel)); err != nil {
		return fmt.Errorf("invalid log level: %s, %w", logLevel, err)
	}

	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{
		AddSource: withSource,
		Level:     level,
	})))

	return nil
}
