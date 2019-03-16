package log

import (
	"io"
	"text/template"
)

type out struct {
	output io.Writer
}

type rec struct {
	Timestamp string
	Namespace string
	Level     string
	Message   string
}

type level int

// NamedLogger is a named logger
// where each part of application
// can write to their own named
// logger, with all having different
// level, output, format etc set
type NamedLogger interface {
	// Log is a generic method that will write
	// to named logger and can accept custom
	// logging levels
	Log(lvl level, msg ...interface{})

	// Debug is a logging method that will write
	// to named logger with DEBUG level
	Debug(msg ...interface{})

	// Info is a logging method that will write
	// to named logger with INFO level
	Info(msg ...interface{})

	// Warn is a logging method that will write
	// to named logger with WARNING level
	Warn(msg ...interface{})

	// Error is a logging method that will write
	// to named logger with ERROR level
	Error(msg ...interface{})

	// SetFormat sets logging format for named logger.
	// It takes a string that is processed by text/template
	// with required values in template:
	//
	// {{.Timestamp}} is a Timestamp format for which is set via SetTimeFormat
	// {{.Namespace}} is a namespace of logger, default logger namespace is "main"
	// {{.Level}} is a logging Level for current record
	// {{.Message}} is actual Message contents
	//
	// If passed format cannot be parsed the function panics.
	SetFormat(format string) NamedLogger

	// SetTimeFormat is a method to set timestamp formatting for
	// named logger separately from other loggers.
	// Returns itself for easy chaining.
	SetTimeFormat(format string) NamedLogger

	// SetTimeFormat is a method to set output for
	// named logger separately from other loggers.
	// Returns itself for easy chaining.
	SetOutput(w io.Writer) NamedLogger
}

type logger struct {
	namespace  *string
	output     *out
	format     *template.Template
	timeFormat *string
}

// Log is a generic method that will write
// to named logger and can accept custom
// logging levels
func (l *logger) Log(lvl level, msg ...interface{}) {
	log(l, lvl, msg...)
}

// Debug is a logging method that will write
// to named logger with DEBUG level
func (l *logger) Debug(msg ...interface{}) {
	log(l, DEBUG, msg...)
}

// Info is a logging method that will write
// to named logger with INFO level
func (l *logger) Info(msg ...interface{}) {
	log(l, INFO, msg...)
}

// Warn is a logging method that will write
// to named logger with WARNING level
func (l *logger) Warn(msg ...interface{}) {
	log(l, WARNING, msg...)
}

// Error is a logging method that will write
// to named logger with ERROR level
func (l *logger) Error(msg ...interface{}) {
	log(l, ERROR, msg...)
}

// SetFormat sets logging format for named logger.
// It takes a string that is processed by text/template
// with required values in template:
//
// {{.Timestamp}} is a Timestamp format for which is set via SetTimeFormat
// {{.Namespace}} is a namespace of logger, default logger namespace is "main"
// {{.Level}} is a logging Level for current record
// {{.Message}} is actual Message contents
//
// If passed format cannot be parsed the function panics.
// Returns itself for easy chaining.
func (l *logger) SetFormat(format string) NamedLogger {
	var err error
	l.format, err = template.New(*l.namespace).Parse(format)
	if err != nil {
		panic(err)
	}

	return l
}

// SetTimeFormat is a method to set timestamp formatting for
// named logger separately from other loggers.
// Returns itself for easy chaining.
func (l *logger) SetTimeFormat(format string) NamedLogger {
	l.timeFormat = &format
	return l
}

// SetTimeFormat is a method to set output for
// named logger separately from other loggers.
// Returns itself for easy chaining.
func (l *logger) SetOutput(w io.Writer) NamedLogger {
	l.output = &out{output: w}
	return l
}
