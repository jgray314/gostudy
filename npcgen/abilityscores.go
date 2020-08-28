package main

import "sort"

type Abilities struct{ strength, dexterity, constitution, intelligence, wisdom, charisma int }

func StandardAbilityValues() []int {
	return []int{15, 14, 13, 12, 10, 8}
}

// Returns a rolled slice of numbers in descending order.
func RollAbilityValues(d Dice) ([]int, error) {
	r := []int{}
	for i := 0; i < 6; i++ {
		v, e := d.RollS("3d6")
		if e != nil {
			return nil, e
		}
		r = append(r, v)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(r)))
	return r, nil
}

// Assigns remaining values to random available ability types
func randomBottom(remain []int, d Dice, a *Abilities) {
	d.r.Shuffle(len(remain), func(i, j int) { remain[i], remain[j] = remain[j], remain[i] })
	idx := 0
	if a.charisma == 0 {
		a.charisma = remain[idx]
		idx++
	}
	if a.constitution == 0 {
		a.constitution = remain[idx]
		idx++
	}
	if a.dexterity == 0 {
		a.dexterity = remain[idx]
		idx++
	}
	if a.intelligence == 0 {
		a.intelligence = remain[idx]
		idx++
	}
	if a.strength == 0 {
		a.strength = remain[idx]
		idx++
	}
	if a.wisdom == 0 {
		a.wisdom = remain[idx]
		idx++
	}
}

func AssignAbilities(c Class, av []int, d Dice) (Abilities, error) {
	a := Abilities{}
	switch c {
	case Barbarian:
		a.strength, a.constitution = av[0], av[1]
	case Bard:
		a.charisma, a.dexterity = av[0], av[1]
	case Cleric:
		v, e := d.Roll(2)
		if e != nil {
			return Abilities{}, e
		}
		if v == 0 {
			a.wisdom, a.strength = av[0], av[1]
		} else {
			a.wisdom, a.constitution = av[0], av[1]
		}
	case Druid:
		a.wisdom, a.constitution = av[0], av[1]
	case Fighter:
		v, e := d.Roll(2)
		if e != nil {
			return Abilities{}, e
		}
		if v == 0 {
			a.strength, a.constitution = av[0], av[1]
		} else {
			a.dexterity, a.constitution = av[0], av[1]
		}
	case Monk:
		a.dexterity, a.wisdom = av[0], av[1]
	case Paladin:
		a.strength, a.charisma = av[0], av[1]
	case Ranger:
		a.dexterity, a.wisdom = av[0], av[1]
	case Rogue:
		v, e := d.Roll(2)
		if e != nil {
			return Abilities{}, e
		}
		if v == 0 {
			a.dexterity, a.intelligence = av[0], av[1]
		} else {
			a.dexterity, a.charisma = av[0], av[1]
		}
	case Sorcerer:
		a.charisma, a.constitution = av[0], av[1]
	case Warlock:
		a.charisma, a.constitution = av[0], av[1]
	case Wizard:
		v, e := d.Roll(2)
		if e != nil {
			return Abilities{}, e
		}
		if v == 0 {
			a.intelligence, a.constitution = av[0], av[1]
		} else {
			a.intelligence, a.dexterity = av[0], av[1]
		}
	}
	randomBottom(av[2:], d, &a)
	return a, nil
}

// TODO: subclass adjustments
func AdjustAbilities(r Race, a *Abilities) {
	switch r {
	case Dwarf:
		a.constitution += 2
	case Elf:
		a.dexterity += 2
	case Halfling:
		a.dexterity += 2
	case Human:
		a.charisma += 1
		a.constitution += 1
		a.dexterity += 1
		a.intelligence += 1
		a.strength += 1
		a.wisdom += 1
	case Dragonborn:
		a.strength += 2
		a.charisma += 1
	case Gnome:
		a.intelligence += 2
	case HalfElf:
		a.charisma += 2
		a.dexterity += 1
		a.wisdom += 1
	case HalfOrc:
		a.strength += 2
		a.constitution += 1
	case Tiefling:
		a.charisma += 2
		a.intelligence += 1
	}
}

func AbilityValueToModifier(n int) int {
	return (n / 2) - 5
}
