package main

import (
	"fmt"
	"strings"
)

// Can assume input string with charcters only in set [a-z]
type Cipher interface {
	Name() string
	Encrypt(string) (string, error)
	Decrypt(string) (string, error)
}

// Normalizes a string to characters in the the set [a-z]
// Common step for classical crypto.
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

func PrintDemo(c Cipher, s string) error {
	n := Normalize(s)
	enc, e := c.Encrypt(n)
	if e != nil {
		return e
	}
	dec, e := c.Decrypt(enc)
	if e != nil {
		return e
	}
	fmt.Printf("Cypher: %s\n\tOriginal: %q\n\tNormalized: %q\n\tEncrypt: %q\n\tDecrypt: %q\n", c.Name(), s, n, enc, dec)
	return nil
}

// Function ideas from http://practicalcryptography.com/ciphers/classical-era/

func main() {
	s := "Hello, playground. I am lengthening this message to checkout more techniques."
	PrintDemo(Atbash{}, s)
	PrintDemo(Caesar{shift: 3}, s)
	PrintDemo(Rot13{}, s)
	PrintDemo(RailFence{height: 3}, s)
	PrintDemo(Combiner{Caesar{7}, RailFence{5}}, s)
}
