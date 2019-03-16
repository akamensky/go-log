package log

import (
	"os"
	"text/template"
)

// Default logger namespace
var defaultNamespace = "main"

// Current logging level
// can be changed by SetLevel()
// is inherited by NamedLogger
// but can be overwritten
var loggingLevel = INFO

// Default log format
// can be changed by SetFormat()
// is inherited by NamedLogger
// but can be overwritten
// is actually text/template format
var defaultFormat = "[{{.Timestamp}}] {{.Namespace}} {{.Level}}: {{.Message}}"

// Default time format
// can be changed by SetTimeFormat()
// is inherited by NamedLogger
// but can be overwritten
// is actually golang builtin time formatter
// (which is horrible by the way)
var defaultTimeFormat = "2006-01-02 15:04:05 -0700"

// Default output where the logs will be written
// can be changed by SetOutput()
// is inherited by NamedLogger
// but can be overwritten
// uses io.Writer as type
var output = &out{output: os.Stderr}

var defaultLogger *logger

func init() {
	format, err := template.New(defaultNamespace).Parse(defaultFormat)
	if err != nil {
		panic(err)
	}

	defaultLogger = &logger{
		namespace:  &defaultNamespace,
		format:     format,
		timeFormat: &defaultTimeFormat,
		output:     output,
	}
}
