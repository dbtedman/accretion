package main

import (
	"github.com/charmbracelet/log"
	"github.com/dbtedman/accretion/cmd"
	"github.com/dbtedman/accretion/config"
)

// version, commit, and date are populated during build
var (
	version = "latest"
	commit  = "n/a"
	date    = "n/a"
)

func main() {
	log.Info(config.Name)
	_ = cmd.RootCommand().Execute()
}
