package main

import (
	"fmt"
	"io"
)

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

type RaceValidator struct{}

func (rv RaceValidator) IsValid(s string) bool {
	for _, v := range SupportedRaces {
		if s == string(v) {
			return true
		}
	}
	return false
}

func (rv RaceValidator) GetSupported() []string {
	r := []string{}
	for _, v := range SupportedRaces {
		r = append(r, string(v))
	}
	return r
}

type RaceGen struct {
	gt GeneratorTable
}

func (rg *RaceGen) Init(d Dice) {
	rg.gt.dice = d
}

func (rg *RaceGen) Detail(detail bool) {
	rg.gt.detail = detail
}

func (rg *RaceGen) enforceConstraints() {
	rg.gt.typevalidator = RaceValidator{}
	rg.gt.size = 100
}

func (rg *RaceGen) LoadFromString(s string) error {
	rg.enforceConstraints()
	return rg.gt.LoadFromString(s)
}

func (rg *RaceGen) LoadFromCsvFile(filename string) error {
	rg.enforceConstraints()
	return rg.gt.LoadFromCsvFile(filename)
}

func (rg *RaceGen) LoadFromCsvIoReader(csvfile io.Reader) error {
	rg.enforceConstraints()
	return rg.gt.LoadFromCsvIoReader(csvfile)
}

func (rg RaceGen) Roll() (Race, error) {
	rg.enforceConstraints()
	rs, e := rg.gt.Roll()
	return Race(rs), e
}

func main() {
	rg := RaceGen{}
	d := Dice{}
	d.Init(121)
	rg.Init(d)
	rg.Detail(true)
	rg.LoadFromString("50 human, 10 dwarf, 10 elf, 5 halfling, 10 halfelf, 5 gnome, 5 halforc, 3 dragonborn, 2 tiefling")
	s, e := rg.Roll()
	fmt.Println(s, e)
	s, e = rg.Roll()
	fmt.Println(s, e)
}
