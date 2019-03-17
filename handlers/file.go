package handlers

import (
	"io"
	"os"
)

// GetFileHandler is a helper function that takes
// a filename, opens it for writing, and returns
// it wrapped as io.Writer. If file is not exist
// it will be created with default permissions.
// This function will panic if it failed to open
// or create a file (i.e. if directory does not
// exist).
func GetFileHandler(filename string) io.Writer {
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND|os.O_SYNC, 0644)
	if err != nil {
		panic(err)
	}

	return f
}
