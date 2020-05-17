package main

import (
	"flag"
	"fmt"
	"errors"
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
	router.Handle("/api/v1/", apiHandler())

	server := &http.Server{
		Addr:         listenAddr,
		Handler:      logging(logger)(router),
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil {
		logger.Fatalf("Fatal error occurred: %v\n", err)
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

func apiHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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


		switch r.URL.Path {
		case "/api/v1/rot":
			parsedKey, _ := strconv.ParseInt(key[0], 10, 0)
			rotText, err := rot(strings.ToLower(action[0]), text[0], int(parsedKey))
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			fmt.Fprint(w, rotText)
		case "/api/v1/gronsfeld":
			gronsfeldText, err := gronsfeld(strings.ToLower(action[0]), text[0], strings.ToLower(key[0]))
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			fmt.Fprint(w, gronsfeldText)
		default:
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		}
	})
}

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
