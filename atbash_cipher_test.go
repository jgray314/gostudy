package main

import "testing"

func TestAtbash(t *testing.T) {
	cases := []struct {
		in, out string
	}{
		{"abc", "zyx"},
		{"hello", "svool"},
	}
	a := atbash{}
	for _, c := range cases {
		e := a.encrypt(c.in)
		if e != c.out {
			t.Errorf("Failed encrypt on %q. Expected:%q Got:%q", c.in, c.out, e)
		}
		d := a.encrypt(c.in)
		if d != c.out {
			t.Errorf("Failed decrypt on %q. Expected:%q Got:%q", c.in, c.out, d)
		}
	}
}
