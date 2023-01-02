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
	gitCloneSSH, _ := git.NewSSH(cloneSSH)
	repository := git.NewRepository(cloneHTTPS, gitCloneSSH, url)

	assert.Equal(t, cloneHTTPS, repository.CloneHTTPS())
}

func TestGitRepositoryHasCloneSSH(t *testing.T) {
	gitCloneSSH, _ := git.NewSSH(cloneSSH)
	repository := git.NewRepository(cloneHTTPS, gitCloneSSH, url)

	assert.Equal(t, cloneSSH, repository.CloneSSH().ToString())
	assert.Implements(t, (*git.SSH)(nil), repository.CloneSSH())
}

func TestGitRepositoryHasURL(t *testing.T) {
	gitCloneSSH, _ := git.NewSSH(cloneSSH)
	repository := git.NewRepository(cloneHTTPS, gitCloneSSH, url)

	assert.Equal(t, url, repository.URL())
}
