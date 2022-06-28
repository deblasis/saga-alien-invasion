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

	North string
	East  string
	South string
	West  string
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
					Name:  "Foo",
					North: "Bar",
					East:  "",
					South: "Qu-ux",
					West:  "Baz",
				},
				"Bar": {
					Name:  "Bar",
					North: "",
					East:  "",
					South: "Foo",
					West:  "Bee",
				},
				"Baz": {
					Name:  "Baz",
					North: "",
					East:  "Foo",
					South: "",
					West:  "",
				},
				"Qu-ux": {
					Name:  "Qu-ux",
					North: "Foo",
					East:  "",
					South: "",
					West:  "",
				},
				"Bee": {
					Name:  "Bee",
					North: "",
					East:  "Bar",
					South: "",
					West:  "",
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
		if city.North != nil {
			lofiCity.North = city.North.Name
		}
		if city.East != nil {
			lofiCity.East = city.East.Name
		}
		if city.South != nil {
			lofiCity.South = city.South.Name
		}
		if city.West != nil {
			lofiCity.West = city.West.Name
		}
		lofiMap[cityName] = *lofiCity
	}

	return lofiMap
}
