package git_test

import (
	"github.com/dbtedman/accretion/internal/domain/git"
	"github.com/stretchr/testify/assert"
	"testing"
)

const cloneHTTPS = "https://github.com/dbtedman/accretion.git"
const cloneSSH = "git@github.com:dbtedman/accretion.git"
const url = "https://github.com/dbtedman/accretion"

func TestGitRepositoryHasCloneHTTPS(t *testing.T) {
	repository := git.NewRepository(cloneHTTPS, cloneSSH, url)

	assert.Equal(t, cloneHTTPS, repository.CloneHTTPS())
}

func TestGitRepositoryHasCloneSSH(t *testing.T) {
	repository := git.NewRepository(cloneHTTPS, cloneSSH, url)

	assert.Equal(t, cloneSSH, repository.CloneSSH())
}

func TestGitRepositoryHasURL(t *testing.T) {
	repository := git.NewRepository(cloneHTTPS, cloneSSH, url)

	assert.Equal(t, url, repository.URL())
}
