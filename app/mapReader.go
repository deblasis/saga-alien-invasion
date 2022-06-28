package app

import "io"

type mapReader struct {
	Cities map[string]*City
}

func NewMapReader() *mapReader {
	return &mapReader{
		Cities: make(map[string]*City),
	}
}

func (mr *mapReader) ParseMapFile(reader io.Reader) (*Map, error) {
	return &Map{
		Cities: mr.Cities,
	}, nil
}
