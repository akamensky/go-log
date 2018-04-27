// Package slog provides drop-in replacement for standard log package. It adds log levels and separating stdout and stderr.
// NOTE: It does not attempt to replace os.Stdout and os.Stderr, which might not be wise for some situations
package slog

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"
	"time"
)

// It is int so that you can add your own levels and use them
// When to include message or not defined by simple comparison
// Current log level should be more or equal to the message level
// for it to be included
const (
	DISABLED = -1 // Logging disabled. Like completely
	FATAL    = 0  // Only show fatal errors when unable to recover and decide to exit
	ERROR    = 10 // Show errors and fatal errors
	WARNING  = 20 // Show warnings and errors
	INFO     = 30 // Show information messages and everything higher than that
	DEBUG    = 40 // Show all including debug messages
)

var levelTag = map[int]string{
	DISABLED: "",
	FATAL:    "FATAL",
	ERROR:    "ERROR",
	WARNING:  "WARN",
	INFO:     "INFO",
	DEBUG:    "DEBUG",
}

var loggingLevel = DEBUG

var stdout io.WriteCloser = os.Stdout
var stderr io.WriteCloser = os.Stderr

var format = "[%s] %s: %s%s"

func log(level int, output io.WriteCloser, msg ...interface{}) {
	var callee string
	if level <= loggingLevel && loggingLevel >= DISABLED {
		if level >= DEBUG {
			pc, _, _, ok := runtime.Caller(2)
			details := runtime.FuncForPC(pc)
			if ok && details != nil {
				fullName := details.Name()
				splitName := strings.Split(fullName, "/")
				callee = splitName[len(splitName)-1:][0] + ": "
			}
		}
		t := time.Now().Format("2006-01-02 15:04:05 -0700")
		record := fmt.Sprintf(format+"\n", t, levelTag[level], callee, fmt.Sprint(msg...))
		output.Write([]byte(record))
	}
	if level == FATAL {
		stdout.Close()
		stderr.Close()
		os.Exit(1)
	}
}

// SetLogLevel will adjust current log level
// default is DEBUG == 40
func SetLogLevel(level int) {
	loggingLevel = level
}

// GetLogLevel returns current log level as int value
func GetLogLevel() int {
	return loggingLevel
}

// GetStdout returns currently set stdout io.WriteCloser interface
func GetStdout() io.WriteCloser {
	return stdout
}

// GetStderr returns currently set stderr io.WriteCloser interface
func GetStderr() io.WriteCloser {
	return stderr
}

// GetOutputs returns both at once
func GetOutputs() (io.WriteCloser, io.WriteCloser) {
	return stdout, stderr
}

// SetStdout changes current stdout to provided io.WriteCloser. Default is os.Stdout
func SetStdout(out io.WriteCloser) {
	if out != nil {
		stdout = out
	}
}

// SetStderr changes current stderr to provided io.WriteCloser. Default is os.Stderr
func SetStderr(out io.WriteCloser) {
	if out != nil {
		stderr = out
	}
}

// SetCombinedOutput changes both stdout and stderr to provided io.WriteCloser.
// This will result in all methods writing to the same destination
func SetCombinedOutput(out io.WriteCloser) {
	if out != nil {
		stdout = out
		stderr = out
	}
}

// SetStdLogFile takes filename string and attempts to open it with minimum privileges
// if that fails then error is returned
// on success open file will be used instead of stdout
func SetStdLogFile(filename string) error {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	stdout = f

	return nil
}

// SetErrLogFile takes filename string and attempts to open it with minimum privileges
// if that fails then error is returned
// on success open file will be used instead of stderr
func SetErrLogFile(filename string) error {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	stderr = f

	return nil
}

// SetCombinedLogFile takes filename string and attempts to open it with minimum privileges
// if that fails then error is returned
// on success open file will be used instead of both stdout and stderr
func SetCombinedLogFile(filename string) error {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	stdout = f
	stderr = f

	return nil
}

// Log will send new log entry to stdout assuming provided level value passes currently set level
func Log(level int, msg ...interface{}) {
	log(level, stdout, msg...)
}

// Debug will send new log entry to stdout with DEBUG level
// Additionally debug will add Package (only final bit) and callee function name to the message
func Debug(msg ...interface{}) {
	log(DEBUG, stdout, msg...)
}

// Info will send new log entry to stdout with INFO level
func Info(msg ...interface{}) {
	log(INFO, stdout, msg...)
}

// Warn will send new log entry to stdout with WARNING level
func Warn(msg ...interface{}) {
	log(WARNING, stdout, msg...)
}

// Err will send new log entry to stderr with ERROR level
func Err(msg ...interface{}) {
	log(ERROR, stderr, msg...)
}

// Fatal assumes unrecoverable error
// it will attempt to send log entry to stderr with FATAL level
// after which it will exit with exit code 1.
// However success of writing is not guaranteed, depending on nature of error
func Fatal(msg ...interface{}) {
	log(FATAL, stderr, msg...)
}
