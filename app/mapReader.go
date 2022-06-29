package app

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"sort"
	"strings"
)

// using compiled regex for increased performance
var (
	//assuming that city names cannot have spaces but only dashes if multi-word
	cityNameRegex    = regexp.MustCompile(`^([\w\-]+)(?:\s)|(^[\w\-]+)$`)
	connectionsRegex = regexp.MustCompile(fmt.Sprintf(`(%v|%v|%v|%v)\s*=([\w|-]+)`, NORTH, EAST, SOUTH, WEST))
)

// MapReader takes care of parsing a mapfile
type MapReader struct {
	cities map[string]*City
}

// NewMapReader returns a mapReader instance
func NewMapReader() *MapReader {
	return &MapReader{
		cities: make(map[string]*City),
	}
}

// ParseMapFile scans a mapfile for valid cities and returns a [Map] or an error accordingly
func (mr *MapReader) ParseMapFile(reader io.Reader) (*Map, error) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		line := scanner.Text()
		err := mr.parseCity(line)
		if err != nil {
			return nil, err
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	// this is because ordering in maps is not preserved in Go for security reasons
	// we use a slice to provide ordering and here we initialize it
	keys := make([]string, 0)
	for k := range mr.cities {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	return &Map{
		Cities:          mr.cities,
		sortedCityNames: keys,
	}, nil
}

func (mr *MapReader) parseCity(line string) error {
	if line == "" {
		return nil
	}
	cityMatch := cityNameRegex.FindAllStringSubmatch(line, -1)
	connectionMatches := connectionsRegex.FindAllString(line, -1)

	if len(cityMatch) == 0 || len(cityMatch[0]) < 2 || (cityMatch[0][1] == "" && cityMatch[0][0] == "") {
		return fmt.Errorf("error processing line %v - %w", line, ErrCityNameNotFound)
	}
	// I am assuming that cities without connections are allowed
	// if len(connectionMatches) == 0 {
	// 	return ErrConnectionsNotFound
	// }
	if len(connectionMatches) > 4 {
		return ErrTooManyConnections
	}

	cityName := cityMatch[0][1]
	if cityName == "" {
		cityName = cityMatch[0][0]
	}
	city := mr.upsertCity(cityName)

	for _, conn := range connectionMatches {
		kv := strings.Split(conn, "=")
		//assuming we match exactly and that the directions are always lowercase
		direction := kv[0]
		destination := kv[1]

		var destCity = mr.upsertCity(destination)

		switch direction {
		case string(NORTH):
			city.Directions[NORTH] = destCity
		case string(EAST):
			city.Directions[EAST] = destCity
		case string(SOUTH):
			city.Directions[SOUTH] = destCity
		case string(WEST):
			city.Directions[WEST] = destCity
		default:
			return ErrInvalidDirection
		}
		destCity.Directions[FlipDirection(Direction(direction))] = city
	}

	return nil
}

func (mr *MapReader) upsertCity(cityMatch string) *City {
	if mr.cities[cityMatch] == nil {
		mr.cities[cityMatch] = NewCity(cityMatch)
	}
	return mr.cities[cityMatch]
}
