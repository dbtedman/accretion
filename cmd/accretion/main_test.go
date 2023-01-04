package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRootCommandWithoutFlags(t *testing.T) {
	command := RootCommand()
	var errConsole bytes.Buffer
	var outConsole bytes.Buffer
	command.SetErr(&errConsole)
	command.SetOut(&outConsole)

	assert.Nil(t, command.Execute())
}
