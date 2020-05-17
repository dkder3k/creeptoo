package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"./ciphers"
)

var listenAddr string

func main() {

	flag.StringVar(&listenAddr, "listen-addr", ":8000", "Address to listen")
	flag.Parse()

	logger := log.New(os.Stdout, "crpt: ", log.LstdFlags)
	logger.Println("Server starting...")

	router := http.NewServeMux()
	router.Handle("/v1/rot", rotHandler())
	router.Handle("/v1/gronsfeld", gronsfeldHandler())

	server := &http.Server{
		Addr:         listenAddr,
		Handler:      logging(logger)(router),
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatalf("Error occurred during server start: %v\n", err)
	}
}

func logging(logger *log.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				logger.Println(r.Method, r.URL.Path, r.RemoteAddr, r.UserAgent())
			}()
			next.ServeHTTP(w, r)
		})
	}
}

func rotHandler() http.Handler {
	return http.HandlerFunc(rot)
}

func gronsfeldHandler() http.Handler {
	return http.HandlerFunc(gronsfeld)
}

func rot(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")

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

	parsedKey, _ := strconv.ParseInt(key[0], 10, 0)

	switch strings.ToLower(action[0]) {
	case "encrypt":
		fmt.Fprintf(w, "%q\n", ciphers.Rot(text[0], int(parsedKey), ciphers.ENCRYPT))
	case "decrypt":
		fmt.Fprintf(w, "%q\n", ciphers.Rot(text[0], int(parsedKey), ciphers.DECRYPT))
	default:
		http.Error(w, "Allowed actions: encrypt, decrypt", http.StatusBadRequest)
	}
}

func gronsfeld(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")

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
		fmt.Fprintf(w, "%q\n", ciphers.Gronsfeld(text[0], key[0], ciphers.ENCRYPT))
	case "decrypt":
		fmt.Fprintf(w, "%q\n", ciphers.Gronsfeld(text[0], key[0], ciphers.DECRYPT))
	default:
		http.Error(w, "Allowed actions: encrypt, decrypt", http.StatusBadRequest)
	}
}
