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

const (
	mapFileFlag            = "mapfile"
	maxTurnsFlag           = "maxTurns"
	verboseFlag            = "verbose"
	numAliensForBattleFlag = "numAliensForBattle"
)

func init() {
	rootCmd.PersistentFlags().String(mapFileFlag, "map.txt", "the file containing the map of the 🌍, [relative path]")
	rootCmd.PersistentFlags().Int(maxTurnsFlag, 10000, "the number of turns before the mothership calls the aliens back home and the program ends")
	rootCmd.PersistentFlags().Int(numAliensForBattleFlag, 2, "the number aliens required to start a battle that will end up destroying a city")
	rootCmd.PersistentFlags().Bool(verboseFlag, false, "if True, sets the loglevel to show DEBUG messages")

	//this is to avoid casting when lookingup the flag values
	viper.BindPFlag(mapFileFlag, rootCmd.PersistentFlags().Lookup(mapFileFlag))
	viper.BindPFlag(maxTurnsFlag, rootCmd.PersistentFlags().Lookup(maxTurnsFlag))
	viper.BindPFlag(numAliensForBattleFlag, rootCmd.PersistentFlags().Lookup(numAliensForBattleFlag))
	viper.BindPFlag(verboseFlag, rootCmd.PersistentFlags().Lookup(verboseFlag))
}

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

		isVerbose := viper.GetBool(verboseFlag)
		if isVerbose {
			zerolog.SetGlobalLevel(zerolog.DebugLevel)
		}

		aliensCount, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}

		world := app.NewWorld(cmd, &app.Config{
			AliensCount:        aliensCount,
			MapfilePath:        viper.GetString(mapFileFlag),
			MaxTurns:           viper.GetInt(maxTurnsFlag),
			NumAliensForBattle: viper.GetInt(numAliensForBattleFlag),
			Verbose:            isVerbose,
		})

		err = world.Setup()
		if err != nil {
			return err
		}

		err = world.Spin()
		if err != nil {
			return err
		}

		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
