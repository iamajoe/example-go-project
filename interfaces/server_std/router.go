package serverstd

import (
	"fmt"
	"net/http"

	"github.com/joesantosio/example-go-project/entity"
)

type endpointHandler struct {
	fn func(w http.ResponseWriter, r *http.Request)
}

func (h *endpointHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.fn(w, r)
}

func setEndpoints(r *http.ServeMux, repos entity.Repositories) {
	for key, val := range GetUserEndpoints(repos) {
		r.Handle(fmt.Sprintf("/users/%s", key), &endpointHandler{val})
	}

	for key, val := range GetFactoryEndpoints(repos) {
		r.Handle(fmt.Sprintf("/factories/%s", key), &endpointHandler{val})
	}
}

func GetRouter(authSecret string, repos entity.Repositories) *http.ServeMux {
	r := http.NewServeMux()
	setEndpoints(r, repos)

	return r
}
