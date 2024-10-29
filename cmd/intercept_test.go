package cmd_test

import (
	"bytes"
	"github.com/dbtedman/accretion/cmd"
	"github.com/dbtedman/accretion/interceptor"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCollectCommand(t *testing.T) {
	errorCh := make(chan error)
	var errConsole bytes.Buffer
	var outConsole bytes.Buffer
	var proxyServer interceptor.Proxy = &interceptor.TestProxy{}
	command := cmd.InterceptCommand(&errorCh, &proxyServer)
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
