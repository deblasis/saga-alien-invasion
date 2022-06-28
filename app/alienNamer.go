package app

import "fmt"

type AlienNamer interface {
	NameAlien() string
}

type ordinalNamer struct {
	seq int
}

func NewAlienOrdinalNamer() AlienNamer {
	return &ordinalNamer{
		seq: 0,
	}
}

func (n *ordinalNamer) NameAlien() string {
	defer func() { n.seq++ }()
	return fmt.Sprintf("Alien %d", n.seq)
}
