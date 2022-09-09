package acl

import (
	"context"
	"net/http"
	"strings"

	"github.com/go-chi/jwtauth"
	"github.com/joesantosio/example-go-project/entity"
	"github.com/joesantosio/example-go-project/httperr"
	"github.com/lestrrat-go/jwx/jwt"
)

func GetToken(userID string, userTokenRepo entity.RepositoryUserToken) (string, error) {
	// TODO: need to setup validity date / cron to validate

	_, tokenString, err := tokenAuth.Encode(
		map[string]interface{}{"userid": userID},
	)

	// we save the token on the repo so we can invalid on our side
	_, err = userTokenRepo.Create(userID, tokenString)
	if err != nil {
		return "", err
	}

	return tokenString, err
}

func GetTokenUserID(tokenStr string, userTokenRepo entity.RepositoryUserToken) (userID string, err error) {
	userID = ""

	if strings.TrimSpace(tokenStr) == "" {
		return userID, nil
	}

	// check if the jwt is valid by checking the database
	// if it doesnt exist, it means that we invalidated it
	ok, err := userTokenRepo.IsTokenValid(tokenStr)
	if !ok || err != nil {
		return userID, err
	}

	// get the user id from the token and register on context
	token, err := tokenAuth.Decode(tokenStr)
	if err != nil {
		return userID, httperr.NewError(http.StatusBadRequest, err)
	}

	if token == nil {
		return userID, nil
	}

	// validate the token
	if err := jwt.Validate(token); err != nil {
		return userID, httperr.NewError(http.StatusBadRequest, err)
	}

	// get the user id from the token
	val, ok := token.Get("userid")
	if !ok {
		return "", nil
	}

	userID, ok = val.(string)
	if !ok {
		return "", httperr.NewError(http.StatusBadRequest, err)
	}

	return userID, nil
}

func SetContextRoles(ctx context.Context, userID string, repos entity.Repositories) (context.Context, error) {
	ctx = context.WithValue(ctx, AclUserID, userID)
	ctx = context.WithValue(ctx, AclPermission, &userRoles{repos, userID, []string{}})

	return ctx, nil
}

func SetAclMiddleware(authSecret string, repos entity.Repositories, errHandler func(http.ResponseWriter, *http.Request, error)) func(next http.Handler) http.Handler {
	Init(authSecret)

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			tokenStr := jwtauth.TokenFromCookie(r)
			if tokenStr == "" {
				tokenStr = jwtauth.TokenFromHeader(r)
			}

			userID, err := GetTokenUserID(tokenStr, repos.GetUserToken())
			if err != nil {
				errHandler(w, r, err)
				return
			}

			ctx, err = SetContextRoles(ctx, userID, repos)
			if err != nil {
				errHandler(w, r, err)
				return
			}

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func Init(authSecret string) {
	tokenAuth = jwtauth.New("HS256", []byte(authSecret), nil)
}
