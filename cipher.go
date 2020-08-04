package main

import (
	"fmt"
	"strings"
)

type cipher interface {
	name() string
	encrypt(string) string
	decrypt(string) string
}

// Normalizes a string to characters in the the set [a-z]
// Common step for classical crypto
func normalize(s string) string {
	s = strings.ToLower(s)
	r := ""
	for _, a := range s {
		if a >= 'a' && a < 'z' {
			r = r + string(a)
		}
	}
	return r
}

func print_demo(c cipher, s string) {
	n := normalize(s)
	e := c.encrypt(n)
	fmt.Printf("Cypher: %s\tOriginal: %s\tNormalized: %s\tEncrypt: %s\tDecrypt: %s\n", c.name(), s, n, e, c.decrypt(e))
}

func cond_append(s string, cond bool, t rune, f rune) string {
	if cond {
		return s + string(t)
	}
	return s + string(f)
}

// Funcation ideas from http://practicalcryptography.com/ciphers/classical-era/
type atbash struct{}

func (a atbash) name() string {
	return "atbash"
}
func (a atbash) encrypt(s string) string {
	r := ""
	for _, a := range s {
		r = r + string('z'-a+'a')
	}
	return r
}

func (a atbash) decrypt(s string) string {
	return a.encrypt(s)
}

func main() {
	s := "Hello, playground"
	print_demo(atbash{}, s)
	print_demo(caesar{shift: 3}, s)
	print_demo(rot13{}, s)
}
