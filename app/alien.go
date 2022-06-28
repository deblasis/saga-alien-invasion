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
}
