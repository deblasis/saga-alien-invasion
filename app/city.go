package app

type City struct {
	Name string

	Directions   map[Direction]*City
	AliensInTown []*Alien

	IsDestroyed bool
}

func NewCity(name string) *City {
	return &City{
		Name:       name,
		Directions: make(map[Direction]*City),
	}
}
