package httphandler

import (
	"encoding/json"
	"net/http"
)

type responseAPI struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

func respSuccessJSON(w http.ResponseWriter, r *http.Request, data interface{}) {
	js, err := json.Marshal(responseAPI{
		Status: http.StatusOK,
		Data:   data,
	})
	if err != nil {
		respErrorText(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
	return
}

func respErrorJSON(w http.ResponseWriter, r *http.Request, status int, errStr string) {
	js, err := json.Marshal(responseAPI{
		Status: status,
		Data: struct {
			Message string `json:"message"`
		}{errStr},
	})
	if err != nil {
		respErrorText(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)
	return
}

func respErrorText(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text")
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("500 - Something bad happened!"))
}

const (
	errBadRequest = "Bad request"
)
