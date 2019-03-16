// Package log provides drop-in replacement for standard log package.
// It adds log levels, io.Writer interface for output and
// named loggers with possibility of separate output.
package log

import (
	"fmt"
	"io"
	"time"
)

// It is int so that you can add your own levels and use them
// When to include message or not defined by simple comparison
// Current log level should be more or equal to the message level
// for it to be included
const (
	DISABLED level = -1 // Logging disabled. Like completely
	ERROR    level = 10 // Show errors and fatal errors
	WARNING  level = 20 // Show warnings and errors
	INFO     level = 30 // Show information messages and everything higher than that
	DEBUG    level = 40 // Show all including debug messages
)

var levelMap = map[level]string{
	DISABLED: "",
	ERROR:    "ERROR",
	WARNING:  "WARN",
	INFO:     "INFO",
	DEBUG:    "DEBUG",
}

var namedLoggers = map[string]NamedLogger{}

// GetNamedLogger will return existing named logger
// or create new one if it does not exist yet
// namespace will be used in output format. See SetFormat()
func GetNamedLogger(namespace string) NamedLogger {
	if logger, ok := namedLoggers[namespace]; ok {
		return logger
	}

	result := &logger{
		namespace:  &namespace,
		format:     defaultLogger.format,
		timeFormat: defaultLogger.timeFormat,
		output:     defaultLogger.output,
	}

	namedLoggers[namespace] = result

	return result
}

// SetLevel sets logging level
// where logging output will include
// all levels lower than currently set
// default level is INFO
func SetLevel(level level) {
	loggingLevel = level
}

// SetFormat sets logging format for default logger.
// Any other loggers that do not overwrite format will inherit it as well.
// It takes a string that is processed by text/template
// with required values in template:
//
// {{.Timestamp}} is a Timestamp format for which is set via SetTimeFormat
// {{.Namespace}} is a namespace of logger, default logger namespace is "main"
// {{.Level}} is a logging Level for current record
// {{.Message}} is actual Message contents
//
// if passed format cannot be parsed the function panics.
func SetFormat(format string) {
	var err error
	defaultLogger.format, err = defaultLogger.format.Parse(format)
	if err != nil {
		panic(err)
	}
}

// SetTimeFormat follows standard Golang time formatting
// for default logger
func SetTimeFormat(format string) {
	*defaultLogger.timeFormat = format
}

// SetOutput sets output for default logger
func SetOutput(w io.Writer) {
	output.output = w
}

func log(l *logger, lvl level, msg ...interface{}) {
	if lvl <= loggingLevel && loggingLevel >= DISABLED {
		_ = l.format.Execute(l.output.output, &rec{
			Timestamp: time.Now().Format(*l.timeFormat),
			Namespace: *l.namespace,
			Level:     levelMap[lvl],
			Message:   fmt.Sprintln(msg...),
		})
	}
}

// Log is a generic method that will write
// to default logger and can accept custom
// logging levels
func Log(lvl level, msg ...interface{}) {
	log(defaultLogger, lvl, msg...)
}

// Debug is a logging method that will write
// to default logger with DEBUG level
func Debug(msg ...interface{}) {
	log(defaultLogger, DEBUG, msg...)
}

// Info is a logging method that will write
// to default logger with INFO level
func Info(msg ...interface{}) {
	log(defaultLogger, INFO, msg...)
}

// Warn is a logging method that will write
// to default logger with WARNING level
func Warn(msg ...interface{}) {
	log(defaultLogger, WARNING, msg...)
}

// Error is a logging method that will write
// to default logger with Error level
func Error(msg ...interface{}) {
	log(defaultLogger, ERROR, msg...)
}
