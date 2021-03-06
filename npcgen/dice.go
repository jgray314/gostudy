package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const UseTime = 0

// Allow the 7 standard dice and a coin flip.
var SupportedSides = []int{2, 4, 6, 8, 10, 12, 20, 100}

type Dice struct {
	detail bool
	r      *rand.Rand
}

func (d *Dice) Detail(detail bool) {
	d.detail = detail
}

// (Re)sets the dice as a random number generator based on seed (if non-zero provided) or the current time.
func (d *Dice) Init(seed int64) {
	if seed == UseTime {
		seed = time.Now().UnixNano()
	}
	d.r = rand.New(rand.NewSource(seed))
}

// Helper function to convert complex string input into
//   num - number of dice to be rolled
//   sides - number of sides of the dice to be rolled
//   plus - additional adjustments to total sum returned after roll
// Returns results of the roll and error messages
// results are undefined behavior if an error is also returned
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

// Returns if sides in the the set of SupportedSides.
func AllowedSides(sides int) bool {
	for _, v := range SupportedSides {
		if v == sides {
			return true
		}
	}
	return false
}

// Simulates a dice roll of form "{}d{}[+-d]"
// Value from 1 to number of sides inclusive
func (d Dice) RollS(s string) (int, error) {
	var num, sides, plus int
	if _, err := scanRoll(s, &num, &sides, &plus); err != nil {
		return 0, err
	}
	var rolls []int
	sum := plus
	for r := 0; r < num; r++ {
		v, e := d.Roll(sides)
		if e != nil {
			return 0, e
		}
		rolls = append(rolls, v)
		sum += v
	}
	if d.detail {
		fmt.Printf("Rolling %v: %v + %d = %d\n", s, rolls, plus, sum)
	}
	return sum, nil
}

// Simulates the roll of a single dice with s sides. Note must be one of the SupportedSides set.
func (d Dice) Roll(s int) (int, error) {
	if !AllowedSides(s) {
		return 0, fmt.Errorf("Attempted number of sides %d not in allowed set %v.\n", s, SupportedSides)
	}
	r := d.r.Intn(s) + 1
	if d.detail {
		fmt.Printf("Rolling %v: %v\n", s, r)
	}
	return r, nil
}
