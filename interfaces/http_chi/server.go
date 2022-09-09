package http_chi

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
	"github.com/joesantosio/simple-game-api/entity"
	"github.com/joesantosio/simple-game-api/interfaces/http_std"
)

var tokenAuth *jwtauth.JWTAuth

func accessControl(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		next.ServeHTTP(w, r)
	})
}

func setEndpoints(mux *chi.Mux, repos entity.Repositories) {
	mux.Route("/users", getUserEndpoints(repos))
	mux.Route("/factories", getFactoryEndpoints(repos))
}

func InitServer(address string, authSecret string, repos entity.Repositories) {
	tokenAuth = jwtauth.New("HS256", []byte(authSecret), nil)
	// _, tokenString, _ := tokenAuth.Encode(map[string]interface{}{"user_id": 123})
	// fmt.Printf("DEBUG: a sample jwt is %s\n\n", tokenString)

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(accessControl)
	r.Use(jwtauth.Verifier(tokenAuth))
	r.Use(http_std.SetAclMiddleware(repos))

	// Set a timeout value on the request context (ctx), that will signal
  // through ctx.Done() that the request has timed out and further
  // processing should be stopped.
  r.Use(middleware.Timeout(60 * time.Second))

	setEndpoints(r, repos)

	fmt.Printf("listening at %s \n", address)
	err := http.ListenAndServe(address, r)
	if err != nil {
		log.Fatalf("server failed to start: %v", err)
	}
}
