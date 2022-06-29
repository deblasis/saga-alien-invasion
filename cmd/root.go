// Package cmd implements the CLI commands, args and flag parsing and generally speaking the entrypoint of the program
package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/deblasis/saga-alien-invasion/app"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.PersistentFlags().String(app.MapFileFlag, "map.txt", "the file containing the map of the 🌍, [relative path]")
	rootCmd.PersistentFlags().Int(app.MaxTurnsFlag, 10000, "the number of turns before the mothership calls the aliens back home and the program ends")
	rootCmd.PersistentFlags().Int(app.NumAliensForBattleFlag, 2, "the number aliens required to start a battle that will end up destroying a city")
	rootCmd.PersistentFlags().Bool(app.VerboseFlag, false, "if True, sets the loglevel to show DEBUG messages")

	//this is to avoid casting when lookingup the flag values
	viper.BindPFlag(app.MapFileFlag, rootCmd.PersistentFlags().Lookup(app.MapFileFlag))
	viper.BindPFlag(app.MaxTurnsFlag, rootCmd.PersistentFlags().Lookup(app.MaxTurnsFlag))
	viper.BindPFlag(app.NumAliensForBattleFlag, rootCmd.PersistentFlags().Lookup(app.NumAliensForBattleFlag))
	viper.BindPFlag(app.VerboseFlag, rootCmd.PersistentFlags().Lookup(app.VerboseFlag))
}

var rootCmd = &cobra.Command{
	Use: "saga-alien-invasion [number of aliens invading]",
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

		isVerbose := viper.GetBool(app.VerboseFlag)
		if isVerbose {
			zerolog.SetGlobalLevel(zerolog.DebugLevel)
		}

		aliensCount, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}

		config := &app.Config{
			AliensCount:        aliensCount,
			MapfilePath:        viper.GetString(app.MapFileFlag),
			MaxTurns:           viper.GetInt(app.MaxTurnsFlag),
			NumAliensForBattle: viper.GetInt(app.NumAliensForBattleFlag),
			Verbose:            isVerbose,
		}

		world := app.NewWorld(cmd, config)

		err = world.Setup()
		if err != nil {
			return err
		}

		printHeader(cmd, config)

		err = world.Spin()
		if err != nil {
			return err
		}
		printSummary(cmd, world)

		return nil
	},
}

// Execute executes the command logic
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func printHeader(cmd *cobra.Command, config *app.Config) {
	cmd.Println(app.Separator)
	cmd.Printf("The alien invasion begins... %d aliens are approaching\n", config.AliensCount)
}

func printSummary(cmd *cobra.Command, w *app.World) {
	cmd.Println(app.Separator)
	cmd.Printf("The 🌍 after %v day(s) of alien invasion:\n", w.CurrentDay+1)
	cmd.Println(app.Separator)
	printMap(cmd, w.Map)
	cmd.Println(app.Separator)
}
func printMap(cmd *cobra.Command, m *app.Map) {
	mw := app.NewMapWriter(m)
	mw.WriteMap(cmd.OutOrStdout())
}
