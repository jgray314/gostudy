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
	c13 := Caesar{shift: 13}
	r := Rot13{}
	for _, c := range cases {
		ce := c13.Encrypt(c.in)
		re := r.Encrypt(c.in)
		if ce != re {
			t.Errorf("Unmatched encrypt for value %q: Caesar 13 = %q, Rot13 %q", c.in, ce, re)
		}
		cd := c13.Decrypt(ce)
		rd := r.Decrypt(re)
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
		enc := Caesar{shift: c.sh}.Encrypt(c.dec)
		if enc != c.enc {
			t.Errorf("Failed encryption with input %q shift %d. Wanted %q, Got:%q", c.dec, c.sh, c.enc, enc)
		}
		dec := Caesar{shift: c.sh}.Decrypt(c.enc)
		if dec != c.dec {
			t.Errorf("Failed decryption with input %q shift %d. Wanted %q, Got:%q", c.enc, c.sh, c.dec, dec)
		}
	}
}

func TestRo13EncryptDecrypt(t *testing.T) {
	cases := []struct {
		in, out string
	}{
		{"abc", "nop"},
		{"xyz", "klm"},
		{"hello", "uryyb"},
	}
	r := Rot13{}
	for _, c := range cases {
		e := r.Encrypt(c.in)
		if e != c.out {
			t.Errorf("Failed encryption with input %q. Wanted %q, Got:%q", c.in, c.out, e)
		}
		d := r.Decrypt(c.in)
		if d != c.out {
			t.Errorf("Failed decryption with input %q. Wanted %q, Got:%q", c.in, c.out, d)
		}
	}
}
