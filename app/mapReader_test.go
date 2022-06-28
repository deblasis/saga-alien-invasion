package app

import (
	"fmt"
	"io"
	"strings"
	"testing"
)

// lofiCity is a simple DTO used to create assertions because we won't know the *City pointers ahead of time. With this, we'll rely on the City.Name
type lofiCity struct {
	Name string

	Directions map[Direction]string
}

func NewLofiCity(name string) *lofiCity {
	return &lofiCity{
		Name: name,
	}
}

func Test_mapReader_ParseMapFile(t *testing.T) {
	type fields struct {
		Cities map[string]*City
	}
	type args struct {
		reader io.Reader
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    map[string]lofiCity
		wantErr bool
	}{
		{
			name: "example map should be parsed correctly",
			fields: fields{
				Cities: map[string]*City{},
			},
			args: args{
				reader: strings.NewReader(`Foo north=Bar west=Baz south=Qu-ux
Bar south=Foo west=Bee`),
			},
			want: map[string]lofiCity{
				"Foo": {
					Name: "Foo",
					Directions: map[Direction]string{
						NORTH: "Bar",
						EAST:  "",
						SOUTH: "Qu-ux",
						WEST:  "Baz",
					},
				},
				"Bar": {
					Name: "Bar",
					Directions: map[Direction]string{
						NORTH: "",
						EAST:  "",
						SOUTH: "Foo",
						WEST:  "Bee",
					},
				},
				"Baz": {
					Name: "Baz",
					Directions: map[Direction]string{
						NORTH: "",
						EAST:  "Foo",
						SOUTH: "",
						WEST:  "",
					},
				},
				"Qu-ux": {
					Name: "Qu-ux",
					Directions: map[Direction]string{
						NORTH: "Foo",
						EAST:  "",
						SOUTH: "",
						WEST:  "",
					},
				},
				"Bee": {
					Name: "Bee",
					Directions: map[Direction]string{
						NORTH: "",
						EAST:  "Bar",
						SOUTH: "",
						WEST:  "",
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mr := &mapReader{
				Cities: tt.fields.Cities,
			}
			got, err := mr.ParseMapFile(tt.args.reader)
			if (err != nil) != tt.wantErr {
				t.Errorf("mapreader.ParseMapFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			gotLofi := map2LofiMap(got)
			if !(fmt.Sprintf("%v", gotLofi) == fmt.Sprintf("%v", tt.want)) {
				t.Errorf("mapreader.ParseMapFile() = %v, want %v", gotLofi, tt.want)
			}
		})
	}
}

func map2LofiMap(m *Map) map[string]lofiCity {
	lofiMap := make(map[string]lofiCity)

	if len(m.Cities) == 0 {
		return lofiMap
	}

	for cityName, city := range m.Cities {
		lofiCity := NewLofiCity(cityName)
		if city.Directions[NORTH] != nil {
			lofiCity.Directions[NORTH] = city.Directions[NORTH].Name
		}
		if city.Directions[EAST] != nil {
			lofiCity.Directions[EAST] = city.Directions[EAST].Name
		}
		if city.Directions[SOUTH] != nil {
			lofiCity.Directions[SOUTH] = city.Directions[SOUTH].Name
		}
		if city.Directions[WEST] != nil {
			lofiCity.Directions[WEST] = city.Directions[WEST].Name
		}
		lofiMap[cityName] = *lofiCity
	}

	return lofiMap
}
