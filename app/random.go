package app

import (
	"math/rand"
	"sync"
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

// Random holds a reference to the randomizer currently in use, by default we use the real one (weak one...) math/rand and we also seed it in init().
// During tests we override this with another randomizer that implements the interface [Rand]
var Random = NewRealRandomizer()

type RealRandomizer struct{}

// NewRealRandomizer instantiates a new RealRandomizer
func NewRealRandomizer() Rand {
	return &RealRandomizer{}
}

// Intn simply wraps the real implementation
func (*RealRandomizer) Intn(n int) int {
	return rand.Intn(n)
}

// FakeRandomizer holds a sequence of numbers that's initialized every time we call Intn, the argument is used as key to keep track of the current value in the sequence and so subsequent calls should provide different but predictable values
// it's a naive implementation because it returns sequential numbers. Perhaps I should consider implementing some sort of in-memory event bus so that I would test event sequences f(events) => state
// instead of trying to recreate these events from the inputs... which are random... I'll look into it
type FakeRandomizer struct {
	sync.RWMutex
	seq  map[int][]int
	cont map[int]int
}

// NewFakeRandomizer instantiates a new FakeRandomizer
func NewFakeRandomizer() Rand {
	return &FakeRandomizer{
		RWMutex: sync.RWMutex{},
		seq:     map[int][]int{},
		cont:    map[int]int{},
	}
}

// Intn tries to return "predictable randomness" in sequences scoped by their range numner (n int) argument
func (f *FakeRandomizer) Intn(n int) int {

	f.Lock()
	defer f.Unlock()

	if f.seq[n] == nil {
		f.seq[n] = make([]int, 0)
	}

	if len(f.seq[n]) == 0 {
		for i := 0; i < n; i++ {
			f.seq[n] = append(f.seq[n], i)
		}
	}

	if f.cont[n] >= len(f.seq[n])-1 {
		f.cont[n] = -1
	}
	f.cont[n]++

	return f.seq[n][f.cont[n]]
}
