package main

import "log"
import "github.com/spf13/cobra"

// version, commit, and date are populated during build
var (
	version = "latest"
	commit  = "n/a"
	date    = "n/a"
)

func main() {
	rootCommand := &cobra.Command{
		Use: "accretion",
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Help()
		},
	}

	rootCommand.AddCommand(
		&cobra.Command{
			Use: "version",
			Run: func(cmd *cobra.Command, args []string) {
				log.Printf("scrutinise version: %s, commit: %s, built at: %s", version, commit, date)
			},
		},
	)

	_ = rootCommand.Execute()
}
