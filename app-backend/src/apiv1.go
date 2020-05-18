package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func apiV1Handler() http.Handler {
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
			fmt.Fprintln(w, rotText)
		case "/api/v1/gronsfeld":
			gronsfeldText, err := gronsfeld(strings.ToLower(action[0]), text[0], strings.ToLower(key[0]))
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			fmt.Fprintln(w, gronsfeldText)
		default:
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		}
	})
}
