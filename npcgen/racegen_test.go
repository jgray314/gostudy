package main

import (
	"io"
	"strings"
	"testing"
)

var coveragestring = "50 human, 10 dwarf, 10 elf, 5 halfling, 10 halfelf, 5 gnome, 5 halforc, 3 dragonborn, 2 tiefling"

func TestRaceGen_LoadFromString(t *testing.T) {
	rg := RaceGen{}
	gt := GeneratorTable{size: 100, typevalidator: StringTypeValidator{}}
	gt.LoadFromString("100 human")
	dullt := gt.table
	gt.LoadFromString(coveragestring)
	allt := gt.table
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		rg      *RaceGen
		args    args
		exp     []string
		wantErr bool
	}{
		{"Dull Success", &rg, args{"100 human"}, dullt, false},
		{"Dull Failure", &rg, args{"99 human"}, nil, true},
		{"Rollable Wrong Number", &rg, args{"20 human"}, nil, true},
		{"Accept All Race", &rg, args{coveragestring}, allt, false},
		{"All Race Wrong Number", &rg, args{"50 human, 10 dwarf, 10 elf, 5 halfling, 10 halfelf, 5 gnome, 5 halforc, 3 dragonborn, 1 tiefling"}, nil, true},
		{"One Invalid", &rg, args{"99 human, 1 diety"}, nil, true},
		{"All Invalid", &rg, args{"25 warg, 25 ent, 50 eagle"}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.rg.LoadFromString(tt.args.s); (err != nil) != tt.wantErr {
				t.Errorf("RaceGen.LoadFromString() error = %v, wantErr %v", err, tt.wantErr)
			}
			// Nil here indicates undefined behavior.
			if tt.exp != nil && strings.Join(tt.rg.gt.table, ",") != strings.Join(tt.exp, ",") {
				t.Errorf("Table does not match expected form.\n Got:      %v\nExpected: %v\n", tt.rg.gt.table, tt.exp)
			}
		})
	}
}

func csvFromTable(table []string) io.Reader {
	return strings.NewReader(strings.Join(table, "\n"))
}

func newTable99(supported []string) []string {
	t := []string{}
	for i := 0; i < 99; i++ {
		val := supported[i%len(supported)]
		t = append(t, string(val))
	}
	return t
}

func TestRaceGen_LoadFromCsvIoReader(t *testing.T) {
	table99 := newTable99(RaceValidator{}.GetSupported())
	table100 := append(newTable99(RaceValidator{}.GetSupported()), "human")
	tablewrong := append(newTable99(RaceValidator{}.GetSupported()), "witcher")
	rg := RaceGen{}

	type args struct {
		csvfile io.Reader
	}
	tests := []struct {
		name          string
		rg            *RaceGen
		args          args
		expectedtable []string
		wantErr       bool
	}{
		{"Table Size Failure", &rg, args{csvFromTable(table99)}, nil, true},
		{"Table Works", &rg, args{csvFromTable(table100)}, table100, false},
		{"Validation Failure", &rg, args{csvFromTable(tablewrong)}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.rg.LoadFromCsvIoReader(tt.args.csvfile); (err != nil) != tt.wantErr {
				t.Errorf("RaceGen.LoadFromCsvIoReader() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.expectedtable != nil && strings.Join(tt.expectedtable, " ") != strings.Join(tt.rg.gt.table, " ") {
				t.Errorf("RaceGen.LoadFromCsvIoReader() did not load expected table. Got: %v Expected: %v",
					tt.rg.gt.table, tt.expectedtable)
			}
		})
	}
}
func TestRaceGen_Roll(t *testing.T) {
	d := Dice{}
	d.Init(121)
	rg := RaceGen{}
	if e := rg.LoadFromString(coveragestring); e != nil {
		t.Errorf("Error loading from string in setup. %v/n", e)
	}
	rg.Init(d)

	m := map[Race]int{}
	for i := 0; i < 10000; i++ {
		r, e := rg.Roll()
		if e != nil {
			t.Errorf("Error on roll %d: %v", i, e)
		}
		v, _ := m[r]
		m[r] = v + 1
	}
	exp := map[Race]int{
		Race("human"):      5020,
		Race("dwarf"):      970,
		Race("elf"):        974,
		Race("halfling"):   517,
		Race("halfelf"):    1008,
		Race("gnome"):      483,
		Race("halforc"):    534,
		Race("dragonborn"): 290,
		Race("tiefling"):   204,
	}
	for k, v := range exp {
		found, _ := m[k]
		if found != v {
			t.Errorf("Count for race \"%v\" was not as expected. Got %d, expected %d.\n", k, found, v)
		}
	}
}
