/*

saga-alien-invasion simulates an alien invasion given a number of aliens and a map

The [number of aliens invading] argument is mandatory.

By default, the command expects a map.txt file in the cwd (current working directory), however it's possible to override this by providing the flag --mapfile described below.

Other than that, the command provides some extra configuration via the flags described below.

Usage:

    saga-alien-invasion [number of aliens invading] [flags]

The flags are:

		-h, --help                 help for saga-alien-invasion
		--mapfile string           the file containing the map of the 🌍, [relative path] (default "map.txt")
		--maxTurns int             the number of turns before the mothership calls the aliens back home and the program ends (default 10000)
		--numAliensForBattle int   the number aliens required to start a battle that will end up destroying a city (default 2)
		--verbose                  if True, sets the loglevel to show DEBUG messages

*/
package main

import (
	"os"

	"github.com/deblasis/saga-alien-invasion/cmd"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	cmd.Execute()
}
