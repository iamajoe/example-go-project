package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func reqCreateUser(repos Repositories) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "not allwowed", http.StatusMethodNotAllowed)
			return
		}

		var user User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			// TODO: what error should we have here??
			handleResponse(w, http.StatusBadRequest, nil, err)
			return
		}

		ok, err := createUser(user.Username, repos.user, repos.factory)
		if err != nil {
			// TODO: what error should we have here??
			handleResponse(w, http.StatusBadRequest, nil, err)
			return
		}

		handleResponse(w, http.StatusOK, fmt.Sprintf("%t", ok), nil)
	}
}

func reqGetDashboard(repos Repositories) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "not allwowed", http.StatusMethodNotAllowed)
			return
		}

		username := r.URL.Query().Get("username")
		user, err := getUserByUsername(username, repos.user, repos.factory)
		if err != nil {
			// TODO: what error should we have here??
			handleResponse(w, http.StatusBadRequest, nil, err)
			return
		}

		handleResponse(w, http.StatusOK, user, nil)
	}
}
