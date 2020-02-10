package ciphers

import "unicode"

// Rot - encrypt/decrypt given text with a key
func Rot(text string, key int, action int) string {
	key %= 26
	var result []rune
	for _, c := range text {
		if unicode.IsLetter(c) {
			shiftedC := rune(int(c) + key * action)
			if !unicode.IsLetter(shiftedC) {
				shiftedC = rune(int(shiftedC) + 26 * action * -1)
			}
			result = append(result, shiftedC)
		} else {
			result = append(result, c)
		}
	}
	return string(result)
}
