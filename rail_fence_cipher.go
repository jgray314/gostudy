package main

// http://practicalcryptography.com/ciphers/classical-era/rail-fence/

type RailFence struct {
	height int
}

func (r RailFence) Name() string {
	return "Rail Fence"
}

func (r RailFence) Encrypt(s string) string {
	if r.height == 1 {
		return s
	}
	repLen := 2*r.height - 2
	sLen := len(s)
	res := ""
	for mod := 0; mod < r.height; mod++ {
		for mul := 0; mul < sLen; mul += repLen {
			idx1 := mul + mod
			idx2 := mul + repLen - mod - 1
			if idx1 < sLen {
				res = res + string([]rune(s)[idx1])
			}
			// Don't want to double list the top or bottom of the fence or overrun the string.
			if mod != 0 && idx2 != idx1 && idx2 < sLen {
				res = res + string([]rune(s)[idx2])
			}
		}
	}
	return res
}

func (r RailFence) Decrypt(s string) string {
	return s
}
