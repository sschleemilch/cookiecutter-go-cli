package logger

import (
	"os"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var initialized = false

func Init(logLevel string, logCaller bool, logFile string, json bool) {
	if initialized {
		return
	}
	level, err := zerolog.ParseLevel(strings.ToLower(logLevel))
	if err != nil {
		log.Fatal().Err(err).Msg("Invalid log level")
		os.Exit(1)
	}
	zerolog.SetGlobalLevel(level)

	var logFileFd *os.File
	if logFile != "" {
		logFileFd, err = os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatal().Err(err).Str("file", logFile).Str("state", "init").Msg("[Log]")
			os.Exit(1)
		}
	}

	consoleWriter := &zerolog.ConsoleWriter{Out: os.Stdout}
	if logFileFd != nil {
		if json {
			log.Logger = zerolog.New(zerolog.MultiLevelWriter(logFileFd, os.Stdout))
		} else {
			log.Logger = zerolog.New(zerolog.MultiLevelWriter(consoleWriter, logFileFd))
		}
	} else {
		if !json {
			log.Logger = zerolog.New(zerolog.MultiLevelWriter(consoleWriter))
		}
	}

	if logCaller {
		log.Logger = log.With().Caller().Logger()
	}

	log.Logger = log.With().Timestamp().Logger()
	initialized = true
}
