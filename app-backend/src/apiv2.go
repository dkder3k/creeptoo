package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func apiV2Handler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")

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
		case "/api/v2/rot":
			parsedKey, _ := strconv.ParseInt(key[0], 10, 0)
			rotText, err := rot(strings.ToLower(action[0]), text[0], int(parsedKey))
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			response := map[string]string{
				"cipher": "ROT*",
				"action": strings.ToLower(action[0]),
				"text": text[0],
				"result": rotText,
			}
			json.NewEncoder(w).Encode(response)
		case "/api/v2/gronsfeld":
			gronsfeldText, err := gronsfeld(strings.ToLower(action[0]), text[0], strings.ToLower(key[0]))
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			response := map[string]string{
				"cipher": "Gronsfeld cipher",
				"action": strings.ToLower(action[0]),
				"text": text[0],
				"result": gronsfeldText,
			}
			json.NewEncoder(w).Encode(response)
		default:
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		}
	})
}
