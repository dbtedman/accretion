package cmd_test

import (
	"bytes"
	"fmt"
	"github.com/dbtedman/accretion/cmd"
	"github.com/dbtedman/accretion/config"
	"github.com/dbtedman/accretion/interceptor"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRootCommand(t *testing.T) {
	errorCh := make(chan error)
	var errConsole bytes.Buffer
	var outConsole bytes.Buffer
	var proxyServer interceptor.Proxy = &interceptor.TestProxy{}
	command := cmd.RootCommand(&errorCh, &proxyServer)
	command.SetErr(&errConsole)
	command.SetOut(&outConsole)

	go func() {
		errorCh <- command.Execute()
	}()
	err := <-errorCh

	assert.Nil(t, err)
	assert.Equal(t, "", errConsole.String())
	assert.Contains(t, outConsole.String(), config.Purpose)
	assert.Contains(t, outConsole.String(), fmt.Sprintf("-h, --help   help for %s", config.Name))
}
