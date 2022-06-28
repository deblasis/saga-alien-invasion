package app

import (
	"bytes"
	"testing"
)

func TestMapWriter_WriteMap(t *testing.T) {

	manchester := &City{
		Name:       "Manchester",
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

	london := &City{
		Name: "London",
		Directions: map[Direction]*City{
			NORTH: manchester,
			SOUTH: rome,
			WEST:  washington,
			EAST:  oslo,
		},
	}

	type fields struct {
		Map *Map
	}
	tests := []struct {
		name       string
		fields     fields
		wantWriter string
		wantErr    bool
	}{
		{
			name: "single city",
			fields: fields{
				Map: &Map{
					Cities: map[string]*City{
						"London": london,
					},
					sortedCityNames: []string{"London", "Manchester", "Oslo", "Rome", "Washington"},
				},
			},
			wantWriter: "London north=Manchester east=Oslo south=Rome west=Washington\n",
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MapWriter{
				Map: tt.fields.Map,
			}
			writer := &bytes.Buffer{}
			if err := m.WriteMap(writer); (err != nil) != tt.wantErr {
				t.Errorf("MapWriter.WriteMap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotWriter := writer.String(); gotWriter != tt.wantWriter {
				t.Errorf("MapWriter.WriteMap() = %v, want %v", gotWriter, tt.wantWriter)
			}
		})
	}
}
