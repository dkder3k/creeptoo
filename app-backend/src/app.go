package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"time"
)

var listenAddr string

func main() {

	flag.StringVar(&listenAddr, "listen-addr", ":8000", "Address to listen")
	flag.Parse()

	logger := log.New(os.Stdout, "crpt: ", log.LstdFlags)
	logger.Println("Server starting...")

	router := http.NewServeMux()
	router.Handle("/api/v1/", apiV1Handler())
	router.Handle("/api/v2/", apiV2Handler())
	router.Handle("/api/v3/", apiV3Handler())
	router.Handle("/health", health())

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

func health() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})
}
