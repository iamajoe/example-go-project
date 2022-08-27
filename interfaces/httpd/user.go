package httpd

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/joesantosio/simple-game-api/domains/user"
	"github.com/joesantosio/simple-game-api/infrastructure"
)

func reqCreateUser(repos *infrastructure.Repositories) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "not allwowed", http.StatusMethodNotAllowed)
			return
		}

		var userRaw struct {
			Username string `json:"username"`
		}
		err := json.NewDecoder(r.Body).Decode(&userRaw)
		if err != nil {
			// TODO: what error should we have here??
			handleResponse(w, http.StatusBadRequest, nil, err)
			return
		}

		ok, err := user.CreateUser(userRaw.Username, repos.User, repos.Factory)
		if err != nil {
			// TODO: what error should we have here??
			handleResponse(w, http.StatusBadRequest, nil, err)
			return
		}

		handleResponse(w, http.StatusOK, fmt.Sprintf("%t", ok), nil)
	}
}

func reqGetDashboard(repos *infrastructure.Repositories) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "not allwowed", http.StatusMethodNotAllowed)
			return
		}

		username := r.URL.Query().Get("username")
		user, err := user.GetUserByUsername(username, repos.User)
		if err != nil {
			// TODO: what error should we have here??
			handleResponse(w, http.StatusBadRequest, nil, err)
			return
		}

		handleResponse(w, http.StatusOK, user, nil)
	}
}
