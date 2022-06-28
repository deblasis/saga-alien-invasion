package app

import (
	"math/rand"
	"time"
)

// init ensures pseudo-randomness
func init() {
	rand.Seed(time.Now().UnixNano())
}

// Rand abstracts the randomizer for the method(s) we'll use so that we can switch implementation in tests with a deterministic one
type Rand interface {
	Intn(n int) int
}

type RealRandomizer struct{}

func NewRealRandomizer() Rand {
	return &RealRandomizer{}
}

// Intn simply wraps the real implementation
func (*RealRandomizer) Intn(n int) int {
	return rand.Intn(n)
}
