package main

import "fmt"

type Combiner struct {
	outer, inner Cipher
}

func (c Combiner) Name() string {
	return fmt.Sprintf("Combiner{%s, %s}", c.outer.Name(), c.inner.Name())
}

func (c Combiner) Encrypt(s string) (string, error) {
	o, e := c.outer.Encrypt(s)
	if e != nil {
		return "", e
	}
	return c.inner.Encrypt(o)
}

func (c Combiner) Decrypt(s string) (string, error) {
	i, e := c.inner.Decrypt(s)
	if e != nil {
		return "", e
	}
	return c.outer.Decrypt(i)
}
