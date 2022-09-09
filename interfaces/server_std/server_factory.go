package serverstd

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/joesantosio/example-go-project/acl"
	"github.com/joesantosio/example-go-project/domain/factory"
	"github.com/joesantosio/example-go-project/entity"
	"github.com/joesantosio/example-go-project/httperr"
	"github.com/joesantosio/example-go-project/interfaces/server"
)

func reqUpgradeFactory(repos entity.Repositories) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "not allwowed", http.StatusMethodNotAllowed)
			return
		}

		userID, ok := r.Context().Value(acl.AclUserID).(string)
		if !ok || len(userID) == 0 {
			err := httperr.NewError(http.StatusForbidden, errors.New("userRoles not set right"))
			server.HandleErrResponse(w, err)
			return
		}

		var bodyResource map[string]string
		if err := json.NewDecoder(r.Body).Decode(&bodyResource); err != nil {
			server.HandleErrResponse(w, err)
			return
		}

		ok, err := factory.UpgradeUserResource(userID, bodyResource["kind"], repos.GetUser(), repos.GetFactory())
		if err != nil {
			server.HandleErrResponse(w, err)
			return
		}

		server.HandleResponse(w, ok)
	}
}

func GetFactoryEndpoints(repos entity.Repositories) map[string]func(w http.ResponseWriter, r *http.Request) {
	endpoints := make(map[string]func(w http.ResponseWriter, r *http.Request))

	endpoints["/upgrade"] = reqUpgradeFactory(repos)

	return endpoints
}
