package server

import (
	"encoding/json"
	"net/http"
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
	if code > http.StatusOK+99 {
		resData.Ok = false
		resData.Err = data.(string)
		resData.Data = nil
	}

	r, marshalErr := json.Marshal(resData)
	if marshalErr != nil {
		HandleErrResponse(w, marshalErr)
		return
	}

	// send the response
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	w.Write(r)
}

func HandleResponse(w http.ResponseWriter, data interface{}) {
	handleResponseRaw(w, http.StatusOK, data)
}

func HandleErrResponse(w http.ResponseWriter, err error) {
	code := http.StatusInternalServerError
	if coder, ok := err.(errorCoded); ok {
		code = coder.StatusCode()
	}

	// msg := err.Error()
	msg := http.StatusText(code)
	// TODO: if not production, we can send the message instead
	// if code >= http.StatusInternalServerError && code <= http.StatusInternalServerError + 99 {
	// }

	handleResponseRaw(w, code, msg)
}
