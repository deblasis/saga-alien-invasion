package app

import (
	"os"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

type World struct {
	config *Config

	Map        *Map
	Aliens     map[uuid.UUID]*Alien
	CurrentDay int

	cmd *cobra.Command
}

func NewWorld(cmd *cobra.Command, config *Config) *World {
	return &World{
		config: config,
		Aliens: make(map[uuid.UUID]*Alien),
		cmd:    cmd,
	}
}

func (w *World) Setup() error {
	err := w.parseMapFile()
	if err != nil {
		return err
	}

	err = w.spawnAliens()
	if err != nil {
		return err
	}

	return nil
}

func (w *World) Spin() error {
	logg := log.With().Str("component", "World.Spin()").Logger()
	logg.Debug().Int("maxTurns", w.config.MaxTurns).Msg("The 🌍 starts spinning...")

	w.printHeader()

	for day := 0; day < w.config.MaxTurns; day++ {
		w.CurrentDay = day

		if len(w.Aliens) == 0 {
			logg.Debug().Msg("All aliens are dead")
			break
		}
		//this is implicit in the rules, I am assuming that the game can end early if battles cannot take place anymore
		if len(w.Aliens) < w.config.NumAliensForBattle {
			w.cmd.Println("The alien force is too weak, the invasion failed")
			logg.Debug().Int("aliensLeft", len(w.Aliens)).Int("numAliensForBattle", w.config.NumAliensForBattle).Msg("Not enough aliens left to start a fight, humanity is saved... for now!")
			break
		}
		//possibly redundant but it codifies a rule for ending the game
		if len(w.Map.Cities) == 0 {
			logg.Debug().Msg("All cities have been destroyed")
			break
		}
		//  TODO: handleBattles
		//  TODO: moveAliens
		//TODO: outputMap

	}

	logg.Debug().Msg("DONE!")

	return nil
}

func (w *World) printHeader() {
	w.cmd.Println("--------------------------------------------------------")
	w.cmd.Println("The alien invasion begins...")
	w.cmd.Println("--------------------------------------------------------")
}

func (w *World) parseMapFile() error {
	logg := log.With().Str("component", "World.parseMapFile()").Str("mapfile", w.config.MapfilePath).Logger()
	logg.Debug().Msg("executing")
	if _, err := os.Stat(w.config.MapfilePath); err != nil {
		return err
	}
	mr := NewMapReader()
	r, err := os.Open(w.config.MapfilePath)
	if err != nil {
		return err
	}
	defer r.Close()
	m, err := mr.ParseMapFile(r)
	if err != nil {
		return err
	}
	w.Map = m
	logg.Debug().Msgf("Map %v", m)
	return nil
}

func (w *World) spawnAliens() error {
	logg := log.With().Str("component", "World.spawnAliens()").Int("aliensCount", w.config.AliensCount).Logger()
	logg.Debug().Msg("executing")
	as := NewAlienSpawner(w.config.AliensCount, w)
	return as.Spawn()
}
