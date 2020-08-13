package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// Allow the 7 standard dice and a coin flip.
var allowedsides = []int{2, 4, 6, 8, 10, 12, 20, 100}

type Dice struct {
	seed  int64
	print bool
	r     *rand.Rand
}

// Resets the dice as a random number generator based on seed (if provided) or the current time.
func (d *Dice) Init() {
	if d.seed == 0 {
		d.seed = time.Now().UnixNano()
	}
	d.r = rand.New(rand.NewSource(d.seed))
}

func scanRoll(s string, num, sides, plus *int) (int, error) {
	*sides = 0
	if s[0] == 'd' {
		*num = 1
		if strings.Contains(s, "+") {
			return fmt.Sscanf(s, "d%d+%d", sides, plus)
		}
		if strings.Contains(s, "-") {
			n, err := fmt.Sscanf(s, "d%d-%d", sides, plus)
			*plus *= -1
			return n, err
		}
		*plus = 0
		return fmt.Sscanf(s, "d%d", sides)
	}
	if strings.Contains(s, "+") {
		return fmt.Sscanf(s, "%dd%d+%d", num, sides, plus)
	}
	if strings.Contains(s, "-") {
		n, err := fmt.Sscanf(s, "%dd%d-%d", num, sides, plus)
		*plus *= -1
		return n, err
	}
	*plus = 0
	return fmt.Sscanf(s, "%dd%d", num, sides)
}

func allowedSides(sides int) bool {
	for _, v := range allowedsides {
		if v == sides {
			return true
		}
	}
	return false
}

// Simulates a dice roll of form "{}d{}[+-d]"
// Value from 1 to number of sides inclusive
func (d Dice) Roll(s string) (int, error) {
	var num, sides, plus int
	if _, err := scanRoll(s, &num, &sides, &plus); err != nil {
		return 0, err
	}
	if !allowedSides(sides) {
		return 0, fmt.Errorf("Attempted number of sides %d not in allowed set %v.", sides, allowedsides)
	}
	var rolls []int
	sum := plus
	for r := 0; r < num; r++ {
		v := d.r.Intn(sides) + 1
		rolls = append(rolls, v)
		sum += v
	}
	if d.print {
		fmt.Printf("Rolling %v: %v + %d = %d\n", s, rolls, plus, sum)
	}
	return sum, nil
}

// temp for demo, next unittest
func main() {
	d := Dice{seed: 0, print: true}
	d.Init()
	d.Roll("3d6")
	d.Roll("1d20")
	d.Roll("1d20+3")
}
