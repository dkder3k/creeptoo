package main

import (
	"./ciphers"
	"errors"
)

func rot(action, text string, key int) (string, error) {
	switch action {
	case "encrypt":
		return ciphers.Rot(text, key, ciphers.ENCRYPT), nil
	case "decrypt":
		return ciphers.Rot(text, key, ciphers.DECRYPT), nil
	default:
		return "", errors.New("Allowed actions: encrypt, decrypt")
	}
}

func gronsfeld(action, text string, key string) (string, error) {
	switch action {
	case "encrypt":
		return ciphers.Gronsfeld(text, key, ciphers.ENCRYPT), nil
	case "decrypt":
		return ciphers.Gronsfeld(text, key, ciphers.DECRYPT), nil
	default:
		return "", errors.New("Allowed actions: encrypt, decrypt")
	}
}