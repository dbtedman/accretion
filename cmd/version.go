package cmd

import (
	"github.com/charmbracelet/log"
	"github.com/dbtedman/accretion/config"
	"github.com/spf13/cobra"
)

// version, commit, and date are populated during build
var (
	version = "latest"
	commit  = "n/a"
	date    = "n/a"
)

func VersionCommand(errorCh *chan error) *cobra.Command {
	return &cobra.Command{
		Use: "version",
		Run: func(cmd *cobra.Command, args []string) {
			log.Info(config.Name, "version", version, "commit", commit, "date", date)
			*errorCh <- nil
		},
	}
}
