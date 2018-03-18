package cliassert

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// CmdResult represents a command execution result
type CmdResult struct {
	code   int
	stdout string
	stderr string
}

// Stdout returns the command stdout
func (r CmdResult) Stdout() string {
	return r.stdout
}

// Stderr returns the command stderr
func (r CmdResult) Stderr() string {
	return r.stderr
}

// Code returns the command exit code
func (r CmdResult) Code() int {
	return r.code
}

// Success returns true if the command exit code is 0 and false otherwhise
func (r CmdResult) Success() bool {
	return r.code == 0
}

// AssertErrorMatch asserts that the command failed and its stderr matches the provided regexp
func (r CmdResult) AssertErrorMatch(t *testing.T, re interface{}) bool {
	if r.AssertError(t) {
		return assert.Regexp(t, re, r.stderr)
	}
	return true
}

// AssertSuccessMatch asserts that the command succeeded and its stdout matches the provided regexp
func (r CmdResult) AssertSuccessMatch(t *testing.T, re interface{}) bool {
	if r.AssertSuccess(t) {
		return assert.Regexp(t, re, r.stdout)
	}
	return true
}

// AssertCode asserts that the command  exit code matched the provided one
func (r CmdResult) AssertCode(t *testing.T, code int) bool {
	return assert.Equal(t, code, r.code, "Expected %d code but got %d", code, r.code)
}

// AssertSuccess asserts that the command succeeded (0 exit code)
func (r CmdResult) AssertSuccess(t *testing.T) bool {
	return assert.True(t, r.Success(), "Expected command to success but got code=%d stderr=%s", r.code, r.stderr)
}

// AssertError asserts that the command failed
func (r CmdResult) AssertError(t *testing.T) bool {
	return assert.False(t, r.Success(), "Expected command to fail")
}
