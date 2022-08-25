package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func reqUpgradeFactory(repos Repositories) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
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

		var bodyResource map[string]string
		if err = json.NewDecoder(r.Body).Decode(&bodyResource); err != nil {
			// TODO: what error should we have here??
			handleResponse(w, http.StatusBadRequest, nil, err)
			return
		}

		ok, err := user.UpgradeResource(bodyResource["kind"])
		if err != nil {
			// TODO: what error should we have here??
			handleResponse(w, http.StatusBadRequest, nil, err)
			return
		}

		handleResponse(w, http.StatusOK, fmt.Sprintf("%t", ok), nil)
	}
}