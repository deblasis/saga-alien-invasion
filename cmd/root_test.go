package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/deblasis/saga-alien-invasion/app"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func init() {
	log.Logger = zerolog.Logger{}
}

func Test_InvalidNoInput(t *testing.T) {
	cmd := rootCmd
	b := bytes.NewBufferString("")
	bErr := bytes.NewBufferString("")
	cmd.SetOut(b)
	cmd.SetErr(bErr)
	cmd.Execute()
	out, err := ioutil.ReadAll(bErr)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(out), "Error: the number of aliens invading the world must be specified") {
		t.Fatalf("expected error got \"%s\"", string(out))
	}
}

func Test_InvalidInputNotInt(t *testing.T) {
	cmd := rootCmd
	b := bytes.NewBufferString("")
	bErr := bytes.NewBufferString("")
	cmd.SetOut(b)
	cmd.SetErr(bErr)
	cmd.SetArgs([]string{fmt.Sprint(0)})
	cmd.Execute()
	out, err := ioutil.ReadAll(bErr)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(out), "Error: the number of aliens must be *checks notes* a positive number") {
		t.Fatalf("expected error got \"%s\"", string(out))
	}
}

func Test_ValidInput(t *testing.T) {
	app.Random = app.NewFakeRandomizer()
	defer func() { app.Random = app.NewRealRandomizer() }()

	cmd := rootCmd
	b := new(bytes.Buffer)
	cmd.SetOut(b)
	cmd.SetErr(b)
	cmd.SetArgs([]string{fmt.Sprint(15), "--" + app.MapFileFlag, "../map.txt", "--verbose"})
	cmd.Execute()
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(out), "Error") {
		t.Fatalf("expected no errors, got: %v", string(out))
	}
}

func Test_FileBased(t *testing.T) {

	type args struct {
		inputFile string
		numAliens int
	}

	tests := []struct {
		name       string
		args       args
		wantGolden string
		wantStdErr string
	}{
		{
			name: "example input with 1 alien",
			args: args{
				inputFile: "example.input",
				numAliens: 1,
			},
			wantGolden: "example_1.golden",
			wantStdErr: "",
		},
		{
			name: "example input with 100 aliens (total annihilation on day 1)",
			args: args{
				inputFile: "example.input",
				numAliens: 100,
			},
			wantGolden: "example_100.golden",
			wantStdErr: "",
		},
		{
			name: "unconnected cities input with 2 aliens",
			args: args{
				inputFile: "unconnected_cities.input",
				numAliens: 2,
			},
			wantGolden: "unconnected_cities.golden",
			wantStdErr: "",
		},
		// single city tests cover also for the fact that connections become cities as well
		{
			name: "single city connected to many with 2 aliens",
			args: args{
				inputFile: "single_city_with_connections.input",
				numAliens: 2,
			},
			wantGolden: "single_city_with_connections_2.golden",
			wantStdErr: "",
		},
		// this tests for an alien force that can cause damage but the game will end with cities left intact
		{
			name: "single city connected to many with 6 aliens",
			args: args{
				inputFile: "single_city_with_connections.input",
				numAliens: 6,
			},
			wantGolden: "single_city_with_connections_6.golden",
			wantStdErr: "",
		},
		// this edge case tests the fact that the battle on day 1 is already enough to weaken the alien force so that it cannot cause any more harm
		{
			name: "single city connected to many with 7 aliens",
			args: args{
				inputFile: "single_city_with_connections.input",
				numAliens: 7,
			},
			wantGolden: "single_city_with_connections_7.golden",
			wantStdErr: "",
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			app.Random = app.NewFakeRandomizer()
			defer func() { app.Random = app.NewRealRandomizer() }()

			path := "../cmd/testdata/" + tt.args.inputFile
			out, outErr, _ := executeCommand(rootCmd, fmt.Sprint(tt.args.numAliens), "--"+app.MapFileFlag, path)

			goldenfile := filepath.Join("testdata", tt.wantGolden)
			want, err := os.ReadFile(goldenfile)
			if err != nil {
				t.Fatal("error reading golden file:", err)
			}

			if !bytes.Equal([]byte(out), want) {
				t.Errorf("\n==== got:\n%s\n==== want:\n%s\n", out, want)
			}

			if !bytes.Equal([]byte(outErr), []byte(tt.wantStdErr)) {
				t.Errorf("\n==== got outErr:\n%s\n==== want:\n%s\n", outErr, tt.wantStdErr)
			}
		})
	}

}

func executeCommand(root *cobra.Command, args ...string) (output string, outErr string, err error) {
	buf := new(bytes.Buffer)
	root.SetOut(buf)
	bufErr := new(bytes.Buffer)
	root.SetErr(bufErr)
	root.SetArgs(args)

	err = root.Execute()
	if err != nil {
		fmt.Println(err)
	}

	return buf.String(), bufErr.String(), err
}
