package slog

import (
	"runtime"
	"strings"
	"time"
	"fmt"
	"os"
	"io"
)

const (
	DISABLED = -1
	FATAL    = 0
	ERROR    = 10
	WARNING  = 20
	INFO     = 30
	DEBUG    = 40
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
		_, err := output.Write([]byte(record))
		if err != nil {
			// TODO: What should it do then?
		}
	}
	if level == FATAL {
		err := stdout.Close()
		if err != nil {
			// TODO: What should it do then?
		}
		err = stderr.Close()
		if err != nil {
			// TODO: What should it do then?
		}
		os.Exit(1)
	}
}

func SetLogLevel(level int) {
	loggingLevel = level
}

func GetLogLevel() int {
	return loggingLevel
}

func SetStdout(out io.WriteCloser) {
	if out != nil {
		stdout = out
	}
}

func SetStderr(out io.WriteCloser) {
	if out != nil {
		stderr = out
	}
}

func SetStdLogFile(filename string) {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log(FATAL, stderr, err.Error())
	}
	stdout = f
}

func SetErrLogFile(filename string) {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log(FATAL, stderr, err.Error())
	}
	stderr = f
}

func SetCombinedLogFile(filename string) {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log(FATAL, stderr, err.Error())
	}
	stdout = f
	stderr = f
}

func Log(level int, msg ...interface{}) {
	log(level, stdout, msg...)
}

func Debug(msg ...interface{}) {
	log(DEBUG, stdout, msg...)
}

func Info(msg ...interface{}) {
	log(INFO, stdout, msg...)
}

func Warn(msg ...interface{}) {
	log(WARNING, stdout, msg...)
}

func Err(msg ...interface{}) {
	log(ERROR, stderr, msg...)
}

func Fatal(msg ...interface{}) {
	log(FATAL, stderr, msg...)
}
