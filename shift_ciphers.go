package main

func shift(s string, shift int) string {
	r := ""
	for shift < 0 {
		shift += 26
	}	
	shift = shift % 26
	for _, a := range s {
		r = cond_append(r, a+rune(shift) > 'z', a+rune(shift)-26, a+rune(shift))
	}
	return r
}

type caesar struct {
	shift int
}

func (c caesar) name() string {
	return "ceasar"
}

func (c caesar) encrypt(s string) string {
	return shift(s, c.shift)
}

func (c caesar) decrypt(s string) string {
	return shift(s, 26 - c.shift)
}

type rot13 struct {}

func (r rot13) name() string {
	return "rot13"
}

func (r rot13) encrypt(s string) string {
	return shift(s, 13)
}

func (r rot13) decrypt(s string) string {
	return shift(s, 13)
}
