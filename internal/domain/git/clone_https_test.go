package git_test

import (
	"github.com/dbtedman/accretion/internal/domain/git"
	"github.com/stretchr/testify/assert"
	"testing"
)

const TheCloneHTTPS = "https://github.com/dbtedman/accretion.git"
const TheCloneHTTP = "http://github.com/dbtedman/accretion.git"

func TestParseCloneHTTPSecure(t *testing.T) {
	result, resultError := git.ParseCloneHTTP(TheCloneHTTPS)

	assert.Equal(t, TheCloneHTTPS, result.String())
	assert.Nil(t, resultError)
}

func TestParseCloneHTTPInsecure(t *testing.T) {
	result, resultError := git.ParseCloneHTTP(TheCloneHTTP)

	assert.Equal(t, TheCloneHTTP, result.String())
	assert.Nil(t, resultError)
}

func TestParseCloneHTTPRejectsMissingDotGit(t *testing.T) {
	result, resultError := git.ParseCloneHTTP("https://github.com/dbtedman/accretion")

	assert.Nil(t, result)
	assert.Error(t, resultError)
}

func TestParseCloneHTTPRejectsMissingProtocol(t *testing.T) {
	result, resultError := git.ParseCloneHTTP("//github.com/dbtedman/accretion.git")

	assert.Nil(t, result)
	assert.Error(t, resultError)
}
