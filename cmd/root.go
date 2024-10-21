package cmd

import (
	"github.com/dbtedman/accretion/config"
	"github.com/spf13/cobra"
)

func RootCommand(errorCh *chan error) *cobra.Command {
	cmd := &cobra.Command{
		Use:   config.Name,
		Short: config.Purpose,
		Run: func(cmd *cobra.Command, args []string) {
			err := cmd.Help()

			if err != nil {
				*errorCh <- err
			} else {
				*errorCh <- nil
			}
		},
	}

	cmd.AddCommand(CollectCommand(errorCh))
	cmd.AddCommand(VersionCommand(errorCh))

	return cmd
}
