package ciphers

import "testing"

func TestGronsfeldEncrypt(t *testing.T) {
	var tests = []struct {
		plainText, want, key string
	}{
		{"hello", "jemuq", "2019"},
		{"HeLLo", "JeMUq", "2019"},
		{"w0nderful", "f0wmnaodu", "9"},
		{"hello w0rld", "jennp w0tmf", "20221"},
	}
	for _, test := range tests {
		if got := Gronsfeld(test.plainText, test.key, ENCRYPT); got != test.want {
			t.Errorf(`GronsfeldEncrypt(%q, %q) expected %q but got %q`, test.plainText, test.key, test.want, got)
		}
	}
}

func BenchmarkGronsfeldEncrypt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Gronsfeld("hello w0rld", "20221", ENCRYPT)
	}
}

func TestGronsfeldDecrypt(t *testing.T) {
	var tests = []struct {
		cipherText, want, key string
	}{
		{"jemuq", "hello", "2019"},
		{"JeMUq", "HeLLo", "2019"},
		{"f0wmnaodu", "w0nderful", "9"},
		{"jennp w0tmf", "hello w0rld", "20221"},
	}
	for _, test := range tests {
		if got := Gronsfeld(test.cipherText, test.key, DECRYPT); got != test.want {
			t.Errorf(`GronsfeldDecrypt(%q, %q) expected %q but got %q`, test.cipherText, test.key, test.want, got)
		}
	}
}

func BenchmarkGronsfeldDecrypt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Gronsfeld("jennp w0tmf", "20221", DECRYPT)
	}
}
