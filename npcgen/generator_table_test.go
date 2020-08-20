package main

import (
	"strings"
	"testing"
)

func TestGeneratorTable_validateSize(t *testing.T) {
	tests := []struct {
		name    string
		gt      GeneratorTable
		wantErr bool
	}{
		{"4 correct", GeneratorTable{size: 4, table: []string{"a", "b", "c", "d"}}, false},
		{"6 correct", GeneratorTable{size: 6, table: []string{"a", "b", "c", "d", "e", "f"}}, false},
		{"6 dups", GeneratorTable{size: 6, table: []string{"ab", "bc", "c", "d", "d", "d"}}, false},
		{"4 empty", GeneratorTable{size: 4, table: []string{}}, true},
		{"4 too few", GeneratorTable{size: 4, table: []string{"a", "b", "c"}}, true},
		{"4 too many", GeneratorTable{size: 4, table: []string{"a", "b", "c", "d", "e"}}, true},
		{"Unsupported num sides", GeneratorTable{size: 3, table: []string{"a", "b", "c"}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.gt.validateSize(); (err != nil) != tt.wantErr {
				t.Errorf("GeneratorTable.validateSize() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

type RejectTypeValidator struct{}

func (rtv RejectTypeValidator) IsValid(s string) bool  { return false }
func (rtv RejectTypeValidator) GetSupported() []string { return []string{} }

func TestGeneratorTable_LoadFromString(t *testing.T) {
	dice := Dice{}
	gt4 := GeneratorTable{name: "Caltrop", size: 4, dice: dice, typevalidator: StringTypeValidator{}}
	gtr := GeneratorTable{name: "Reject", size: 4, dice: dice, typevalidator: RejectTypeValidator{}}
	type args struct {
		s string
	}
	tests := []struct {
		name          string
		gt            *GeneratorTable
		args          args
		expectedtable []string
		wantErr       bool
	}{
		{"Basic", &gt4, args{"1 normal, 3 basic"}, []string{"normal", "basic", "basic", "basic"}, false},
		{"All Same Cap Mix", &gt4, args{"4 BaSiC"}, []string{"basic", "basic", "basic", "basic"}, false},
		{"All Same Another Way", &gt4, args{"2 basic, 2 basic"}, []string{"basic", "basic", "basic", "basic"}, false},
		{"Interleave", &gt4, args{"1 a, 1 b, 1 a, 1 c"}, []string{"a", "b", "a", "c"}, false},
		{"Space Strings", &gt4, args{"2 aba cad, 2 foo bar"}, []string{"aba cad", "aba cad", "foo bar", "foo bar"}, false},
		{"Fail Split", &gt4, args{"2close, 2num"}, nil, true},
		{"Fail Count", &gt4, args{"1 value, 2 few"}, nil, true},
		{"Fail Verifier", &gtr, args{"4 BaSiC"}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.gt.LoadFromString(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("GeneratorTable.LoadFromString() error = %v, wantErr %v", err, tt.wantErr)
			}
			// Only define state of table if no error is thrown
			if err == nil && strings.Join(tt.gt.table, " ") != strings.Join(tt.expectedtable, " ") {
				t.Errorf("GeneratorTable.LoadFromString() unexpected result.\nGot:     %v\nExpected:%v", tt.gt.table, tt.expectedtable)
			}
		})
	}
}

func TestGeneratorTable_Roll(t *testing.T) {
	dice := Dice{}
	dice.Init(121) // For reproducability. Also know this particular seed has coverage
	gt10 := GeneratorTable{name: "Decile", size: 10, dice: dice, typevalidator: StringTypeValidator{}}
	gt10.LoadFromString("5 human, 3 elf, 2 dwarf")
	tests := []struct {
		name    string
		gt      GeneratorTable
		want    string
		wantErr bool
	}{
		{"Uninitialized", GeneratorTable{name: "Decile", size: 10, dice: dice, typevalidator: StringTypeValidator{}}, "", true},
		{"Roll 1", gt10, "elf", false},
		{"Roll 2", gt10, "human", false},
		{"Roll 3", gt10, "elf", false},
		{"Roll 4", gt10, "human", false},
		{"Roll 5", gt10, "human", false},
		{"Roll 6", gt10, "dwarf", false},
		{"Roll 7", gt10, "human", false},
		{"Roll 8", gt10, "dwarf", false},
		{"Roll 9", gt10, "dwarf", false},
		{"Roll 10", gt10, "human", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.gt.Roll()
			if (err != nil) != tt.wantErr {
				t.Errorf("GeneratorTable.Roll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GeneratorTable.Roll() = %v, want %v", got, tt.want)
			}
		})
	}
}
