package main

type atbash struct{}

func (a atbash) name() string {
	return "atbash"
}
func (a atbash) encrypt(s string) string {
	r := ""
	for _, a := range s {
		r = r + string('z'-a+'a')
	}
	return r
}

func (a atbash) decrypt(s string) string {
	return a.encrypt(s)
}