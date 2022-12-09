package main_test

import (
	"bytes"
	"github.com/dbtedman/accretion/cmd/nuptials"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRootCommandWithoutFlags(t *testing.T) {
	command := main.RootCommand()
	var errConsole bytes.Buffer
	var outConsole bytes.Buffer
	command.SetErr(&errConsole)
	command.SetOut(&outConsole)

	assert.Nil(t, command.Execute())
}
