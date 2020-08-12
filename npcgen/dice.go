package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Allow the 7 standard dice and a coin flip.
//const allowedsides []int = []int{2,4,6,8,10,12,20,100}

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

// Simulates a dice roll of form "{}d{}[+d]"
func (d Dice) Roll(s string) (int, error) {
	num, sides, plus := 0, 0, 0
	n, err := fmt.Sscanf(s, "%dd%d+%d", &num, &sides, &plus)
	if n < 2 {
		return 0, err
	}
	//	if !Find(allowedsides, sides) {
	//		return 0, fmt.Errorf("Attempted number of sides %d not in allowed set %s.", sides, allowedsides)
	//	}
	var rolls []int
	sum := plus
	for r := 0; r < num; r++ {
		v := d.r.Intn(sides)
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
