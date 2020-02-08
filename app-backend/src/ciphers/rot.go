package ciphers

import "unicode"

// RotEncrypt - encrypt string using ROT cipher
func RotEncrypt(plainText string, key uint8) string {
	key %= 26
	var result []rune
	for _, c := range plainText {
		if unicode.IsLetter(c) {
			shiftedC := rune(uint8(c) + key)
			if !unicode.IsLetter(shiftedC) {
				shiftedC = rune(shiftedC - 26)
			}
			result = append(result, shiftedC)
		} else {
			result = append(result, c)
		}
	}
	return string(result)
}

// RotDecrypt - decrypt ROT cipher encrypted string
func RotDecrypt(cipherText string, key uint8) string {
	key %= 26
	var result []rune
	for _, c := range cipherText {
		if unicode.IsLetter(c) {
			shiftedC := rune(uint8(c) - key)
			if !unicode.IsLetter(shiftedC) {
				shiftedC = rune(shiftedC + 26)
			}
			result = append(result, shiftedC)
		} else {
			result = append(result, c)
		}
	}
	return string(result)
}
