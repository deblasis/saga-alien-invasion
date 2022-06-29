package app

import (
	"github.com/rs/zerolog/log"
)

// Mothership decides where to deploy aliens and also where they should move when the invasion started
type Mothership struct {
	numAliens      int
	world          *World
	alienLocations map[*Alien]string
	cities         []string
}

// NewMothership returns a Mothership instance
func NewMothership(numAliens int, w *World) *Mothership {
	return &Mothership{
		numAliens: numAliens,
		world:     w,
		cities:    w.Map.sortedCityNames,
	}
}

// DeployAlien recruits, names and deploys aliens in one of the cities
func (m *Mothership) DeployAlien() error {
	logg := log.With().Str("component", "mothership.DeployAlien()").Logger()

	// alternative implementations
	//namer := NewAlienOrdinalNamer()
	namer := NewFamousAliensNamer()

	for i := 0; i < m.numAliens; i++ {
		alien := NewAlien(namer.NameAlien(), m)

		target := m.getSpaceTargetCity()

		target.AliensInTown = append(target.AliensInTown, alien)
		m.world.Aliens[alien.Id] = alien
		alien.Location = target

		logg.Debug().Str("targetCity", target.Name).Str("alien", alien.Name).Msg("Alien deployed!")
	}

	return nil
}

// getSpaceTargetCity encapsulates the logic that selects the next target when aliens can be deployed anywhere on the map
func (m *Mothership) getSpaceTargetCity() *City {
	logg := log.With().Str("component", "mothership.getSpaceTargetCity()").Logger()
	logg.Debug().Msg("executing")
	n := len(m.world.Map.Cities)
	tIdx := Random.Intn(n)
	tCity := m.cities[tIdx]
	target := m.world.Map.Cities[tCity]
	logg.Debug().Msgf("")
	return target
}

func (m *Mothership) getLandTargetCity(a *Alien) *City {
	logg := log.With().Str("component", "mothership.getLandTargetCity()").Str("alien", a.Name).Str("location", a.Location.Name).Logger()
	logg.Debug().Msg("executing")
	loc := a.Location
	choices := make([]Direction, 0)
	for _, dir := range AllDirections {
		city := loc.Directions[dir]
		if city != nil {
			choices = append(choices, dir)
		}
	}

	n := len(choices)
	if n == 0 {
		logg.Debug().Msg("the alien has nowhere else to go")
		return nil
	}

	tIdx := Random.Intn(len(choices))
	return a.Location.Directions[choices[tIdx]]
}
