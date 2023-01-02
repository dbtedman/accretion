package git_test

import (
	"github.com/dbtedman/accretion/internal/domain/git"
	"github.com/stretchr/testify/assert"
	"testing"
)

const ssh = "git@github.com:dbtedman/accretion.git"

func TestSSHToString(t *testing.T) {
	assert.Equal(t, ssh, git.NewSSH(ssh).ToString())
}
