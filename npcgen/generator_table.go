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
		return fmt.Errorf("Table is not of a size that can be roll. Got %d, Expected one of %v", gt.size, SupportedSides)
	}
	len := len(gt.table)
	if len != gt.size {
		return fmt.Errorf("Invalid table size. Got %d, Expected %d./n", len, gt.size)
	}
	return nil
}

func (gt *GeneratorTable) LoadFromString(s string) error {
	if gt.detail {
		fmt.Printf("Loading from table %v from string: %v", gt.name, s)
	}
	// Start by clearing
	gt.table = nil
	for _, p := range strings.Split(s, ",") {
		var n int
		var r string
		_, e := fmt.Sscanf(p, "%d %s", &n, &r)
		if e != nil {
			return e
		}
		r = strings.ToLower(r)
		if !gt.typevalidator.IsValid(r) {
			return fmt.Errorf("Unsupported value: %v. Expected one of %v", r, gt.typevalidator.GetSupported())
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
	v, e := gt.dice.Roll(gt.size)
	if e != nil {
		return "", e
	}
	return gt.table[v-1], nil
}

func (gt GeneratorTable) PrintState() {
	fmt.Println(gt)
}
