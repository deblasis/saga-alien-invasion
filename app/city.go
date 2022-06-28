package app

type City struct {
	Name string

	Directions map[Direction]*City

	IsDestroyed bool
}

func NewCity(name string) *City {
	return &City{
		Name:       name,
		Directions: make(map[Direction]*City),
	}
}
