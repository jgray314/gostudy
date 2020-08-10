package main

import (
	"fmt"
	"strings"
)

// http://practicalcryptography.com/ciphers/classical-era/rail-fence/

type RailFence struct {
	height int
}

func (r RailFence) Name() string {
	return "Rail Fence"
}

func rangeBoundAppend(i int, from, to string) string {
	if i < len(from) {
		return to + string([]rune(from)[i])
	}
	return to
}

func (r RailFence) repLen() int {
	return 2*r.height - 2
}

// Assumes zero based from top
func (r RailFence) getEncryptLevel(l int, s string) string {
	res := ""
	o1, o2 := l, r.repLen()-l
	for b := 0; b < len(s); b += r.repLen() {
		res = rangeBoundAppend(b+o1, s, res)
		if o2 != o1 && o2 != r.repLen() {
			res = rangeBoundAppend(b+o2, s, res)
		}
	}
	return res
}

func (r RailFence) Encrypt(s string) (string, error) {
	if r.height == 1 {
		return s, nil
	}
	var res []string
	for lvl := 0; lvl < r.height; lvl++ {
		res = append(res, r.getEncryptLevel(lvl, s))
	}
	return strings.Join(res, ""), nil
}

// Returns if read rune from the string at the provided index was writeable,
// and if so writes it to the provided index in writeRune and increments the writeIdx. Errors if writeIdx is out of bounds.
func runeAssign(readIdx *int, writeIdx int, readStr string, writeRune []rune) (bool, error) {
	if writeIdx < 0 || writeIdx >= len(writeRune) {
		return false, fmt.Errorf("Out of bounds on runeAssign write index. writeIdx = %d writeRune = %q",
			writeIdx, string(writeRune))
	}
	if *readIdx < 0 || *readIdx >= len(readStr) {
		return false, nil
	}
	writeRune[writeIdx] = []rune(readStr)[*readIdx]
	*readIdx++
	return true, nil
}

func (r RailFence) Decrypt(s string) (string, error) {
	if r.height == 1 {
		return s, nil
	}
	res := make([]rune, len(s))
	di := 0
	for lvl := 0; lvl < r.height; lvl++ {
		o1, o2 := lvl, r.repLen()-lvl
		for b := 0; b+o1 < len(s) && di < len(s); b += r.repLen() {
			if _, err := runeAssign(&di, b+o1, s, res); err != nil {
				return "", err
			}
			if o2 != o1 && o2 != r.repLen() && b+o2 < len(s) && di < len(s) {
				if _, err := runeAssign(&di, b+o2, s, res); err != nil {
					return "", err
				}
			}
		}
	}
	return string(res), nil
}
