package main

import "testing"

func TestRot13CaesarMatch(t *testing.T) {
	cases := []struct {
		in string
		// TODO: Add error handling testcases
	}{
		{"helloworld"},
		{"something"},
		{"abcdxyz"},
	}
	c13 := Caesar{shift: 13}
	r := Rot13{}
	for _, c := range cases {
		cenc, _ := c13.Encrypt(c.in)
		renc, _ := r.Encrypt(c.in)
		if cenc != renc {
			t.Errorf("Unmatched encrypt for value %q: Caesar 13 = %q, Rot13 %q", c.in, cenc, renc)
		}
		cdec, _ := c13.Decrypt(cenc)
		rdec, _ := r.Decrypt(renc)
		if cdec != rdec {
			t.Errorf("Unmatched decrypt for value %q: Caesar 13 = %q, Rot13 %q", c.in, cdec, rdec)
		}
		if cdec != c.in {
			t.Errorf("Failed decrypt. Wanted %q, Got:%q", cdec, c.in)
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
		enc, _ := Caesar{shift: c.sh}.Encrypt(c.dec)
		if enc != c.enc {
			t.Errorf("Failed encryption with input %q shift %d. Wanted %q, Got:%q", c.dec, c.sh, c.enc, enc)
		}
		dec, _ := Caesar{shift: c.sh}.Decrypt(c.enc)
		if dec != c.dec {
			t.Errorf("Failed decryption with input %q shift %d. Wanted %q, Got:%q", c.enc, c.sh, c.dec, dec)
		}
	}
}

func TestRo13EncryptDecrypt(t *testing.T) {
	cases := []struct {
		in, out string
		// TODO: Add in error handling cases
	}{
		{"abc", "nop"},
		{"xyz", "klm"},
		{"hello", "uryyb"},
	}
	r := Rot13{}
	for _, c := range cases {
		enc, _ := r.Encrypt(c.in)
		if enc != c.out {
			t.Errorf("Failed encryption with input %q. Wanted %q, Got:%q", c.in, c.out, enc)
		}
		dec, _ := r.Decrypt(c.in)
		if dec != c.out {
			t.Errorf("Failed decryption with input %q. Wanted %q, Got:%q", c.in, c.out, dec)
		}
	}
}
