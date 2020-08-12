package main

type Nil struct{}

func (n Nil) Name() string {
	return "Nil"
}

func (n Nil) Encrypt(s string) (string, error) {
	return s, nil
}

func (n Nil) Decrypt(s string) (string, error) {
	return s, nil
}
