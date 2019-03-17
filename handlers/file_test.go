package handlers

import (
	"github.com/akamensky/go-log"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestFileHandler(t *testing.T) {
	filename := "./test.log"

	log.SetHandler(GetFileHandler(filename))

	log.SetLevel(log.INFO)

	log.SetFormat("{{.Message}}")

	testString := "Test"

	log.Info(testString)

	// Now read the file and compare the content
	f, err := os.Open(filename)
	if err != nil {
		t.Errorf("Cannot open log file '%s' because of: %s", filename, err)
		return
	}
	defer func() {
		_ = os.Remove(filename)
	}()

	contentBytes, err := ioutil.ReadAll(f)
	if err != nil {
		t.Errorf("Cannot read log file '%s' because of: %s", filename, err)
		return
	}

	content := strings.TrimSpace(string(contentBytes))

	if content != testString {
		t.Errorf("Unexpected file contents. Expected '%s', got '%s'", testString, content)
	}
}
