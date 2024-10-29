package cmd

import (
	"github.com/dbtedman/accretion/config"
	"github.com/dbtedman/accretion/interceptor"
	"github.com/spf13/cobra"
)

func RootCommand(errorCh *chan error, proxyServer *interceptor.Proxy) *cobra.Command {
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

	cmd.AddCommand(CollectCommand(errorCh, proxyServer))
	cmd.AddCommand(VersionCommand(errorCh))

	return cmd
}
