package main

import "testing"

func TestAtbash(t *testing.T) {
	cases := []struct {
		in, out string
		// TODO: add error handling tests
	}{
		{"abc", "zyx"},
		{"hello", "svool"},
	}
	a := Atbash{}
	for _, c := range cases {
		enc, _ := a.Encrypt(c.in)
		if enc != c.out {
			t.Errorf("Failed encrypt on %q. Expected:%q Got:%q", c.in, c.out, enc)
		}
		dec, _ := a.Decrypt(c.in)
		if dec != c.out {
			t.Errorf("Failed decrypt on %q. Expected:%q Got:%q", c.in, c.out, dec)
		}
	}
}
