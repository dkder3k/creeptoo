package ciphers

import "testing"

func TestROTEncrypt(t *testing.T) {
	var tests = []struct {
		plainText, want string
		key             int
	}{
		{"HelLo", "LipPs", 4},
		{"Hello", "Hello", 0},
		{"Hello", "Hello", 26},
		{"Hello", "Ifmmp", 53},
		{"Hello", "Jgnnq", 54},
		{"hello w0rld", "olssv d0ysk", 7},
	}
	for _, test := range tests {
		if got := Rot(test.plainText, test.key, ENCRYPT); got != test.want {
			t.Errorf(`RotEncrypt(%q, %d) expected %q but got %q`, test.plainText, test.key, test.want, got)
		}
	}
}

func TestROTDecrypt(t *testing.T) {
	var tests = []struct {
		cipherText, want string
		key              int
	}{
		{"LiPps", "HeLlo", 4},
		{"Hello", "Hello", 0},
		{"Hello", "Hello", 26},
		{"Ifmmp", "Hello", 53},
		{"Jgnnq", "Hello", 54},
		{"olssv d0ysk", "hello w0rld", 7},
	}
	for _, test := range tests {
		if got := Rot(test.cipherText, test.key, DECRYPT); got != test.want {
			t.Errorf(`RotDecrypt(%q, %d) expected %q but got %q`, test.cipherText, test.key, test.want, got)
		}
	}
}
