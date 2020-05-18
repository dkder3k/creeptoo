package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type task struct {
	Action	string `json:"action"`
	Key		string `json:"key"`
	Text	string `json:"text"`
}

func formTask(request []byte) (task, error) {
	var t task
	if err := json.Unmarshal(request, &t); err != nil {
		return task{}, err
	}
	return t, nil
}

func apiV3Handler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json")

		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		currentTask, err := formTask(reqBody)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if currentTask.Action == "" {
			http.Error(w, "Action should be specified", http.StatusBadRequest)
			return
		}
		if currentTask.Text == "" {
			http.Error(w, "Text should be specified", http.StatusBadRequest)
			return
		}
		if currentTask.Key == "" {
			http.Error(w, "Key should be specified", http.StatusBadRequest)
			return
		}


		switch r.URL.Path {
		case "/api/v3/rot":
			parsedKey, _ := strconv.ParseInt(currentTask.Key, 10, 0)
			rotText, err := rot(strings.ToLower(currentTask.Action), currentTask.Text, int(parsedKey))
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			response := map[string]string{
				"cipher": "ROT*",
				"action": strings.ToLower(currentTask.Action),
				"text": currentTask.Text,
				"result": rotText,
			}
			json.NewEncoder(w).Encode(response)
		case "/api/v3/gronsfeld":
			gronsfeldText, err := gronsfeld(strings.ToLower(currentTask.Action), currentTask.Text, strings.ToLower(currentTask.Key))
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			response := map[string]string{
				"cipher": "Gronsfeld cipher",
				"action": strings.ToLower(currentTask.Action),
				"text": currentTask.Text,
				"result": gronsfeldText,
			}
			json.NewEncoder(w).Encode(response)
		default:
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		}
	})
}
