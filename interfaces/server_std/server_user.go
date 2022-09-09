package serverstd

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/joesantosio/example-go-project/acl"
	"github.com/joesantosio/example-go-project/domain/user"
	"github.com/joesantosio/example-go-project/entity"
	"github.com/joesantosio/example-go-project/httperr"
	"github.com/joesantosio/example-go-project/interfaces/server"
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
			server.HandleErrResponse(w, err)
			return
		}

		ok, err := user.Create(userRaw.Username, repos.GetUser(), repos.GetFactory())
		if err != nil {
			server.HandleErrResponse(w, err)
			return
		}

		server.HandleResponse(w, ok)
	}
}

func reqGetDashboard(repos entity.Repositories) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, ok := r.Context().Value(acl.AclUserID).(string)
		if !ok || len(userID) == 0 {
			err := httperr.NewError(http.StatusForbidden, errors.New("userRoles not set right"))
			server.HandleErrResponse(w, err)
			return
		}

		user, err := user.GetByUserID(userID, repos.GetUser())
		if err != nil {
			server.HandleErrResponse(w, err)
			return
		}

		server.HandleResponse(w, user)
	}
}

func GetUserEndpoints(repos entity.Repositories) map[string]func(w http.ResponseWriter, r *http.Request) {
	endpoints := make(map[string]func(w http.ResponseWriter, r *http.Request))

	endpoints["/user"] = reqCreateUser(repos)
	// TODO: should enforce user roles
	endpoints["/dashboard"] = reqGetDashboard(repos)

	return endpoints
}
