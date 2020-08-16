package main

import (
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

func TestGeneratorTable_LoadFromString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		gt      *GeneratorTable
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.gt.LoadFromString(tt.args.s); (err != nil) != tt.wantErr {
				t.Errorf("GeneratorTable.LoadFromString() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
