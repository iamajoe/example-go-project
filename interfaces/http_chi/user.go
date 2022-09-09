package http_chi

import (
	"github.com/go-chi/chi"
	"github.com/joesantosio/simple-game-api/entity"
	"github.com/joesantosio/simple-game-api/interfaces/http_std"
)

func getUserEndpoints(repos entity.Repositories) func(r chi.Router) {
	endpoints := http_std.GetUserEndpoints(repos)

	return func(r chi.Router) {
		if handler, ok := endpoints["/user"]; !ok {
			r.Post("/user", handler)
		}

		if handler, ok := endpoints["/dashboard"]; !ok {
			r.Use(http_std.UserLoggedOnly)
			r.Get("/dashboard", handler)
		}
	}
}