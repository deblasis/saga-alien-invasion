package app

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"regexp"
	"sort"
	"strings"
)

// using compiled regex for increased performance
var (
	cityNameRegex    = regexp.MustCompile(`^([\w\-]+)`)
	connectionsRegex = regexp.MustCompile(fmt.Sprintf(`(%v|%v|%v|%v)\s*=([\w|-]+)`, NORTH, EAST, SOUTH, WEST))
)

type mapReader struct {
	Cities map[string]*City
}

func NewMapReader() *mapReader {
	return &mapReader{
		Cities: make(map[string]*City),
	}
}

func (mr *mapReader) ParseMapFile(reader io.Reader) (*Map, error) {
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
	for k, _ := range mr.Cities {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	return &Map{
		Cities:          mr.Cities,
		sortedCityNames: keys,
	}, nil
}

func (mr *mapReader) parseCity(line string) error {
	if line == "" {
		return nil
	}
	cityMatch := cityNameRegex.FindString(line)
	connectionMatches := connectionsRegex.FindAllString(line, -1)

	if cityMatch == "" {
		return fmt.Errorf("error processing line %v - %w", line, ErrCityNameNotFound)
	}
	if len(connectionMatches) == 0 {
		return ErrConnectionsNotFound
	}
	if len(connectionMatches) > 4 {
		return ErrTooManyConnections
	}

	city := mr.upsertCity(cityMatch)

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

func (mr *mapReader) upsertCity(cityMatch string) *City {
	if mr.Cities[cityMatch] == nil {
		mr.Cities[cityMatch] = NewCity(cityMatch)
	}
	return mr.Cities[cityMatch]
}

var (
	ErrCityNameNotFound    = errors.New("city name not found")
	ErrConnectionsNotFound = errors.New("connections not found")
	ErrTooManyConnections  = errors.New("too many connections")
	ErrInvalidDirection    = errors.New("invalid direction")
)
