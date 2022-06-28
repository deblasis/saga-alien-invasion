package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "invasion [number of aliens invading]",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("the number of aliens invading the world must be specified")
		}
		var n int
		n, err := strconv.Atoi(args[0])
		if err != nil || n <= 0 {
			return fmt.Errorf("the number of aliens must be *checks notes* a positive number! received: %s", args[0])
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {

		aliensCount, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}
		cmd.Printf("DOING SOMETHING with aliensCount: %d\n", aliensCount)
		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
