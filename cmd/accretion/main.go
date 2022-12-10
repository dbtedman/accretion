package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

const commandError = 1

func main() {
	if err := RootCommand().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(commandError)
	}
}

func RootCommand() *cobra.Command {
	rootCommand := &cobra.Command{
		Use:   "accretion",
		Short: "Manage internal technical documentation that is enriched with live data accreted from your environment.",
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Help()
		},
	}

	return rootCommand
}
