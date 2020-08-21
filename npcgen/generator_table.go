package main

import (
	"fmt"
	"strings"
)

type TypeValidator interface {
	IsValid(string) bool
	GetSupported() []string
}

type StringTypeValidator struct{}

func (stv StringTypeValidator) IsValid(s string) bool  { return true }
func (stv StringTypeValidator) GetSupported() []string { return []string{"all strings allowed"} }

type GeneratorTable struct {
	name          string
	size          int
	detail        bool
	table         []string
	typevalidator TypeValidator
	dice          Dice
}

func (gt GeneratorTable) validateSize() error {
	if !AllowedSides(gt.size) {
		return fmt.Errorf("Table is not of a size that can be roll. Got %d, Expected one of %v\n", gt.size, SupportedSides)
	}
	len := len(gt.table)
	if len != gt.size {
		return fmt.Errorf("Invalid table size. Got %d, Expected %d.\n", len, gt.size)
	}
	return nil
}

/* Takes a comma separated string of number of entries and then the entry string value. Must sum to configured size.
* 	"50 human, 25 elf, 5 halfelf, 20 dwarf"
*	"2 something, 3 something else, 1 nothing"
 */
func (gt *GeneratorTable) LoadFromString(s string) error {
	if gt.detail {
		fmt.Printf("Loading from table %v from string: %v\n", gt.name, s)
	}
	// Start by clearing
	gt.table = nil
	for _, p := range strings.Split(s, ",") {
		p = strings.TrimSpace(p)
		idx := strings.Index(p, " ")
		if idx == -1 {
			return fmt.Errorf("Invalided comma separated value. Need a number followed by a space and a string. Got: %q\n", p)
		}
		var n int
		_, e := fmt.Sscanf(p[:idx], "%d", &n)
		if e != nil {
			return fmt.Errorf("Bad number in pair for %q. %s\n", p[:idx], e.Error())
		}
		r := strings.TrimSpace(strings.ToLower(p[idx:]))
		if !gt.typevalidator.IsValid(r) {
			return fmt.Errorf("Unsupported value: %v. Expected one of %v\n", r, gt.typevalidator.GetSupported())
		}
		for i := 0; i < n; i++ {
			gt.table = append(gt.table, r)
		}
	}
	if gt.detail {
		fmt.Println("Loaded table to be validated:\n", gt.table)
	}
	return gt.validateSize()
}

func (gt GeneratorTable) Roll() (string, error) {
	if e := gt.validateSize(); e != nil {
		return "", fmt.Errorf("Attempt to roll against improperly initialized table. %s\n", e.Error())
	}
	v, e := gt.dice.Roll(gt.size)
	if e != nil {
		return "", e
	}
	if gt.detail {
		fmt.Printf("Rolled %d of %d, returning %s\n", v, gt.size, gt.table[v-1])
	}
	return gt.table[v-1], nil
}

func (gt GeneratorTable) PrintState() {
	fmt.Println(gt)
}
