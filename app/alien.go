package app

import (
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

type Alien struct {
	// Id probably redundant but I don't want to rely on the unicity of names provided externally in NewAlien() that could lead to collisions
	Id       uuid.UUID
	Name     string
	Location *City
}

func NewAlien(name string) *Alien {
	return &Alien{
		Id:   uuid.New(),
		Name: name,
	}
}

func (a *Alien) Move() {
	logg := log.With().Str("component", "Alien.Move()").Str("alien", a.Name).Str("location", a.Location.Name).Logger()
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
		logg.Debug().Msg("nowhere to go")
		return
	}

	tIdx := Random.Intn(n)
	a.Location = loc.Directions[choices[tIdx]]
	logg.Debug().Str("destination", a.Location.Name).Msg("moving")
}
