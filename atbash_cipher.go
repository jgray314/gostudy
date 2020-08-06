package main

type Atbash struct{}

func (a Atbash) Name() string {
	return "Atbash"
}
func (a Atbash) Encrypt(s string) string {
	r := ""
	for _, a := range s {
		r = r + string('z'-a+'a')
	}
	return r
}

func (a Atbash) Decrypt(s string) string {
	return a.Encrypt(s)
}