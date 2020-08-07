package main

import (
	"fmt"
	"strings"
)

// Can assume input string with charcters only in set [a-z]
type Cipher interface {
	Name() string
	Encrypt(string) string
	Decrypt(string) string
}

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

func PrintDemo(c Cipher, s string) {
	n := Normalize(s)
	e := c.Encrypt(n)
	fmt.Printf("Cypher: %s\tOriginal: %s\tNormalized: %s\tEncrypt: %s\tDecrypt: %s\n", c.Name(), s, n, e, c.Decrypt(e))
}

// Function ideas from http://practicalcryptography.com/ciphers/classical-era/

func main() {
	s := "Hello, playground"
	PrintDemo(Atbash{}, s)
	PrintDemo(Caesar{shift: 3}, s)
	PrintDemo(Rot13{}, s)
	PrintDemo(RailFence{height: 3}, s)
}
