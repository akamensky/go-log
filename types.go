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
	// Log at custom level
	Log(lvl level, msg ...interface{})
	// Log at predefined levels
	Debug(msg ...interface{})
	Info(msg ...interface{})
	Warn(msg ...interface{})
	Error(msg ...interface{})

	SetFormat(format string) NamedLogger
	SetTimeFormat(format string) NamedLogger
	SetOutput(w io.Writer) NamedLogger
}

type logger struct {
	namespace  *string
	output     *out
	format     *template.Template
	timeFormat *string
}

func (l *logger) Log(lvl level, msg ...interface{}) {
	log(l, lvl, msg...)
}

func (l *logger) Debug(msg ...interface{}) {
	log(l, DEBUG, msg...)
}

func (l *logger) Info(msg ...interface{}) {
	log(l, INFO, msg...)
}

func (l *logger) Warn(msg ...interface{}) {
	log(l, WARNING, msg...)
}

func (l *logger) Error(msg ...interface{}) {
	log(l, ERROR, msg...)
}

func (l *logger) SetFormat(format string) NamedLogger {
	var err error
	l.format, err = template.New(*l.namespace).Parse(format)
	if err != nil {
		panic(err)
	}

	return l
}

func (l *logger) SetTimeFormat(format string) NamedLogger {
	l.timeFormat = &format
	return l
}

func (l *logger) SetOutput(w io.Writer) NamedLogger {
	l.output = &out{output: w}
	return l
}
