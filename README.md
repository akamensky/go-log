# Golang slog
[![GoDoc](https://godoc.org/github.com/akamensky/slog?status.svg)](https://godoc.org/github.com/akamensky/slog) [![Go Report Card](https://goreportcard.com/badge/github.com/akamensky/slog)](https://goreportcard.com/report/github.com/akamensky/slog) [![cover.run](https://cover.run/go/github.com/akamensky/slog.svg?style=flat&tag=golang-1.10)](https://cover.run/go?tag=golang-1.10&repo=github.com%2Fakamensky%2Fslog) [![Build Status](https://travis-ci.org/akamensky/slog.svg?branch=master)](https://travis-ci.org/akamensky/slog)

Simple drop-in replacement for golang log package which adds log levels 
and better log message formatting.

#### Install

```
$ go get -u -v -x github.com/akamensky/slog
```

#### Usage example

```go
package main

import (
	"github.com/akamensky/slog"
)

func main() {
	// Set log files (or can be anything else implementing io.WriteCloser)
	logFileName := "/var/log/testapp.log"
	errLogFileName := "/var/log/testapp-error.log"
	err := slog.SetStdLogFile(logFileName)
	if err != nil {
		panic(err)
	}
	err = slog.SetErrLogFile(errLogFileName)
	if err != nil {
		panic(err)
	}
	
	// Set log level
	slog.SetLogLevel(slog.WARNING)
	
	// Write your messages
	slog.Debug("Debug message, won't show up because of log level")
	slog.Info("Info message, won't show up because of log level")
	slog.Warn("Warning message", "something might go wrong")
	slog.Err("Error message", "some error happened")
	
	// Adjust log level while still running
	slog.SetLogLevel(slog.DEBUG)
	
	// Write some more messages
	slog.Debug("Debug message will go to log")
	
	// Hm, something went horribly wrong, I better ditch it while I can
	slog.Fatal("Something horrible happened. I am ditching this s**t")
}
```

#### License

[Read license here](https://github.com/akamensky/slog/blob/master/LICENSE)
