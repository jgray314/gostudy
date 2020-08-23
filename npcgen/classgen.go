package main

import "io"

type Class string

// Limiting to PHB core classes.
const (
	Barbarian Class = "barbarian"
	Bard            = "bard"
	Cleric          = "cleric"
	Druid           = "druid"
	Fighter         = "fighter"
	Monk            = "monk"
	Paladin         = "paladin"
	Ranger          = "ranger"
	Rogue           = "rogue"
	Sorcerer        = "sorcerer"
	Warlock         = "warlock"
	Wizard          = "wizard"
)

var SupportedClasses = []Class{Barbarian, Bard, Cleric, Druid, Fighter, Monk, Paladin, Ranger, Rogue, Sorcerer, Warlock, Wizard}

// TODO: Add in class specializations such as schools
type ClassValidator struct{}

func (cv ClassValidator) IsValid(s string) bool {
	for _, v := range SupportedClasses {
		if s == string(v) {
			return true
		}
	}
	return false
}

func (cv ClassValidator) GetSupported() []string {
	r := []string{}
	for _, v := range SupportedClasses {
		r = append(r, string(v))
	}
	return r
}

type ClassGen struct {
	dice      Dice
	detail    bool
	defaultgt GeneratorTable
	// racetogt  map[Race]GeneratorTable  // For future exploration
}

func (cg ClassGen) enforceConstraints(gt *GeneratorTable) {
	gt.dice = cg.dice
	gt.detail = cg.detail
	gt.typevalidator = ClassValidator{}
	gt.size = 100
}

// Following fields allow operations around class generation that do not consider character race.
// In a less complicated interaction thes are all that are needed.
func (cg *ClassGen) LoadFromString(s string) error {
	cg.enforceConstraints(&cg.defaultgt)
	return cg.defaultgt.LoadFromString(s)
}

func (cg *ClassGen) LoadFromCsvFile(filename string) error {
	cg.enforceConstraints(&cg.defaultgt)
	return cg.defaultgt.LoadFromCsvFile(filename)
}

func (cg *ClassGen) LoadFromCsvIoReader(csvfile io.Reader) error {
	cg.enforceConstraints(&cg.defaultgt)
	return cg.defaultgt.LoadFromCsvIoReader(csvfile)
}

func (cg ClassGen) Roll() (Class, error) {
	cg.enforceConstraints(&cg.defaultgt)
	cs, e := cg.defaultgt.Roll()
	return Class(cs), e
}
