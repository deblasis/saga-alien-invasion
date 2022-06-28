package app

import (
	"os"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

type World struct {
	config *Config

	Map    *Map
	Aliens map[uuid.UUID]*Alien
}

func NewWorld(config *Config) *World {
	return &World{
		config: config,
	}
}

func (w *World) Setup() error {
	err := w.parseMapFile()
	if err != nil {
		return err
	}

	//TODO: spawnAliens
	return nil
}

func (w *World) Spin() error {
	//TODO: loop
	//  TODO: handleBattles
	//  TODO: moveAliens
	//TODO: outputMap

	return nil
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