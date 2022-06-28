package app

import (
	"github.com/rs/zerolog/log"
)

type alienSpawner struct {
	NumAliens      int
	World          *World
	AlienLocations map[*Alien]string

	cities []string
}

func NewAlienSpawner(numAliens int, w *World) *alienSpawner {
	return &alienSpawner{
		NumAliens: numAliens,
		World:     w,
		cities:    w.Map.sortedCityNames,
	}
}

func (a *alienSpawner) Spawn() error {
	logg := log.With().Str("component", "alienSpawner.Spawn()").Logger()

	namer := NewAlienOrdinalNamer()

	for i := 0; i < a.NumAliens; i++ {
		alien := NewAlien(namer.NameAlien())

		n := len(a.World.Map.Cities)
		tIdx := Random.Intn(n)
		tCity := a.cities[tIdx]
		target := a.World.Map.Cities[tCity]
		a.World.Aliens[alien.Id] = alien
		alien.Location = target

		logg.Debug().Str("randomCity", target.Name).Str("alien", alien.Name).Msg("Alien spawned!")
	}

	return nil
}
