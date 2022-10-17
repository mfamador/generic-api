// Package logger defines the logger bootstrap
package logger

import (
	"genericsapi/internal/config"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var _ = func() bool { a := &config.Config; return a != nil }()

func init() {
	level := config.Config.Logger.Level
	SetPretty(config.Config.Logger.Pretty)

	switch level {
	case "panic":
		zerolog.SetGlobalLevel(zerolog.PanicLevel)
	case "fatal":
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
	case "error":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case "warn":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case "info":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}
}

// SetPretty switches pretty logs on and off
func SetPretty(b bool) {
	if b {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: "2006/01/02 15:04:05"})
	} else {
		zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs
		zerolog.MessageFieldName = "msg"
	}
}
