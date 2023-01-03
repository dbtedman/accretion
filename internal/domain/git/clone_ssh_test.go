package git_test

import (
	"github.com/dbtedman/accretion/internal/domain/git"
	"github.com/stretchr/testify/assert"
	"testing"
)

const TheCloneSSH = "git@github.com:dbtedman/accretion.git"

func TestSSHToString(t *testing.T) {
	result, resultError := git.ParseCloneSSH(TheCloneSSH)

	assert.Equal(t, TheCloneSSH, result.String())
	assert.Nil(t, resultError)
}
