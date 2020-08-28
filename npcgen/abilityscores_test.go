package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRollAbilityValues(t *testing.T) {
	d := Dice{}
	d.Init(21)
	rollHist := make(map[int]int)
	for i := 0; i < 1000; i++ {
		l, e := RollAbilityValues(d)

		if e != nil {
			t.Error("Unexpected error for RollAbilityValues(): ", e)
		}
		if len(l) != 6 {
			t.Error("Unexpected length for RollAbilityValues(). Expected length 6. Got: ", l)
		}
		for i, v := range l {
			if i != 0 && v > l[i-1] {
				t.Error("RollAbilityValues() not in decending order: ", l)
			}
			if v > 18 || v < 3 {
				t.Error("RollAbilityValues() returned value out of supported range: ", v)
			}
			rollHist[v] += 1
		}
	}
	fmt.Printf("Histogram of raw abilities: %v\n", rollHist)
}

func Test_randomBottom(t *testing.T) {
	d := Dice{}
	d.Init(21)
	std := StandardAbilityValues()
	top, second := std[0], std[1]
	stdbottom := std[2:]
	tests := []struct {
		name   string
		remain []int
		a      Abilities
		exp    Abilities
	}{
		{"STR, CON", stdbottom, Abilities{strength: top, constitution: second}, Abilities{15, 12, 14, 8, 10, 13}},
		{"INT, WIS", stdbottom, Abilities{intelligence: top, wisdom: second}, Abilities{13, 8, 10, 15, 14, 12}},
		{"DEX, CHA", stdbottom, Abilities{dexterity: top, charisma: second}, Abilities{13, 15, 10, 12, 8, 14}},
		// Different from first
		{"STR, CON Again", stdbottom, Abilities{strength: top, constitution: second}, Abilities{15, 10, 14, 13, 12, 8}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			randomBottom(tt.remain, d, &tt.a)
			if tt.a != tt.exp {
				t.Errorf("Unexpected randomizedBottom result. Got: %v Expected:%v", tt.a, tt.exp)
			}
		})
	}
}

func TestAssignAbilities(t *testing.T) {
	type args struct {
		c  Class
		av []int
		d  Dice
	}
	tests := []struct {
		name    string
		args    args
		want    Abilities
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AssignAbilities(tt.args.c, tt.args.av, tt.args.d)
			if (err != nil) != tt.wantErr {
				t.Errorf("AssignAbilities() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AssignAbilities() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdjustAbilities(t *testing.T) {
	tests := []struct {
		r   Race
		exp Abilities
	}{
		{Human, Abilities{16, 15, 14, 13, 11, 9}},
		{Dragonborn, Abilities{17, 14, 13, 12, 10, 9}},
		{HalfOrc, Abilities{17, 14, 14, 12, 10, 8}},
		{Tiefling, Abilities{15, 14, 13, 13, 10, 10}},
		// TODO: add in more checkes once account for subraces
	}
	for _, tt := range tests {
		t.Run(string(tt.r), func(t *testing.T) {
			av := StandardAbilityValues()
			a := Abilities{av[0], av[1], av[2], av[3], av[4], av[5]}
			AdjustAbilities(tt.r, &a)
		})
	}
}

func TestAbilityValueToModifier(t *testing.T) {

	tests := []struct {
		n    int
		want int
	}{
		{1, -5},
		{3, -4},
		{8, -1},
		{9, -1},
		{10, 0},
		{11, 0},
		{12, 1},
		{13, 1},
		{14, 2},
		{17, 3},
		{19, 4},
		{20, 5},
		{23, 6},
		{30, 10},
	}
	for _, tt := range tests {
		t.Run(string(tt.n), func(t *testing.T) {
			if got := AbilityValueToModifier(tt.n); got != tt.want {
				t.Errorf("AbilityValueToModifier(%d) = %v, want %v", tt.n, got, tt.want)
			}
		})
	}
}
