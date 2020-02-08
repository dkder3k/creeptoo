package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"./ciphers"
)

func main() {
	http.HandleFunc("/v1/rot", rotHandler)
	http.HandleFunc("/v1/gronsfeld", gronsfeldHandler)
	http.ListenAndServe(":8080", nil)
}

func rotHandler(w http.ResponseWriter, r *http.Request) {

	action, ok := r.URL.Query()["action"]
	if !ok {
		http.Error(w, "Action should be specified", http.StatusBadRequest)
		return
	}
	text, ok := r.URL.Query()["text"]
	if !ok {
		http.Error(w, "Text should be specified", http.StatusBadRequest)
		return
	}
	key, ok := r.URL.Query()["key"]
	if !ok {
		http.Error(w, "Key should be specified", http.StatusBadRequest)
		return
	}

	uintKey, _ := strconv.ParseUint(key[0], 10, 8)

	switch strings.ToLower(action[0]) {
	case "encrypt":
		fmt.Fprintf(w, "%q\n", ciphers.RotEncrypt(text[0], uint8(uintKey)))
	case "decrypt":
		fmt.Fprintf(w, "%q\n", ciphers.RotDecrypt(text[0], uint8(uintKey)))
	default:
		http.Error(w, "Allowed actions: encrypt, decrypt", http.StatusBadRequest)
	}
}

func gronsfeldHandler(w http.ResponseWriter, r *http.Request) {

	action, ok := r.URL.Query()["action"]
	if !ok {
		http.Error(w, "Action should be specified", http.StatusBadRequest)
		return
	}
	text, ok := r.URL.Query()["text"]
	if !ok {
		http.Error(w, "Text should be specified", http.StatusBadRequest)
		return
	}
	key, ok := r.URL.Query()["key"]
	if !ok {
		http.Error(w, "Key should be specified", http.StatusBadRequest)
		return
	}

	switch strings.ToLower(action[0]) {
	case "encrypt":
		fmt.Fprintf(w, "%q\n", ciphers.GronsfeldEncrypt(text[0], key[0]))
	case "decrypt":
		fmt.Fprintf(w, "%q\n", ciphers.GronsfeldDecrypt(text[0], key[0]))
	default:
		http.Error(w, "Allowed actions: encrypt, decrypt", http.StatusBadRequest)
	}
}
