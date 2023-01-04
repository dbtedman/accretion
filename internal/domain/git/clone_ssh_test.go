package git_test

import (
	"testing"

	"github.com/dbtedman/accretion/internal/domain/git"
	"github.com/stretchr/testify/assert"
)

const TheCloneSSH = "git@github.com:dbtedman/accretion.git"

func TestSSHToString(t *testing.T) {
	result, resultError := git.ParseCloneSSH(TheCloneSSH)

	assert.Equal(t, TheCloneSSH, result.String())
	assert.Nil(t, resultError)
}
