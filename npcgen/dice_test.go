package main

import (
	"testing"
)

func Test_scanRoll(t *testing.T) {
	var num, sides, plus int
	type args struct {
		s     string
		num   *int
		sides *int
		plus  *int
	}
	tests := []struct {
		name                               string
		args                               args
		want, wantNum, wantSides, wantPlus int
		wantErr                            bool
	}{
		{"Min", args{"d12", &num, &sides, &plus}, 1, 1, 12, 0, false},
		{"Mult", args{"2d12", &num, &sides, &plus}, 2, 2, 12, 0, false},
		{"Plus", args{"d12+2", &num, &sides, &plus}, 2, 1, 12, 2, false},
		{"MultPlus", args{"2d12+2", &num, &sides, &plus}, 3, 2, 12, 2, false},
		{"Minus", args{"d12-2", &num, &sides, &plus}, 2, 1, 12, -2, false},
		{"MultMinus", args{"2d12-2", &num, &sides, &plus}, 3, 2, 12, -2, false},
		{"OddValid", args{"21d5-27", &num, &sides, &plus}, 3, 21, 5, -27, false},
		{"Just d", args{"d", &num, &sides, &plus}, 0, 1, 0, 0, true},
		{"No sides", args{"2d+2", &num, &sides, &plus}, 2, 2, 2, 0, true},
		{"Extra letter mid", args{"2d3a+2", &num, &sides, &plus}, 2, 2, 3, 0, true},
		{"Extra letter end", args{"2d3+2a", &num, &sides, &plus}, 3, 2, 3, 2, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := scanRoll(tt.args.s, tt.args.num, tt.args.sides, tt.args.plus)
			if (err != nil) != tt.wantErr {
				t.Errorf("scanRoll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("scanRoll() = %v, want %v", got, tt.want)
			}
			if num != tt.wantNum || sides != tt.wantSides || plus != tt.wantPlus {
				t.Errorf("scanRoll() did not modify as expected. Got = %d %d %d, Wanted = %d %d %d",
					num, sides, plus, tt.wantNum, tt.wantSides, tt.wantPlus)
			}
		})
	}
}

func TestDice_Roll(t *testing.T) {
	d := Dice{}
	d.Init(100)
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		d       Dice
		args    args
		want    int
		wantErr bool
	}{
		{"Basic", d, args{"1d6"}, 2, false},
		{"Multiple", d, args{"6d6"}, 15, false},
		{"Plus", d, args{"1d6+6"}, 10, false},
		{"Coin", d, args{"1d2"}, 1, false},
		{"d4", d, args{"1d4"}, 2, false},
		{"d8", d, args{"1d8"}, 2, false},
		{"d10", d, args{"1d10"}, 3, false},
		{"d12", d, args{"1d12"}, 2, false},
		{"Main", d, args{"1d20"}, 17, false},
		{"Centile", d, args{"1d100"}, 24, false},
		{"No Count", d, args{"d20"}, 13, false},
		{"Invalid Sides", d, args{"1d7"}, 0, true},
		{"NAN", d, args{"1d6+a"}, 0, true},
		{"Minus", d, args{"1d20-2"}, -1, false},
		{"Minus again", d, args{"1d20-2"}, 5, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.d.Roll(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("Dice.Roll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Dice.Roll() = %v, want %v", got, tt.want)
			}
		})
	}
}
