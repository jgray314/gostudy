package main

import (
	"testing"
)

func TestCombiner_Name(t *testing.T) {
	tests := []struct {
		name string
		c    Combiner
		want string
	}{
		{"Nil Nil", Combiner{Nil{}, Nil{}}, "Combiner{Nil, Nil}"},
		{"Nil Atbash", Combiner{Nil{}, Atbash{}}, "Combiner{Nil, Atbash}"},
		{"Atbash Nil", Combiner{Atbash{}, Nil{}}, "Combiner{Atbash, Nil}"},
		{"Atbash Rot13", Combiner{Atbash{}, Rot13{}}, "Combiner{Atbash, Rot13}"},
		{"Rot13 RailFence4", Combiner{Rot13{}, RailFence{4}}, "Combiner{Rot13, Rail Fence}"},
		{"Atbash Caesar8+RailFence4",
			Combiner{Atbash{}, Combiner{Caesar{8}, RailFence{4}}},
			"Combiner{Atbash, Combiner{Caesar, Rail Fence}}"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Name(); got != tt.want {
				t.Errorf("Combiner.Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCombiner_Encrypt(t *testing.T) {
	defaultstring := "abcdefghij"
	atbashenc, _ := Atbash{}.Encrypt(defaultstring)
	atbashrot13enc, _ := Rot13{}.Encrypt(atbashenc)
	rotenc, _ := Rot13{}.Encrypt(defaultstring)
	rotrailenc, _ := RailFence{4}.Encrypt(rotenc)
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		c       Combiner
		args    args
		want    string
		wantErr bool
	}{
		{"Nil Nil", Combiner{Nil{}, Nil{}}, args{defaultstring}, defaultstring, false},
		{"Nil Atbash", Combiner{Nil{}, Atbash{}}, args{defaultstring}, atbashenc, false},
		{"Atbash Nil", Combiner{Atbash{}, Nil{}}, args{defaultstring}, atbashenc, false},
		{"Atbash Rot13", Combiner{Atbash{}, Rot13{}}, args{defaultstring}, atbashrot13enc, false},
		{"Rot13 RailFence4", Combiner{Rot13{}, RailFence{4}}, args{defaultstring}, rotrailenc, false},
		{"Atbash Caesar8+RailFence4 Long String",
			Combiner{Atbash{}, Combiner{Caesar{8}, RailFence{4}}},
			args{"doingalongeroneoftheseasaproofofconcept"},
			"ewtahtdthtquodpsccfszbuddcphqtfuoubtdtt",
			false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Encrypt(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("Combiner.Encrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Combiner.Encrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCombiner_Decrypt(t *testing.T) {
	defaultstring := "abcdefghij"
	atbashenc, _ := Atbash{}.Encrypt(defaultstring)
	atbashrot13enc, _ := Rot13{}.Encrypt(atbashenc)
	rotenc, _ := Rot13{}.Encrypt(defaultstring)
	rotrailenc, _ := RailFence{4}.Encrypt(rotenc)
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		c       Combiner
		args    args
		want    string
		wantErr bool
	}{
		{"Nil Nil", Combiner{Nil{}, Nil{}}, args{defaultstring}, defaultstring, false},
		{"Nil Atbash", Combiner{Nil{}, Atbash{}}, args{atbashenc}, defaultstring, false},
		{"Atbash Nil", Combiner{Atbash{}, Nil{}}, args{atbashenc}, defaultstring, false},
		{"Atbash Rot13", Combiner{Atbash{}, Rot13{}}, args{atbashrot13enc}, defaultstring, false},
		{"Rot13 RailFence4", Combiner{Rot13{}, RailFence{4}}, args{rotrailenc}, defaultstring, false},
		{"Atbash Caesar8+RailFence4 Long String",
			Combiner{Atbash{}, Combiner{Caesar{8}, RailFence{4}}},
			args{"ewtahtdthtquodpsccfszbuddcphqtfuoubtdtt"},
			"doingalongeroneoftheseasaproofofconcept",
			false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Decrypt(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("Combiner.Decrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Combiner.Decrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}
