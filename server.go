package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func handleResponse(w http.ResponseWriter, code int, data interface{}, err error) {
	if err != nil {
		codeToBe := code
		if code == http.StatusOK {
			codeToBe = http.StatusBadRequest
		}

		handleResponse(w, codeToBe, err.Error(), nil)
		return
	}

	// prepare the response
	resData := struct {
		Ok   bool        `json:"ok"`
		Code int         `json:"code"`
		Data interface{} `json:"data,omitempty"`
		Err  string      `json:"err,omitempty"`
	}{true, code, data, ""}
	if code != http.StatusOK {
		resData.Ok = false
		resData.Err = data.(string)
		resData.Data = ""
	}

	r, marshalErr := json.Marshal(resData)
	if marshalErr != nil {
		handleResponse(w, http.StatusBadRequest, nil, marshalErr)
		return
	}

	// send the response
	w.WriteHeader(code)
	w.Write(r)
}

func initServer(repos Repositories) {
	sm := http.NewServeMux()

	sm.HandleFunc("/user", reqCreateUser(repos))
	sm.HandleFunc("/dashboard", reqGetDashboard(repos))
	sm.HandleFunc("/factory/upgrade", reqUpgradeFactory(repos))

	fmt.Println("listening at :4040")
	err := http.ListenAndServe(":4040", sm)
	if err != nil {
		log.Fatalf(err.Error())
	}
}
