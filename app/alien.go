package app

import (
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

// Alien is the entity that can be deployed at a Location [City] and that receives orders from [Mothership]
type Alien struct {
	// Id probably redundant but I don't want to rely on the unicity of names provided externally in NewAlien() that could lead to collisions
	Id       uuid.UUID
	Name     string
	Location *City

	m *Mothership
}

func NewAlien(name string, m *Mothership) *Alien {
	return &Alien{
		Id:   uuid.New(),
		Name: name,
		m:    m,
	}
}

// Move makes the Alien execute the move as dictated by Mothership (my assumption since I am envisaging that Mothership has better knowledge about the invasion strategy... also this is for Dependency Injection and reusability. In my world, Alien is dumb.)
func (a *Alien) Move() {
	logg := log.With().Str("component", "Alien.Move()").Str("alien", a.Name).Str("location", a.Location.Name).Logger()
	logg.Debug().Msg("executing")

	// this is the alien asking the mothership where to move next since they know better (my assumption)
	loc := a.m.getLandTargetCity(a)
	if loc == nil {
		return
	}
	a.Location = loc
	logg.Debug().Str("destination", a.Location.Name).Msg("moving")
}
