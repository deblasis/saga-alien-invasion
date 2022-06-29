package app

import (
	"errors"
	"fmt"
)

var (
	ErrCityNameNotFound    = errors.New("city name not found")
	ErrConnectionsNotFound = errors.New("connections not found")
	ErrTooManyConnections  = errors.New("too many connections")
	ErrInvalidDirection    = errors.New("invalid direction")
	ErrMapFileNotFound     = fmt.Errorf("please provide a map.txt file in the current working directory or specify a mapfile path via the --%v flag", MapFileFlag)
)
