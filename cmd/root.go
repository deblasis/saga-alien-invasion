package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "invasion [number of aliens invading]",
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.Println("DOING SOMETHING")
		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
