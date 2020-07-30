package main

import (
	"fmt"
	"strings"
)

// Normalizes a string to characters in the the set [a-z]
func Normalize(s string) string {
	s = strings.ToLower(s)
	r := ""
	for _, a := range s {
		if a >= 'a' && a < 'z' {
			r = r + string(a)
		} else {
			r = r + " "
		}
	}
	return r
}

func CondAppend(s string, cond bool, t rune, f rune) string {
	if cond {
		return s + string(t)
	}
	return s + string(f)
}


// Funcation ideas from http://practicalcryptography.com/ciphers/classical-era/
func Atbash(s string) string {
	r := ""
	for _, a := range s {
		r = CondAppend(r, a == ' ', ' ', 'z'-a+'a')
	}
	return r
}

func Rot13(s string) string {
	r := ""
	for _, a := range s {
		if a == ' ' {
			r = r + " "
		} else {
			r = CondAppend(r, a > 'm', a-13, a+13)
		}
	}
	return r
}



func main() {
	str := "Hello, playground"
	fmt.Println(str)
	str = Normalize(str)
	fmt.Println(Atbash(str))
	fmt.Println(Rot13(str))
}
