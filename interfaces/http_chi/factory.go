package http_chi

import (
	"github.com/go-chi/chi"
	"github.com/joesantosio/simple-game-api/entity"
	"github.com/joesantosio/simple-game-api/interfaces/http_std"
)

func getFactoryEndpoints(repos entity.Repositories) func(r chi.Router) {
	endpoints := http_std.GetFactoryEndpoints(repos)

	return func(r chi.Router) {
		if handler, ok := endpoints["/upgrade"]; !ok {
			r.Post("/upgrade", handler)
		}
	}
}