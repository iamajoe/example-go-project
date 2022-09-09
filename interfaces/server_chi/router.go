package serverchi

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/joesantosio/example-go-project/acl"
	"github.com/joesantosio/example-go-project/entity"
	"github.com/joesantosio/example-go-project/httperr"
	"github.com/joesantosio/example-go-project/interfaces/server"
)

func accessControl(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		next.ServeHTTP(w, r)
	})
}

func setEndpoints(r *chi.Mux, repos entity.Repositories) {
	r.Route("/users", getUserEndpoints(repos))
	r.Route("/factories", getFactoryEndpoints(repos))
}

func GetParamID(r *http.Request, param string) (int, error) {
	idStr := chi.URLParam(r, param)
	if idStr == "" {
		return -1, httperr.NewError(
			http.StatusBadRequest,
			errors.New(fmt.Sprintf("%s required", param)),
		)
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return -1, httperr.NewError(http.StatusBadRequest, err)
	}

	return id, nil
}

func GetRouter(authSecret string, repos entity.Repositories) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.CleanPath)
	r.Use(middleware.NoCache)
	r.Use(middleware.Timeout(60 * time.Second))
	r.Use(accessControl)
	r.Use(acl.SetAclMiddleware(authSecret, repos, func(w http.ResponseWriter, r *http.Request, err error) {
		server.HandleErrResponse(w, err)
	}))

	r.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		server.HandleErrResponse(w, httperr.NewError(http.StatusMethodNotAllowed, errors.New("not allowed")))
	})
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		server.HandleErrResponse(w, httperr.NewError(http.StatusNotFound, errors.New("not allowed")))
	})

	setEndpoints(r, repos)

	return r
}
