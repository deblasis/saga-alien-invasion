package app

type City struct {
	Name string

	North *City
	East  *City
	South *City
	West  *City

	IsDestroyed bool
}

func NewCity(name string) *City {
	return &City{
		Name: name,
	}
}
