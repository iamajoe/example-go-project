package serverchi

import (
	"github.com/go-chi/chi"
	"github.com/joesantosio/example-go-project/entity"
	serverstd "github.com/joesantosio/example-go-project/interfaces/server_std"
)

func getFactoryEndpoints(repos entity.Repositories) func(r chi.Router) {
	endpoints := serverstd.GetFactoryEndpoints(repos)

	return func(r chi.Router) {
		if handler, ok := endpoints["/upgrade"]; !ok {
			r.Post("/upgrade", handler)
		}
	}
}
