package ciphers

import "unicode"

// Gronsfeld - encrypt/decrypt given text with a key
func Gronsfeld(text, key string, action int) string {
	result := make([]rune, 0, len(text))
	for i, c := range text {
		if unicode.IsLetter(c) {
			shiftedC := rune(int(c) + (int(key[i%len(key)])-48)*action)
			if !unicode.IsLetter(shiftedC) ||
				unicode.IsUpper(c) && unicode.IsLower(shiftedC) ||
				unicode.IsLower(c) && unicode.IsUpper(shiftedC) {
				shiftedC = rune(int(shiftedC) + 26*action*-1)
			}
			result = append(result, shiftedC)
		} else {
			result = append(result, c)
		}
	}
	return string(result)
}
