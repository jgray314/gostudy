package main

import (
	"testing"
)

func Test_runeAssign(t *testing.T) {
	type args struct {
		readIdx   int
		writeIdx  int
		readStr   string
		writeRune []rune
	}
	tests := []struct {
		name           string
		args           args
		readIdxFinal   int
		writeRuneFinal []rune
		want           bool
		wantErr        bool
	}{
		{"Simple 0->0", args{0, 0, "abc", []rune("   ")}, 1, []rune("a  "), true, false},
		{"Simple 2->1", args{2, 1, "abc", []rune("   ")}, 3, []rune(" c "), true, false},
		{"Input OutofBounds", args{3, 0, "abc", []rune("   ")}, 3, []rune("   "), false, false},
		{"Output OutofBounds", args{0, 4, "abc", []rune("   ")}, 0, []rune("   "), false, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := runeAssign(&tt.args.readIdx, tt.args.writeIdx, tt.args.readStr, tt.args.writeRune)
			if (err != nil) != tt.wantErr {
				t.Errorf("runeAssign() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("runeAssign() = %v, want %v", got, tt.want)
			}
			if string(tt.args.readIdx) != string(tt.readIdxFinal) {
				t.Errorf("runeAssign updated readIdx to %d, wanted %d", tt.args.readIdx, tt.readIdxFinal)
			}
			if string(tt.args.writeRune) != string(tt.writeRuneFinal) {
				t.Errorf("runeAssign updated writeRune to %v, wanted %v",
					string(tt.args.writeRune), string(tt.writeRuneFinal))
			}
		})
	}
}

func TestRailFence_EngcrypDecrypt(t *testing.T) {
	cases := []struct {
		key            int
		clear, encrypt string
	}{
		{1, "hello", "hello"},
		{5, "hello", "hello"},
		{27, "hello", "hello"},
		{3, "hello", "hoell"},
		{3, "helloplayground", "hoyuelpagonllrd"},
		{4, "helloplayground", "hluepaonloyrdlg"},
		{7, "helloplayground", "hueonlrdlgoypal"},
		{2, "abcdefg", "acegbdf"},
		{3, "abcdefg", "aebdfcg"},
	}
	for _, c := range cases {
		r := RailFence{c.key}
		if enc, _ := r.Encrypt(c.clear); enc != c.encrypt {
			t.Errorf("key = %d s = %q RailFence.Encrypt() = %q, want %q", c.key, c.clear, enc, c.encrypt)
		}
		if dec, _ := r.Decrypt(c.encrypt); dec != c.clear {
			t.Errorf("key = %d s = %q RailFence.Decrypt() = %q, want %q", c.key, c.encrypt, dec, c.clear)
		}
	}
}
