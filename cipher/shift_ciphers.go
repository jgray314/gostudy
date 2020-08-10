package main

import "fmt"

func CondAppend(s string, cond bool, t rune, f rune) string {
	if cond {
		return s + string(t)
	}
	return s + string(f)
}

func shift(s string, shift int) (string, error) {
	r := ""
	for shift < 0 {
		shift += 26
	}
	shift = shift % 26
	for _, a := range s {
		if a < 'a' || a > 'z' {
			return "", fmt.Errorf("Unexpected character in provided string: %q", a)
		}
		r = CondAppend(r, a+rune(shift) > 'z', a+rune(shift)-26, a+rune(shift))
	}
	return r, nil
}

type Caesar struct {
	shift int
}

func (c Caesar) Name() string {
	return "Caesar"
}

func (c Caesar) Encrypt(s string) (string, error) {
	return shift(s, c.shift)
}

func (c Caesar) Decrypt(s string) (string, error) {
	return shift(s, 26-c.shift)
}

type Rot13 struct{}

func (r Rot13) Name() string {
	return "Rot13"
}

func (r Rot13) Encrypt(s string) (string, error) {
	return shift(s, 13)
}

func (r Rot13) Decrypt(s string) (string, error) {
	return shift(s, 13)
}
