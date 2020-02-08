package ciphers

import "unicode"

// GronsfeldEncrypt - encrypt string using Gronsfeld cipher
func GronsfeldEncrypt(plainText, key string) string {
	var result []rune
	for i, c := range plainText {
		if unicode.IsLetter(c) {
			shiftedC := rune(uint8(c) + uint8(key[i%len(key)]-48))
			if !unicode.IsLetter(shiftedC) {
				shiftedC = rune(shiftedC - 26)
			}
			if unicode.IsUpper(c) && unicode.IsLower(shiftedC) {
				shiftedC = rune(shiftedC - 26)
			}
			result = append(result, shiftedC)
		} else {
			result = append(result, c)
		}
	}
	return string(result)
}

// GronsfeldDecrypt - decrypt Gronsfeld cipher encrypted string
func GronsfeldDecrypt(cipherText, key string) string {
	var result []rune
	for i, c := range cipherText {
		if unicode.IsLetter(c) {
			shiftedC := rune(uint8(c) - uint8(key[i%len(key)]-48))
			if !unicode.IsLetter(shiftedC) {
				shiftedC = rune(shiftedC + 26)
			}
			if unicode.IsLower(c) && unicode.IsUpper(shiftedC) {
				shiftedC = rune(shiftedC + 26)
			}
			result = append(result, shiftedC)
		} else {
			result = append(result, c)
		}
	}
	return string(result)
}
