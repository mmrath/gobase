package log

import (
	"fmt"
	_ "github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var Logger = log.Logger

// StandardLogger enforces specific log message formats
type StandardLogger struct {
}

// NewLogger initializes the standard Logger
func newLogger() StandardLogger {
	return StandardLogger{}
}

// InvalidArg is a standard error message
func InvalidArg(argumentName string) {
	Logger.Error().Str("argName", argumentName).Msg("invalid argument")
}

// InvalidArgValue is a standard error message
func InvalidArgValue(argumentName string, argumentValue string) {
	Logger.Error().Str("argName", argumentName).Interface("argValue", argumentValue).Msg("invalid argument value")
}

// MissingArg is a standard error message
func MissingArg(argumentName string) {
	Logger.Error().Str("argName", argumentName).Msg("invalid argument")
}

func (l StandardLogger) Debug(args ...interface{}) {
	Logger.Debug().Msgf(fmt.Sprint(args))
}

func (l StandardLogger) Info(args ...interface{}) {
	Logger.Info().Msgf(fmt.Sprint(args))
}

func (l StandardLogger) Warn(args ...interface{}) {
	Logger.Warn().Msgf(fmt.Sprint(args))
}

func (l StandardLogger) Error(args ...interface{}) {
	Logger.Error().Msgf(fmt.Sprint(args))
}

func (l StandardLogger) Panic(args ...interface{}) {
	Logger.Panic().Msgf(fmt.Sprint(args))
}

func (l StandardLogger) Fatal(args ...interface{}) {
	Logger.Fatal().Msgf(fmt.Sprint(args))
}

func (l StandardLogger) Debugf(format string, args ...interface{}) {
	Logger.Debug().Msgf(format, args...)
}

func (l StandardLogger) Infof(format string, args ...interface{}) {
	Logger.Info().Msgf(format, args...)
}

func (l StandardLogger) Warnf(format string, args ...interface{}) {
	Logger.Warn().Msgf(format, args...)
}

func (l StandardLogger) Errorf(format string, args ...interface{}) {
	Logger.Error().Msgf(format, args...)
}

func (l StandardLogger) Panicf(format string, args ...interface{}) {
	Logger.Panic().Msgf(format, args...)
}

func (l StandardLogger) Fatalf(format string, args ...interface{}) {
	Logger.Fatal().Msgf(format, args...)
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
