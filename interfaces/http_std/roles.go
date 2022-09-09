package http_std

import (
	"context"
	"errors"
	"net/http"

	"github.com/go-chi/jwtauth"
	"github.com/joesantosio/simple-game-api/entity"
	"github.com/joesantosio/simple-game-api/httperr"
)

var (
	ACL_PERMISSION = "acl.permission"
)

type UserRoles struct {
	repos entity.Repositories
	userId string
	isUserRegistered bool
}

func (r *UserRoles) IsUserRegistered() bool {
	// TODO: we can check here user repo

	if r.userId != "" { 
		return true
	}

	return r.isUserRegistered
}
func NewUserRoles(repos entity.Repositories, userId string) *UserRoles {
	return &UserRoles{repos, userId, false}
}

func UserLoggedOnly(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    ctx := r.Context()
    perm, ok := ctx.Value(ACL_PERMISSION).(UserRoles)

		if !ok || !perm.IsUserRegistered() {
			err := httperr.NewError(http.StatusForbidden, errors.New("not registered"))
			HandleErrResponse(w, err)
      return
    }

    next.ServeHTTP(w, r)
  })
}

func SetAclMiddleware(repos entity.Repositories) func (next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			_, claims, _ := jwtauth.FromContext(ctx)
			
			userIdToBe := ""
			if userId, ok := claims["user_id"].(string); ok {
				userIdToBe = userId
			}

			ctx = context.WithValue(ctx, ACL_PERMISSION, NewUserRoles(repos, userIdToBe))

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}