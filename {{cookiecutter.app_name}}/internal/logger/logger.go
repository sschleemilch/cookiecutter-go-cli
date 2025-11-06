package logger

import (
	"fmt"
	"log/slog"
	"os"
)

var initialized = false

func parseLevel(logLevel string) (slog.Level, error) {
	var level slog.Level
	err := level.UnmarshalText([]byte(logLevel))
	return level, err
}

func Init(logLevel string, logCaller bool, asJson bool) error {
	if initialized {
		return nil
	}

	level, err := parseLevel(logLevel)
	if err != nil {
		return fmt.Errorf("parsing level: %w", err)
	}

	options := &slog.HandlerOptions{
		AddSource: logCaller,
		Level:     level,
	}

	var handler slog.Handler
	if asJson {
		handler = slog.NewJSONHandler(os.Stdout, options)
	} else {
		handler = slog.NewTextHandler(os.Stdout, options)
	}
	slog.SetDefault(slog.New(handler))

	initialized = true

	return nil
}
