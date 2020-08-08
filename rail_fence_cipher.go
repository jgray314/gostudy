package main

import (
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

// Assumes zero based from top
func (r RailFence) getEncryptLevel(l int, s string) string {
	repLen := 2*r.height - 2
	res := ""
	o1, o2 := l, repLen-l
	for b := 0; b < len(s); b += repLen {
		res = rangeBoundAppend(b+o1, s, res)
		if o2 != o1 && o2 != repLen {
			res = rangeBoundAppend(b+o2, s, res)
		}
	}
	return res
}

func (r RailFence) Encrypt(s string) string {
	if r.height == 1 {
		return s
	}
	var res []string
	for lvl := 0; lvl < r.height; lvl++ {
		res = append(res, r.getEncryptLevel(lvl, s))
	}
	return strings.Join(res, "")
}

func (r RailFence) Decrypt(s string) string {
	return s
}
