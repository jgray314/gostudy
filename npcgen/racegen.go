package main

// TODO: Generalize different tables want to support / map to allowed dice sizes. Create tests.

import (
	"fmt"
	"strings"
)

// lowercase as cannonical
type Race string

// Fun with enums
const (
	Dwarf      Race = "dwarf"
	Elf             = "elf"
	Halfling        = "halfling"
	Human           = "human"
	Dragonborn      = "dragonborn"
	Gnome           = "gnome"
	HalfElf         = "halfelf"
	HalfOrc         = "halforc"
	Tiefling        = "tiefling"
)

var SupportedRaces = []Race{Dwarf, Elf, Halfling, Human, Dragonborn, Gnome, HalfElf, HalfOrc, Tiefling}

func isRace(s string) bool {
	for _, v := range SupportedRaces {
		if s == string(v) {
			return true
		}
	}
	return false
}

type RaceGen struct {
	centIdx []Race
	detail  bool
}

func (rg *RaceGen) Detail(detail bool) {
	rg.detail = detail
}

func (rg RaceGen) validateSize() error {
	if len(rg.centIdx) != 100 {
		return fmt.Errorf("Table incorrect size. Expected 100, Got %d", len(rg.centIdx))
	}
	return nil
}

// Takes a comma separated string for centile counts and races initializes RaceGen to match. Must sum to 100.
// eg. " 50 human, 25 elf, 5 halfelf, 20 dwarf"
func (rg *RaceGen) LoadFromString(s string) error {
	if rg.detail {
		fmt.Println("Loading from string: ", s)
	}
	// Start by clearing
	rg.centIdx = nil
	for _, p := range strings.Split(s, ",") {
		var n int
		var rs string
		_, e := fmt.Sscanf(p, "%d %s", &n, &rs)
		if e != nil {
			return e
		}
		rs = strings.ToLower(rs)
		if !isRace(rs) {
			return fmt.Errorf("Unsupported race: %v. Expected one of %v", s, SupportedRaces)
		}
		r := Race(strings.ToLower(rs))
		for i := 0; i < n; i++ {
			rg.centIdx = append(rg.centIdx, r)
		}
	}
	if rg.detail {
		fmt.Println("Loaded table to be validated:\n", rg.centIdx)
	}
	return rg.validateSize()
}

// TODO add load from CSV

func main() {
	rg := RaceGen{}
	rg.Detail(true)
	rg.LoadFromString("50 human, 5 dwarf, 10 elf, 5 halfling, 5 halfelf, 5 gnome, 5 halforc, 3 dragonborn, 2 tiefling")
}
