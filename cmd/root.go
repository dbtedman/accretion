package cmd

import (
	"github.com/dbtedman/accretion/config"
	"github.com/spf13/cobra"
)

func RootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: config.Name,
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Help()
		},
	}

	cmd.AddCommand(VersionCommand())

	return cmd
}
