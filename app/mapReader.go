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

	//TODO: scan line by line
	//	TODO: parseCity (regex?)

	//TODO: ensure determinism

	return &Map{
		Cities: mr.Cities,
	}, nil
}
