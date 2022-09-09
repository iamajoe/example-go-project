package http_std

import (
	"encoding/json"
	"net/http"

	"github.com/joesantosio/simple-game-api/domain/factory"
	"github.com/joesantosio/simple-game-api/domain/user"
	"github.com/joesantosio/simple-game-api/entity"
)

func reqUpgradeFactory(repos entity.Repositories) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "not allwowed", http.StatusMethodNotAllowed)
			return
		}

		username := r.URL.Query().Get("username")
		_, err := user.GetUserByUsername(username, repos.GetUser())
		if err != nil {
			HandleErrResponse(w, err)
			return
		}

		var bodyResource map[string]string
		if err = json.NewDecoder(r.Body).Decode(&bodyResource); err != nil {
			HandleErrResponse(w, err)
			return
		}

		ok, err := factory.UpgradeUserResource(username, bodyResource["kind"], repos.GetUser(), repos.GetFactory())
		if err != nil {
			HandleErrResponse(w, err)
			return
		}

		HandleResponse(w, ok)
	}
}

func GetFactoryEndpoints(repos entity.Repositories) map[string]func(w http.ResponseWriter, r *http.Request) {
	endpoints := make(map[string]func(w http.ResponseWriter, r *http.Request))

	endpoints["/upgrade"] = reqUpgradeFactory(repos)

	return endpoints
}