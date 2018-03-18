package cliassert

import (
	"bytes"
	"os/exec"
	"syscall"
)

// ExecCommand executes the provided command and returns a CmdResult
func ExecCommand(bin string, args ...string) CmdResult {
	var buffStdout, buffStderr bytes.Buffer
	code := 0

	cmd := exec.Command(bin, args...)
	cmd.Stdout = &buffStdout
	cmd.Stderr = &buffStderr

	err := cmd.Run()

	if err != nil {
		code = err.(*exec.ExitError).Sys().(syscall.WaitStatus).ExitStatus()
	}

	return CmdResult{code: code, stdout: buffStdout.String(), stderr: buffStderr.String()}
}
