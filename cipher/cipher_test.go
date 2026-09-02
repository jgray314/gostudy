package main

import "testing"

func TestNormalize(t *testing.T) {
	cases := []struct {
		in, out string
	}{
		{"Hello, World!", "helloworld"},
		{"already-lower", "alreadylower"},
		{"", ""},
		{"zebra puzzle", "zebrapuzzle"},
	}
	for _, c := range cases {
		if got := Normalize(c.in); got != c.out {
			t.Errorf("Normalize(%q) = %q, want %q", c.in, got, c.out)
		}
	}
}
