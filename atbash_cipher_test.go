package main

import "testing"

func TestAtbash(t *testing.T) {
	cases := []struct {
		in, out string
	}{
		{"abc", "zyx"},
		{"hello", "svool"},
	}
	a := Atbash{}
	for _, c := range cases {
		e := a.Encrypt(c.in)
		if e != c.out {
			t.Errorf("Failed encrypt on %q. Expected:%q Got:%q", c.in, c.out, e)
		}
		d := a.Decrypt(c.in)
		if d != c.out {
			t.Errorf("Failed decrypt on %q. Expected:%q Got:%q", c.in, c.out, d)
		}
	}
}
