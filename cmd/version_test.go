package cmd_test

import (
	"bytes"
	"github.com/dbtedman/accretion/cmd"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVersionCommand(t *testing.T) {
	errorCh := make(chan error)
	var errConsole bytes.Buffer
	var outConsole bytes.Buffer
	command := cmd.VersionCommand(&errorCh)
	command.SetErr(&errConsole)
	command.SetOut(&outConsole)

	go func() {
		errorCh <- command.Execute()
	}()
	err := <-errorCh

	assert.Nil(t, err)
	assert.Equal(t, "", errConsole.String())
	assert.Equal(t, "", outConsole.String())
}
