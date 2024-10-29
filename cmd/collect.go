package cmd

import (
	"github.com/dbtedman/accretion/config"
	"github.com/dbtedman/accretion/interceptor"
	"github.com/spf13/cobra"
	"net/url"
)

func CollectCommand(errorCh *chan error, proxyServer *interceptor.Proxy) *cobra.Command {
	var listenAddress string
	var proxyAddress string

	cmd := &cobra.Command{
		Use: "collect",
		Run: func(cmd *cobra.Command, args []string) {
			proxyAddressURL, _ := url.Parse(proxyAddress)
			interceptor.ListenHTTPWithProxy(*proxyServer, *proxyAddressURL, listenAddress, errorCh)
		},
	}

	cmd.PersistentFlags().StringVar(&listenAddress, "from", config.DefaultListenAddress, "")
	cmd.PersistentFlags().StringVar(&proxyAddress, "to", config.DefaultProxyAddress, "")

	return cmd
}
