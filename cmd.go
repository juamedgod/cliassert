package cliassert

import (
	"bytes"
	"os/exec"
	"syscall"
)

// ExecCommand executes the provided command and returns a CmdResult
func ExecCommand(bin string, args ...string) CmdResult {
	cmd := NewCommand()
	return cmd.Exec(bin, args...)
}

// Command defines a command executor
type Command struct {
	stdout *bytes.Buffer
	stderr *bytes.Buffer
	stdin  string
}

// NewCommand returns a new Command
func NewCommand() *Command {
	return &Command{
		stdout: &bytes.Buffer{},
		stderr: &bytes.Buffer{},
	}
}

// SetStdin sets the stdin for the command to execute
func (c *Command) SetStdin(stdin string) {
	c.stdin = stdin
}

// Exec executes the provided command and returns a CmdResult
func (c *Command) Exec(bin string, args ...string) CmdResult {
	code := 0

	cmd := exec.Command(bin, args...)
	cmd.Stdout = c.stdout
	cmd.Stderr = c.stderr
	if c.stdin != "" {
		cmd.Stdin = bytes.NewBufferString(c.stdin)
	}
	err := cmd.Run()

	if err != nil {
		code = err.(*exec.ExitError).Sys().(syscall.WaitStatus).ExitStatus()
	}

	return CmdResult{code: code, stdout: c.stdout.String(), stderr: c.stderr.String()}
}
