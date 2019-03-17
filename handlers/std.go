package handlers

import (
	"io"
	"os"
)

func GetStdoutHandler() io.Writer {
	return os.Stdout
}

func GetStderrHandler() io.Writer {
	return os.Stderr
}
