package http_std

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/joesantosio/simple-game-api/entity"
)

type errorCoded interface {
	Error() string
	StatusCode() int
}

func handleResponseRaw(w http.ResponseWriter, code int, data interface{}) {
	// prepare the response
	resData := struct {
		Ok   bool        `json:"ok"`
		Code int         `json:"code"`
		Data interface{} `json:"data,omitempty"`
		Err  string      `json:"err,omitempty"`
	}{true, code, data, ""}
	if code > http.StatusOK + 99 {
		resData.Ok = false
		resData.Err = data.(string)
		resData.Data = ""
	}

	r, marshalErr := json.Marshal(resData)
	if marshalErr != nil {
		handleErrResponse(w, marshalErr)
		return
	}

	// send the response
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	w.Write(r)
}

func handleResponse(w http.ResponseWriter, data interface{}) {
	handleResponseRaw(w, http.StatusOK, data)
}

func handleErrResponse(w http.ResponseWriter, err error) {
	code := http.StatusInternalServerError
	if coder, ok := err.(errorCoded); ok {
		code = coder.StatusCode()
	}

	msg := err.Error()
	if code >= http.StatusInternalServerError && code <= http.StatusInternalServerError + 99 {
		msg = http.StatusText(code)
	}

	handleResponseRaw(w, code, msg)
}

type endpointHandler struct {
	fn func(w http.ResponseWriter, r *http.Request)
}
func (h *endpointHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.fn(w, r)
}
func newEndpointHandler(fn func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return &endpointHandler{fn}
}

func setEndpoints(mux *http.ServeMux, repos entity.Repositories) {
	for key, val := range getUserEndpoints(repos) {
		mux.Handle(key, val)
	}

	for key, val := range getFactoryEndpoints(repos) {
		mux.Handle(key, val)
	}
}

func InitServer(address string, repos entity.Repositories) {
	mux := http.NewServeMux()
	setEndpoints(mux, repos)

	srv := &http.Server{
		Addr:         address,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      mux,
	}

	fmt.Printf("listening at %s \n", address)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalf("server failed to start: %v", err)
	}
}
