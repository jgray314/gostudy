package main

import "testing"

func TestRot13CaesarMatch(t *testing.T) {
	cases := []struct {
		in string
	}{
		{"helloworld"},
		{"something"},
		{"abcdxyz"},
	}
	c13 := caesar{shift: 13}
	r := rot13{}
	for _, c := range cases {
		ce := c13.encrypt(c.in)
		re := r.encrypt(c.in)
		if ce != re {
			t.Errorf("Unmatched encrypt for value %q: Caesar 13 = %q, Rot13 %q", c.in, ce, re)
		}
		cd := c13.decrypt(ce)
		rd := r.decrypt(re)
		if cd != rd {
			t.Errorf("Unmatched decrypt for value %q: Caesar 13 = %q, Rot13 %q", c.in, cd, rd)
		}
		if cd != c.in {
			t.Errorf("Failed decrypt. Wanted %q, Got:%q", cd, c.in)
		}
	}
}

func TestCaesarEncryptDecrypt(t *testing.T) {
	cases := []struct {
		dec string
		sh  int
		enc string
	}{
		{"abc", 1, "bcd"},
		{"abc", 2, "cde"},
		{"abcdxyz", -1, "zabcwxy"},
		{"abc", 27, "bcd"},
		{"abcdxyz", 0, "abcdxyz"},
		{"abcdxyz", 26, "abcdxyz"},
		{"abcdxyz", -26, "abcdxyz"},
	}
	for _, c := range cases {
		enc := caesar{shift: c.sh}.encrypt(c.dec)
		if enc != c.enc {
			t.Errorf("Failed encryption with input %q shift %d. Wanted %q, Got:%q", c.dec, c.sh, c.enc, enc)
		}
		dec := caesar{shift: c.sh}.decrypt(c.enc)
		if dec != c.dec {
			t.Errorf("Failed decryption with input %q shift %d. Wanted %q, Got:%q", c.enc, c.sh, c.dec, dec)
		}
	}
}

func TestCaesarDecrypt(t *testing.T) {
	cases := []struct {
		in, out string
	}{
		{"abc", "nop"},
		{"xyz", "klm"},
		{"hello", "uryyb"},
	}
	r := rot13{}
	for _, c := range cases {
		e := r.encrypt(c.in)
		if e != c.out {
			t.Errorf("Failed encryption with input %q. Wanted %q, Got:%q", c.in, c.out, e)
		}
		d := r.decrypt(c.in)
		if d != c.out {
			t.Errorf("Failed decryption with input %q. Wanted %q, Got:%q", c.in, c.out, d)
		}
	}
}
