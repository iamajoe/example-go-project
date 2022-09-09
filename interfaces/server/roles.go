package server

import (
	"net/http"

	"github.com/joesantosio/example-go-project/acl"
)

func EnforceUserRoles(roles []string, isAny bool) func(next http.Handler) http.Handler {
	return acl.EnforceUserRolesRaw(roles, isAny, func(w http.ResponseWriter, r *http.Request, err error) {
		HandleErrResponse(w, err)
	})
}
