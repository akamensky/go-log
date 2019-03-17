package log

import (
	"bytes"
	"strings"
	"testing"
)

func TestSetHandler(t *testing.T) {
	buff := bytes.NewBuffer(nil)

	SetHandler(buff)

	if defaultLogger.output.output != buff {
		t.Error("default logger output does not match with expected value")
	}
}

func TestSetLevel(t *testing.T) {
	SetLevel(INFO)

	if loggingLevel != INFO {
		t.Errorf("Unexpected default logging level. Expected [%d], got [%d]", INFO, loggingLevel)
	}
}

func TestLog(t *testing.T) {
	buff := bytes.NewBuffer(nil)

	SetHandler(buff)

	SetLevel(DEBUG)

	Log(DEBUG, "Test")

	str := strings.TrimSpace(buff.String())

	if !strings.HasSuffix(str, "Test") {
		t.Errorf("Unexpected log record. Expected ending on [%s], got [%s]", "Test", str)
	}
}

func TestInfo(t *testing.T) {
	buff := bytes.NewBuffer(nil)

	SetHandler(buff)

	SetLevel(INFO)

	Info("Test")

	str := strings.TrimSpace(buff.String())

	if !strings.HasSuffix(str, "Test") {
		t.Errorf("Unexpected log record. Expected ending on [%s], got [%s]", "Test", str)
	}
}

func TestWarn(t *testing.T) {
	buff := bytes.NewBuffer(nil)

	SetHandler(buff)

	SetLevel(WARNING)

	Warn("Test")

	str := strings.TrimSpace(buff.String())

	if !strings.HasSuffix(str, "Test") {
		t.Errorf("Unexpected log record. Expected ending on [%s], got [%s]", "Test", str)
	}
}

func TestErr(t *testing.T) {
	buff := bytes.NewBuffer(nil)

	SetHandler(buff)

	SetLevel(ERROR)

	Error("Test")

	str := strings.TrimSpace(buff.String())

	if !strings.HasSuffix(str, "Test") {
		t.Errorf("Unexpected log record. Expected ending on [%s], got [%s]", "Test", str)
	}
}

func TestDebug(t *testing.T) {
	buff := bytes.NewBuffer(nil)

	SetHandler(buff)

	SetLevel(DEBUG)

	Debug("Test")

	str := strings.TrimSpace(buff.String())

	if !strings.HasSuffix(str, "Test") {
		t.Errorf("Unexpected log record. Expected ending on [%s], got [%s]", "Test", str)
	}
}

func TestNamedLoggerSameOutput(t *testing.T) {
	buff := bytes.NewBuffer(nil)

	namedLogger := GetNamedLogger("sublogger")

	SetHandler(buff)

	SetLevel(INFO)

	namedLogger.Info("Test")

	str := strings.TrimSpace(buff.String())

	if !strings.HasSuffix(str, "Test") {
		t.Errorf("Unexpected log record. Expected ending on [%s], got [%s]", "Test", str)
	}

	buff = bytes.NewBuffer(nil)

	SetHandler(buff)

	SetLevel(INFO)

	Info("Test2")

	str = strings.TrimSpace(buff.String())

	if !strings.HasSuffix(str, "Test2") {
		t.Errorf("Unexpected log record. Expected ending on [%s], got [%s]", "Test2", str)
	}
}

func TestNamedLoggerDiffOutput(t *testing.T) {
	buff := bytes.NewBuffer(nil)
	namedBuff := bytes.NewBuffer(nil)

	namedLogger := GetNamedLogger("sublogger")
	namedLogger.SetOutput(namedBuff)

	SetHandler(buff)

	SetLevel(INFO)

	Info("Default")
	namedLogger.Info("Named")

	str := strings.TrimSpace(buff.String())
	namedStr := strings.TrimSpace(namedBuff.String())

	if !strings.HasSuffix(str, "Default") {
		t.Errorf("Unexpected log record. Expected ending on [%s], got [%s]", "Default", str)
	}
	if !strings.HasSuffix(namedStr, "Named") {
		t.Errorf("Unexpected log record. Expected ending on [%s], got [%s]", "Named", namedStr)
	}
}
