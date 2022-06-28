package app

import (
	"fmt"
	"io"
	"strings"

	"github.com/rs/zerolog/log"
)

type MapWriter struct {
	Map *Map
}

func NewMapWriter(m *Map) *MapWriter {
	return &MapWriter{
		Map: m,
	}
}

func (m *MapWriter) WriteMap(writer io.Writer) error {
	logg := log.With().Str("component", "MapWriter.WriteMap()").Logger()
	logg.Debug().Msg("executing")
	if len(m.Map.sortedCityNames) == 0 {
		writer.Write([]byte("COMPLETE ANNIHILATION - All cities have been destroyed ☠️\n"))
		return nil
	}
	for _, cityName := range m.Map.sortedCityNames {
		connections := make([]string, 0)
		city := m.Map.Cities[cityName]
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
