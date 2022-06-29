package app

import (
	"fmt"
	"io"
	"strings"

	"github.com/rs/zerolog/log"
)

// MapWriter outputs the map to an [io.Writer]
type MapWriter struct {
	m *Map
}

// NewMapWriter returns an instance of a MapWriter
func NewMapWriter(m *Map) *MapWriter {
	return &MapWriter{
		m: m,
	}
}

// WriteMap outputs the map in text format
func (m *MapWriter) WriteMap(writer io.Writer) error {
	logg := log.With().Str("component", "MapWriter.WriteMap()").Logger()
	logg.Debug().Msg("executing")
	if len(m.m.sortedCityNames) == 0 {
		writer.Write([]byte("COMPLETE ANNIHILATION - All cities have been destroyed ☠️\n"))
		return nil
	}
	for _, cityName := range m.m.sortedCityNames {
		connections := make([]string, 0)
		city := m.m.Cities[cityName]
		if city == nil {
			continue
		}

		for _, dir := range AllDirections {
			c := city.Directions[dir]
			if c == nil {
				continue
			}
			connections = append(connections, fmt.Sprintf("%v=%v", string(dir), city.Directions[dir].Name))
		}
		if len(connections) == 0 {
			writer.Write([]byte(cityName + "\n"))
		} else {
			writer.Write([]byte(fmt.Sprintf("%v %v\n", cityName, strings.Join(connections, " "))))
		}
	}
	return nil
}
