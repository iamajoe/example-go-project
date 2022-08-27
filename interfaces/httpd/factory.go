package httpd

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/joesantosio/simple-game-api/domains/factory"
	"github.com/joesantosio/simple-game-api/domains/user"
	"github.com/joesantosio/simple-game-api/infrastructure"
)

func reqUpgradeFactory(repos *infrastructure.Repositories) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "not allwowed", http.StatusMethodNotAllowed)
			return
		}

		username := r.URL.Query().Get("username")
		_, err := user.GetUserByUsername(username, repos.User)
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

		ok, err := factory.UpgradeUserResource(username, bodyResource["kind"], repos.User, repos.Factory)
		if err != nil {
			// TODO: what error should we have here??
			handleResponse(w, http.StatusBadRequest, nil, err)
			return
		}

		handleResponse(w, http.StatusOK, fmt.Sprintf("%t", ok), nil)
	}
}
