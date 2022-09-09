package serverchi

import (
	"github.com/go-chi/chi"
	"github.com/joesantosio/example-go-project/acl"
	"github.com/joesantosio/example-go-project/entity"
	"github.com/joesantosio/example-go-project/interfaces/server"
	serverstd "github.com/joesantosio/example-go-project/interfaces/server_std"
)

func getUserEndpoints(repos entity.Repositories) func(r chi.Router) {
	endpoints := serverstd.GetUserEndpoints(repos)

	return func(r chi.Router) {
		if handler, ok := endpoints["/user"]; !ok {
			r.Post("/user", handler)
		}

		if handler, ok := endpoints["/dashboard"]; !ok {
			r.With(
				server.EnforceUserRoles([]string{acl.AclRoleAuth}, false),
			).Get("/dashboard", handler)
		}
	}
}
