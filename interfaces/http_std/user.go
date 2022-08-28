package http_std

import (
	"encoding/json"
	"net/http"

	"github.com/joesantosio/simple-game-api/domain/user"
	"github.com/joesantosio/simple-game-api/entity"
)

func reqCreateUser(repos entity.Repositories) func(w http.ResponseWriter, r *http.Request) {
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
			handleErrResponse(w, err)
			return
		}

		ok, err := user.CreateUser(userRaw.Username, repos.GetUser(), repos.GetFactory())
		if err != nil {
			handleErrResponse(w, err)
			return
		}

		handleResponse(w, ok)
	}
}

func reqGetDashboard(repos entity.Repositories) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "not allwowed", http.StatusMethodNotAllowed)
			return
		}

		username := r.URL.Query().Get("username")
		user, err := user.GetUserByUsername(username, repos.GetUser())
		if err != nil {
			handleErrResponse(w, err)
			return
		}

		handleResponse(w, user)
	}
}

func getUserEndpoints(repos entity.Repositories) map[string]http.Handler {
	endpoints := make(map[string]http.Handler)

	endpoints["/users/user"] = newEndpointHandler(reqCreateUser(repos))
	endpoints["/users/dashboard"] = newEndpointHandler(reqGetDashboard(repos))

	return endpoints
}