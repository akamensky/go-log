package slog

import (
	"bytes"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"testing"
)

type BufferCloser struct {
	*bytes.Buffer
}

func (o *BufferCloser) Close() (err error) {
	return
}

func TestSetCombinedLogFile(t *testing.T) {
	logFileName := "./testcombined.log"

	SetCombinedLogFile(logFileName)

	if stdout != stderr {
		t.Error("stdout does not match stderr")
	}

	syscall.Unlink(logFileName)

	err := SetCombinedLogFile("")
	if err != nil {
		return
	}

	t.Errorf("Unexpected result. Expected error, got none")
}

func TestSetStdLogFile(t *testing.T) {
	logFileName := "./testout.log"

	SetStdLogFile(logFileName)

	if stdout == os.Stdout || stdout == stderr {
		t.Error("stdout either was not changed or it is same with stderr")
	}

	if f, ok := stdout.(*os.File); ok {
		if f.Name() != logFileName {
			t.Errorf("log file name does not match. expected [%s], got [%s]", logFileName, f.Name())
		}
	} else {
		t.Error("stdout is not a file")
	}

	syscall.Unlink(logFileName)

	err := SetStdLogFile("")
	if err != nil {
		return
	}

	t.Errorf("Unexpected result. Expected error, got none")
}

func TestSetErrLogFile(t *testing.T) {
	logFileName := "./testerr.log"

	SetErrLogFile(logFileName)

	if stderr == os.Stderr || stdout == stderr {
		t.Error("stderr either was not changed or it is same with stdout")
	}

	if f, ok := stderr.(*os.File); ok {
		if f.Name() != logFileName {
			t.Errorf("log file name does not match. expected [%s], got [%s]", logFileName, f.Name())
		}
	} else {
		t.Error("stdout is not a file")
	}

	syscall.Unlink(logFileName)

	err := SetErrLogFile("")
	if err != nil {
		return
	}

	t.Errorf("Unexpected result. Expected error, got none")
}

func TestSetCombinedOutput(t *testing.T) {
	outBuff := &BufferCloser{bytes.NewBuffer(nil)}

	SetCombinedOutput(outBuff)

	if stdout != outBuff || stderr != outBuff {
		t.Error("stdout or stderr does not match with expected value")
	}
}

func TestSetStdout(t *testing.T) {
	stdoutBuff := &BufferCloser{bytes.NewBuffer(nil)}

	SetStdout(stdoutBuff)

	if stdout != stdoutBuff {
		t.Error("stderr does not match with expected value")
	}
}

func TestSetStderr(t *testing.T) {
	stderrBuff := &BufferCloser{bytes.NewBuffer(nil)}

	SetStderr(stderrBuff)

	if stderr != stderrBuff {
		t.Error("stderr does not match with expected value")
	}
}

func TestSetLogLevel(t *testing.T) {
	if loggingLevel != DEBUG {
		t.Errorf("Unexpected default logging level. Expected [%d], got [%d]", DEBUG, loggingLevel)
	}

	SetLogLevel(INFO)

	if loggingLevel != INFO {
		t.Errorf("Unexpected default logging level. Expected [%d], got [%d]", INFO, loggingLevel)
	}
}

func TestLog(t *testing.T) {
	stdoutBuff := &BufferCloser{bytes.NewBuffer(nil)}

	SetStdout(stdoutBuff)

	SetLogLevel(DEBUG)

	Log(DEBUG, "Test")

	str := stdoutBuff.String()

	if !strings.HasSuffix(str, "Test\n") {
		t.Errorf("Unexpected log record. Expected ending on [%s], got [%s]", "Test\n", str)
	}
}

func TestInfo(t *testing.T) {
	stdoutBuff := &BufferCloser{bytes.NewBuffer(nil)}

	SetStdout(stdoutBuff)

	SetLogLevel(INFO)

	Info("Test")

	str := stdoutBuff.String()

	if !strings.HasSuffix(str, "Test\n") {
		t.Errorf("Unexpected log record. Expected ending on [%s], got [%s]", "Test\n", str)
	}
}

func TestWarn(t *testing.T) {
	stdoutBuff := &BufferCloser{bytes.NewBuffer(nil)}

	SetStdout(stdoutBuff)

	SetLogLevel(WARNING)

	Warn("Test")

	str := stdoutBuff.String()

	if !strings.HasSuffix(str, "Test\n") {
		t.Errorf("Unexpected log record. Expected ending on [%s], got [%s]", "Test\n", str)
	}
}

func TestErr(t *testing.T) {
	stderrBuff := &BufferCloser{bytes.NewBuffer(nil)}

	SetStderr(stderrBuff)

	SetLogLevel(ERROR)

	Err("Test")

	str := stderrBuff.String()

	if !strings.HasSuffix(str, "Test\n") {
		t.Errorf("Unexpected log record. Expected ending on [%s], got [%s]", "Test\n", str)
	}
}

func TestFatal(t *testing.T) {
	if os.Getenv("FATAL") == "1" {
		SetLogLevel(FATAL)
		Fatal("Test")
	} else {
		cmd := exec.Command(os.Args[0], "-test.run=TestFatal")
		cmd.Env = append(os.Environ(), "FATAL=1")
		output, err := cmd.CombinedOutput()
		if e, ok := err.(*exec.ExitError); ok && !e.Success() && strings.HasSuffix(string(output), "Test\n") {
			return
		}
		t.Errorf("Unexpected output or exit code. Expected output ending [%s], got [%s]", "Test\n", string(output))
	}
}

func TestDebug(t *testing.T) {
	stdoutBuff := &BufferCloser{bytes.NewBuffer(nil)}

	SetStdout(stdoutBuff)

	SetLogLevel(DEBUG)

	Debug("Test")

	str := stdoutBuff.String()

	if !strings.HasSuffix(str, "Test\n") {
		t.Errorf("Unexpected log record. Expected ending on [%s], got [%s]", "Test\n", str)
	}
}

func TestGetLogLevel(t *testing.T) {
	if GetLogLevel() != DEBUG {
		t.Errorf("Unexpected default log level. Expected [%d], got [%d]", DEBUG, GetLogLevel())
	}

	SetLogLevel(INFO)

	if GetLogLevel() != INFO {
		t.Errorf("Unexpected default log level. Expected [%d], got [%d]", INFO, GetLogLevel())
	}
}

func TestGetOutputs(t *testing.T) {
	out, err := GetOutputs()

	if out != stdout || err != stderr {
		t.Errorf("Received stdout or stderr do  not match with used in package")
	}
}

func TestGetStdout(t *testing.T) {
	out := GetStdout()

	if out != stdout {
		t.Errorf("Received stdout does not match with the one used in package")
	}
}

func TestGetStderr(t *testing.T) {
	err := GetStderr()

	if err != stderr {
		t.Errorf("Received stderr does not match with the one used in package")
	}
}
