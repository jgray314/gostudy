package main

import (
	"fmt"
	"strings"
)

// Normalizes a string to characters in the the set [a-z]
// Common step for classical crypto
func Normalize(s string) string {
	s = strings.ToLower(s)
	r := ""
	for _, a := range s {
		if a >= 'a' && a < 'z' {
			r = r + string(a)
		}
	}
	return r
}

func CondAppend(s string, cond bool, t rune, f rune) string {
	if cond {
		return s + string(t)
	}
	return s + string(f)
}

// Funcation ideas from http://practicalcryptography.com/ciphers/classical-era/
func Atbash(s string) string {
	r := ""
	for _, a := range s {
		r = r + string('z'-a+'a')
	}
	return r
}

func Caesar(s string, shift int) string {
	r := ""
	shift = shift % 26 // More adjustments needed for negative shifts
	for _, a := range s {
		r = CondAppend(r, a+rune(shift) > 'z', a+rune(shift)-26, a+rune(shift))
	}
	return r
}

func Rot13(s string) string {
	return Caesar(s, 13)
}

func main() {
	str := "Hello, playground"
	fmt.Println(str)
	str = Normalize(str)
	fmt.Println(str)
	fmt.Println(Atbash(str))
	fmt.Println(Caesar(str, 1))
	fmt.Println(Rot13(str))
}
