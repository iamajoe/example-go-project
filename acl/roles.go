package acl

import (
	"errors"
	"net/http"

	"github.com/go-chi/jwtauth"
	"github.com/joesantosio/example-go-project/entity"
	"github.com/joesantosio/example-go-project/httperr"
)

var (
	AclUserID     = "acl.userid"
	AclPermission = "acl.permission"
	AclRoleAuth   = "acl.role.auth"
)
var tokenAuth *jwtauth.JWTAuth

type userRoles struct {
	repos  entity.Repositories
	userID string
	roles  []string
}

func (r *userRoles) getRoleChecker(role string, req *http.Request) func() (bool, error) {
	switch role {
	case AclRoleAuth:
		return func() (bool, error) {
			users, err := r.repos.GetUser().GetByIDs([]string{r.userID})
			return err == nil && len(users) > 0, nil
		}
	}

	return func() (bool, error) {
		return false, nil
	}
}

func (r *userRoles) HasUserRole(role string, req *http.Request) (bool, error) {
	// have we already checked if the role is in?
	for _, val := range r.roles {
		if val == role {
			return true, nil
		}
	}

	// fetch the role
	checkFn := r.getRoleChecker(role, req)
	in, err := checkFn()
	if err != nil {
		return false, err
	}

	// cache so we don't need to request again
	if in {
		r.roles = append(r.roles, role)
	}

	return in, nil
}

func EnforceUserRolesRaw(roles []string, isAny bool, errHandler func(http.ResponseWriter, *http.Request, error)) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			perm, ok := r.Context().Value(AclPermission).(*userRoles)
			if !ok {
				err := httperr.NewError(http.StatusForbidden, errors.New("userRoles not set right"))
				errHandler(w, r, err)
				return
			}

			hasRolesIn := []string{}
			for _, role := range roles {
				ok, err := perm.HasUserRole(role, r)
				if ok && err == nil {
					hasRolesIn = append(hasRolesIn, role)

					// being any, we don't need to keep checking
					if isAny {
						break
					}
				}
			}

			if (isAny && len(hasRolesIn) == 0) ||
				(!isAny && len(hasRolesIn) != len(roles)) {
				errHandler(w, r, httperr.NewError(http.StatusForbidden, errors.New("unauthorized")))
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
