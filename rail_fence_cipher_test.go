package main

import "testing"

func TestRailFence(t *testing.T) {
	cases := []struct {
		key            int
		clear, encrypt string
	}{
		{1, "hello", "hello"},
		{5, "hello", "hello"},
		{27, "hello", "hello"},
		{3, "hello", "hoell"},
		{3, "helloplayground", "hoyuelplgrndlelprgdn"},
	}
	for _, c := range cases {
		r := RailFence{c.key}
		if e := r.Encrypt(c.clear); e != c.encrypt {
			t.Errorf("key = %d s = %q RailFence.Encrypt() = %q, want %q", c.key, c.clear, e, c.encrypt)
		}
		if d := r.Decrypt(c.encrypt); d != c.clear {
			t.Errorf("key = %d s = %q RailFence.Decrypt() = %q, want %q", c.key, c.encrypt, d, c.clear)
		}
	}
}
