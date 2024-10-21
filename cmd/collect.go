package cmd

import (
	"github.com/dbtedman/accretion/config"
	"github.com/spf13/cobra"
)

func CollectCommand(errorCh *chan error) *cobra.Command {
	var listenAddress string
	var proxyAddress string

	cmd := &cobra.Command{
		Use: "collect",
		Run: func(cmd *cobra.Command, args []string) {},
	}

	cmd.PersistentFlags().StringVar(&listenAddress, "from", config.DefaultListenAddress, "")
	cmd.PersistentFlags().StringVar(&proxyAddress, "to", config.DefaultProxyAddress, "")

	return cmd
}
