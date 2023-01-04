package git_test

import (
	"net/url"
	"testing"

	"github.com/dbtedman/accretion/internal/domain/git"
	"github.com/stretchr/testify/assert"
)

const theURL = "https://github.com/dbtedman/accretion"

func TestGitRepositoryHasCloneHTTPS(t *testing.T) {
	aCloneSSH, aURL := setup()

	result := git.NewRepository(TheCloneHTTPS, aCloneSSH, aURL).CloneHTTPS()

	assert.Equal(t, TheCloneHTTPS, result)
}

func TestGitRepositoryHasCloneSSH(t *testing.T) {
	aCloneSSH, aURL := setup()

	result := git.NewRepository(TheCloneHTTPS, aCloneSSH, aURL).CloneSSH()

	assert.Equal(t, aCloneSSH, result)
}

func TestGitRepositoryHasURL(t *testing.T) {
	aGitCloneSSH, aParsedURL := setup()

	result := git.NewRepository(TheCloneHTTPS, aGitCloneSSH, aParsedURL).URL()

	assert.Equal(t, aParsedURL, result)
}

func setup() (git.CloneSSH, url.URL) {
	aGitCloneSSH, _ := git.ParseCloneSSH(TheCloneSSH)
	aParsedURL, _ := url.Parse(theURL)

	return aGitCloneSSH, *aParsedURL
}
