package app

// Direction points to a direction relative to a City, currently only the 4 cardinal points but extendable to anything that could be put into string format, eg: routes, coordinates, etc.
type Direction string

// instead of harcoding direction names inside *City, let's use a map keyed to this custom type that will allow extensibility as well as easier iteration
const (
	NORTH Direction = "north"
	EAST  Direction = "east"
	SOUTH Direction = "south"
	WEST  Direction = "west"
)

// AllDirection is going to be used to provide deterministic ordering, particularly necessary in tests (reason: maps are not ordered in Go)
var AllDirections = []Direction{NORTH, EAST, SOUTH, WEST}

// FlipDirection simply flips a direction for relative positioning of cities while parsing the map
func FlipDirection(dir Direction) Direction {
	var flipped Direction
	switch dir {
	case NORTH:
		flipped = SOUTH
	case EAST:
		flipped = WEST
	case SOUTH:
		flipped = NORTH
	case WEST:
		flipped = EAST
	}
	return flipped
}
