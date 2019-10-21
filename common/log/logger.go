package log

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// StandardLogger enforces specific log message formats
type StandardLogger struct {
	logger zerolog.Logger
}

// NewLogger initializes the standard Logger
func NewLogger() StandardLogger {
	return StandardLogger{
		logger: zerolog.New(os.Stderr).With().Timestamp().Logger(),
	}
}


func (l StandardLogger) Debug(args ...interface{}) {
	l.logger.Debug().Msgf(fmt.Sprint(args))
}

func (l StandardLogger) Info(args ...interface{}) {
	l.logger.Info().Msgf(fmt.Sprint(args))
}

func (l StandardLogger) Warn(args ...interface{}) {
	l.logger.Warn().Msgf(fmt.Sprint(args))
}

func (l StandardLogger) Error(args ...interface{}) {
	l.logger.Error().Msgf(fmt.Sprint(args))
}

func (l StandardLogger) Panic(args ...interface{}) {
	l.logger.Panic().Msgf(fmt.Sprint(args))
}

func (l StandardLogger) Fatal(args ...interface{}) {
	l.logger.Fatal().Msgf(fmt.Sprint(args))
}

func (l StandardLogger) Debugf(format string, args ...interface{}) {
	l.logger.Debug().Msgf(format, args...)
}

func (l StandardLogger) Infof(format string, args ...interface{}) {
	l.logger.Info().Msgf(format, args...)
}

func (l StandardLogger) Warnf(format string, args ...interface{}) {
	l.logger.Warn().Msgf(format, args...)
}

func (l StandardLogger) Errorf(format string, args ...interface{}) {
	l.logger.Error().Msgf(format, args...)
}

func (l StandardLogger) Panicf(format string, args ...interface{}) {
	l.logger.Panic().Msgf(format, args...)
}

func (l StandardLogger) Fatalf(format string, args ...interface{}) {
	l.logger.Fatal().Msgf(format, args...)
}






func Debugf(format string, args ...interface{}) {
	log.Debug().Msgf(format, args...)
}

func Infof(format string, args ...interface{}) {
	log.Info().Msgf(format, args...)
}

func Warnf(format string, args ...interface{}) {
	log.Warn().Msgf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	log.Error().Msgf(format, args...)
}

func Panicf(format string, args ...interface{}) {
	log.Panic().Msgf(format, args...)
}

func Fatalf(format string, args ...interface{}) {
	log.Fatal().Msgf(format, args...)
}
