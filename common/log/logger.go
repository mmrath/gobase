package log

import (
	"github.com/sirupsen/logrus"
)

var Logger = newLogger()

// Event stores messages to log later, from our standard interface
type Event struct {
	id      int
	message string
}

// StandardLogger enforces specific log message formats
type StandardLogger struct {
	*logrus.Logger
}

// NewLogger initializes the standard Logger
func newLogger() *StandardLogger {
	var baseLogger = logrus.New()

	var standardLogger = &StandardLogger{baseLogger}

	standardLogger.Formatter = &logrus.JSONFormatter{}

	return standardLogger
}

// Declare variables to store log messages as new Events
var (
	invalidArgMessage      = Event{1, "Invalid arg: %s"}
	invalidArgValueMessage = Event{2, "Invalid value for argument: %s: %v"}
	missingArgMessage      = Event{3, "Missing arg: %s"}
)

// InvalidArg is a standard error message
func (l *StandardLogger) InvalidArg(argumentName string) {
	l.Errorf(invalidArgMessage.message, argumentName)
}

// InvalidArgValue is a standard error message
func (l *StandardLogger) InvalidArgValue(argumentName string, argumentValue string) {
	l.Errorf(invalidArgValueMessage.message, argumentName, argumentValue)
}

// MissingArg is a standard error message
func (l *StandardLogger) MissingArg(argumentName string) {
	l.Errorf(missingArgMessage.message, argumentName)
}

func WithField(key string, value interface{}) *logrus.Entry {
	return Logger.WithField(key, value)
}

// Trace logs a message at level Trace on the standard Logger.
func Trace(args ...interface{}) {
	Logger.Trace(args...)
}

// Debug logs a message at level Debug on the standard Logger.
func Debug(args ...interface{}) {
	Logger.Debug(args...)
}

// Print logs a message at level Info on the standard Logger.
func Print(args ...interface{}) {
	Logger.Print(args...)
}

// Info logs a message at level Info on the standard Logger.
func Info(args ...interface{}) {
	Logger.Info(args...)
}

// Warn logs a message at level Warn on the standard Logger.
func Warn(args ...interface{}) {
	Logger.Warn(args...)
}

// Warning logs a message at level Warn on the standard Logger.
func Warning(args ...interface{}) {
	Logger.Warning(args...)
}

// Error logs a message at level Error on the standard Logger.
func Error(args ...interface{}) {
	Logger.Error(args...)
}

// Panic logs a message at level Panic on the standard Logger.
func Panic(args ...interface{}) {
	Logger.Panic(args...)
}

// Fatal logs a message at level Fatal on the standard Logger then the process will exit with status set to 1.
func Fatal(args ...interface{}) {
	Logger.Fatal(args...)
}

// Tracef logs a message at level Trace on the standard Logger.
func Tracef(format string, args ...interface{}) {
	Logger.Tracef(format, args...)
}

// Debugf logs a message at level Debug on the standard Logger.
func Debugf(format string, args ...interface{}) {
	Logger.Debugf(format, args...)
}

// Printf logs a message at level Info on the standard Logger.
func Printf(format string, args ...interface{}) {
	Logger.Printf(format, args...)
}

// Infof logs a message at level Info on the standard Logger.
func Infof(format string, args ...interface{}) {
	Logger.Infof(format, args...)
}

// Warnf logs a message at level Warn on the standard Logger.
func Warnf(format string, args ...interface{}) {
	Logger.Warnf(format, args...)
}

// Warningf logs a message at level Warn on the standard Logger.
func Warningf(format string, args ...interface{}) {
	Logger.Warningf(format, args...)
}

// Errorf logs a message at level Error on the standard Logger.
func Errorf(format string, args ...interface{}) {
	Logger.Errorf(format, args...)
}

// Panicf logs a message at level Panic on the standard Logger.
func Panicf(format string, args ...interface{}) {
	Logger.Panicf(format, args...)
}

// Fatalf logs a message at level Fatal on the standard Logger then the process will exit with status set to 1.
func Fatalf(format string, args ...interface{}) {
	Logger.Fatalf(format, args...)
}

func Traceln(args ...interface{}) {
	Logger.Traceln(args...)
}

func Debugln(args ...interface{}) {
	Logger.Debugln(args...)
}

func Println(args ...interface{}) {
	Logger.Println(args...)
}

func Infoln(args ...interface{}) {
	Logger.Infoln(args...)
}

func Warnln(args ...interface{}) {
	Logger.Warnln(args...)
}

func Warningln(args ...interface{}) {
	Logger.Warningln(args...)
}

func Errorln(args ...interface{}) {
	Logger.Errorln(args...)
}

func Panicln(args ...interface{}) {
	Logger.Panicln(args...)
}

func Fatalln(args ...interface{}) {
	Logger.Fatalln(args...)
}
