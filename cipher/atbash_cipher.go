package main

import "fmt"

type Atbash struct{}

func (a Atbash) Name() string {
	return "Atbash"
}
func (a Atbash) Encrypt(s string) (string, error) {
	r := ""
	for _, a := range s {
		if a < 'a' || a > 'z' {
			return "", fmt.Errorf("Unexpected character in provided string: %q", a)
		}
		r = r + string('z'-a+'a')
	}
	return r, nil
}

func (a Atbash) Decrypt(s string) (string, error) {
	return a.Encrypt(s)
}
