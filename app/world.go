package app

import (
	"fmt"
	"os"
	"strings"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// World is the "container object" that holds basically the state of the... World
type World struct {
	config *Config

	Map        *Map
	Aliens     map[uuid.UUID]*Alien
	CurrentDay int

	cmd *cobra.Command
}

// NewWorld returns an instance of World given a [cobra.Command] (so that we can capture/redirect stdout and stderr) and [Config]
func NewWorld(cmd *cobra.Command, config *Config) *World {
	return &World{
		config: config,
		Aliens: make(map[uuid.UUID]*Alien),
		cmd:    cmd,
	}
}

// Setup initializes the World
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

// Spin starts the loop that can be related to the passing of time (turns/days are assumed to be the same unit of measure here)
// inside this loop things happen, such as battles between aliens, cities being destroyed, aliens moving
func (w *World) Spin() error {
	logg := log.With().Str("component", "World.Spin()").Logger()
	logg.Debug().Int("maxTurns", w.config.MaxTurns).Msg("The 🌍 starts spinning...")

	for day := 0; day < w.config.MaxTurns; day++ {
		w.CurrentDay = day

		if err := w.handleBattles(); err != nil {
			return err
		}

		if len(w.Aliens) == 0 {
			logg.Debug().Msg("All aliens are dead")
			break
		}
		//this is implicit in the rules, I am assuming that the game can end early if battles cannot take place anymore
		if len(w.Aliens) < w.config.NumAliensForBattle {
			w.cmd.Println(Separator)
			w.cmd.Println("The alien force is too weak, the invasion failed")
			logg.Debug().Int("aliensLeft", len(w.Aliens)).Int("numAliensForBattle", w.config.NumAliensForBattle).Msg("Not enough aliens left to start a fight, humanity is saved... for now!")
			break
		}
		//possibly redundant but it codifies a rule for ending the game
		if len(w.Map.Cities) == 0 {
			logg.Debug().Msg("All cities have been destroyed")
			break
		}
		//logg.Debug().Int("day", day).Msg("Executing day")

		if err := w.moveAliens(); err != nil {
			return err
		}
	}

	logg.Debug().Msg("DONE!")

	return nil
}

func (w *World) parseMapFile() error {
	logg := log.With().Str("component", "World.parseMapFile()").Str("mapfile", w.config.MapfilePath).Logger()
	logg.Debug().Msg("executing")
	if _, err := os.Stat(w.config.MapfilePath); err != nil {
		return fmt.Errorf("%v not found - %w", w.config.MapfilePath, ErrMapFileNotFound)
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
	as := NewMothership(w.config.AliensCount, w)
	return as.DeployAlien()
}

func (w *World) handleBattles() error {
	logg := log.With().Str("component", "World.handleBattles()").Int("day", w.CurrentDay).Logger()
	logg.Debug().Msg("executing")

	destroyedCities := make([]string, 0)

	for _, cityName := range w.Map.sortedCityNames {
		city := w.Map.Cities[cityName]
		if city == nil {
			logg.Debug().Str("city", cityName).Msg("🔥 already destroyed 🔥")
			continue
		}
		if len(city.AliensInTown) >= w.config.NumAliensForBattle {
			deadAliens := make([]string, 0)
			destroyedCities = append(destroyedCities, cityName)
			for _, alien := range city.AliensInTown {
				delete(w.Aliens, alien.Id)
				deadAliens = append(deadAliens, alien.Name)
			}
			logg.Debug().Str("city", cityName).Str("deadAliens", strings.Join(deadAliens, ",")).Msg("🔥 just destroyed 🔥")
			w.cmd.Printf("%v has been destroyed by %v!\n", cityName, strings.Join(deadAliens, " and "))
		}
	}

	for _, c := range destroyedCities {
		w.destroyCity(c)
	}

	return nil
}

func (w *World) destroyCity(cityName string) {
	delete(w.Map.Cities, cityName)
	var idx int
	for i, c := range w.Map.sortedCityNames {
		if cityName == c {
			idx = i
			break
		}
	}
	w.Map.sortedCityNames = append(w.Map.sortedCityNames[:idx], w.Map.sortedCityNames[idx+1:]...)
}

func (w *World) moveAliens() error {
	logg := log.With().Str("component", "World.moveAliens()").Int("day", w.CurrentDay).Logger()
	logg.Debug().Msg("executing")

	for _, alien := range w.Aliens {
		if alien == nil {
			//logg.Debug().Str("alienId", id.String()).Msg("is dead")
			continue
		}
		alien.Move()
	}
	return nil
}
