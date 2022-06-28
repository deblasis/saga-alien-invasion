package app

import (
	"reflect"
	"testing"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	log.Logger = zerolog.Logger{}
}

// Test_alienSpawner_Spawn accounts for "predictable randomness", we expect an evenly distributed spawning, whether if we have 4 cities and n aliens we should get the same number of aliens in every city if n % 4 == 0
func Test_alienSpawner_Spawn(t *testing.T) {
	london := &City{
		Name:       "London",
		Directions: map[Direction]*City{},
	}

	rome := &City{
		Name:       "Rome",
		Directions: map[Direction]*City{},
	}

	oslo := &City{
		Name:       "Oslo",
		Directions: map[Direction]*City{},
	}

	washington := &City{
		Name:       "Washington",
		Directions: map[Direction]*City{},
	}

	type fields struct {
		NumAliens      int
		World          *World
		AlienLocations map[*Alien]string
		cities         []string
	}
	tests := []struct {
		name         string
		fields       fields
		aliensInCity map[string]int
		wantErr      bool
	}{
		{
			name: "4 cities 15 aliens",
			fields: fields{
				NumAliens: 15,
				World: &World{
					config: &Config{},
					Map: &Map{
						Cities: map[string]*City{
							"London":     london,
							"Rome":       rome,
							"Oslo":       oslo,
							"Washington": washington,
						},
						sortedCityNames: []string{"London", "Oslo", "Rome", "Washington"},
					},
					Aliens: map[uuid.UUID]*Alien{},
				},
				AlienLocations: map[*Alien]string{},
				cities:         []string{"London", "Oslo", "Rome", "Washington"},
			},
			aliensInCity: map[string]int{
				"London":     3,
				"Oslo":       4,
				"Rome":       4,
				"Washington": 4,
			},
			wantErr: false,
		},

		{
			name: "4 cities 16 aliens",
			fields: fields{
				NumAliens: 16,
				World: &World{
					config: &Config{},
					Map: &Map{
						Cities: map[string]*City{
							"London":     london,
							"Rome":       rome,
							"Oslo":       oslo,
							"Washington": washington,
						},
						sortedCityNames: []string{"London", "Oslo", "Rome", "Washington"},
					},
					Aliens: map[uuid.UUID]*Alien{},
				},
				AlienLocations: map[*Alien]string{},
				cities:         []string{"London", "Oslo", "Rome", "Washington"},
			},
			aliensInCity: map[string]int{
				"London":     4,
				"Oslo":       4,
				"Rome":       4,
				"Washington": 4,
			},
			wantErr: false,
		},
		{
			name: "4 cities 100 aliens",
			fields: fields{
				NumAliens: 100,
				World: &World{
					config: &Config{},
					Map: &Map{
						Cities: map[string]*City{
							"London":     london,
							"Rome":       rome,
							"Oslo":       oslo,
							"Washington": washington,
						},
						sortedCityNames: []string{"London", "Oslo", "Rome", "Washington"},
					},
					Aliens: map[uuid.UUID]*Alien{},
				},
				AlienLocations: map[*Alien]string{},
				cities:         []string{"London", "Oslo", "Rome", "Washington"},
			},
			aliensInCity: map[string]int{
				"London":     25,
				"Oslo":       25,
				"Rome":       25,
				"Washington": 25,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Random = NewFakeRandomizer()
			defer func() { Random = NewRealRandomizer() }()
			a := &alienSpawner{
				NumAliens:      tt.fields.NumAliens,
				World:          tt.fields.World,
				AlienLocations: tt.fields.AlienLocations,
				cities:         tt.fields.cities,
			}
			if err := a.Spawn(); (err != nil) != tt.wantErr {
				t.Errorf("alienSpawner.Spawn() error = %v, wantErr %v", err, tt.wantErr)
			}

			gotAliensInCity := make(map[string]int)
			for _, alien := range a.World.Aliens {
				gotAliensInCity[alien.Location.Name]++
			}

			if !reflect.DeepEqual(gotAliensInCity, tt.aliensInCity) {
				t.Errorf("alienSpawner.Spawn() expected aliensInCity = %v, got %v", tt.aliensInCity, gotAliensInCity)
			}

		})
	}
}
