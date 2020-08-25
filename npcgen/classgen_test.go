package main

import (
	"io"
	"strings"
	"testing"
)

var dullstring = "100 fighter"
var demostring = "20 druid, 10 monk, 20 paladin, 30 rogue, 20 warlock"

func TestClassGen_LoadFromString(t *testing.T) {
	cg := ClassGen{}
	gt := GeneratorTable{size: 100, typevalidator: StringTypeValidator{}}
	gt.LoadFromString(dullstring)
	dullt := gt.table
	gt.LoadFromString(demostring)
	demot := gt.table
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		cg      *ClassGen
		args    args
		exp     []string
		wantErr bool
	}{
		{"Dull", &cg, args{dullstring}, dullt, false},
		{"Mix", &cg, args{demostring}, demot, false},
		{"Miscount", &cg, args{"19" + demostring[2:]}, nil, true},
		{"Validation", &cg, args{demostring[:len(demostring)-4]}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.cg.LoadFromString(tt.args.s); (err != nil) != tt.wantErr {
				t.Errorf("ClassGen.LoadFromString() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.exp != nil && strings.Join(tt.exp, " ") != strings.Join(tt.cg.defaultgt.table, " ") {
				t.Errorf("Unexpected loaded string. Got: %v Expected: %v", tt.cg.defaultgt.table, tt.exp)
			}
		})
	}
}

func TestClassGen_LoadFromCsvIoReader(t *testing.T) {
	table99 := newTable99(ClassValidator{}.GetSupported())
	table100 := append(newTable99(ClassValidator{}.GetSupported()), "cleric")
	tablewrong := append(newTable99(ClassValidator{}.GetSupported()), "priest")
	cg := ClassGen{}
	type args struct {
		csvfile io.Reader
	}
	tests := []struct {
		name    string
		cg      *ClassGen
		args    args
		exp     []string
		wantErr bool
	}{
		{"Wrong length", &cg, args{csvFromTable(table99)}, nil, true},
		{"Valid", &cg, args{csvFromTable(table100)}, table100, false},
		{"Validation failure", &cg, args{csvFromTable(tablewrong)}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.cg.LoadFromCsvIoReader(tt.args.csvfile); (err != nil) != tt.wantErr {
				t.Errorf("ClassGen.LoadFromCsvIoReader() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.exp != nil && strings.Join(tt.exp, " ") != strings.Join(tt.cg.defaultgt.table, " ") {
				t.Errorf("Unexpected loaded string. Got: %v Expected: %v", tt.cg.defaultgt.table, tt.exp)
			}
		})
	}
}

func TestClassGen_Roll(t *testing.T) {
	d := Dice{}
	d.Init(121)
	cg := ClassGen{dice: d}
	if e := cg.LoadFromString(demostring); e != nil {
		t.Errorf("Error loading from string in setup. %v/n", e)
	}

	m := map[Class]int{}
	for i := 0; i < 10000; i++ {
		c, e := cg.Roll()
		if e != nil {
			t.Errorf("Error on roll %d: %v", i, e)
		}
		v, _ := m[c]
		m[c] = v + 1
	}
	exp := map[Class]int{
		Class("druid"):   1995,
		Class("monk"):    1004,
		Class("paladin"): 2021,
		Class("rogue"):   2980,
		Class("warlock"): 2000,
	}
	for k, v := range exp {
		found, _ := m[k]
		if found != v {
			t.Errorf("Count for class \"%v\" was not as expected. Got %d, expected %d.\n", k, found, v)
		}
	}
}
