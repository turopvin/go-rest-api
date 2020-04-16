package apiserver

import (
	"encoding/json"
	"net/http"
)

func SendError(w http.ResponseWriter, r *http.Request, code int, err error) {
	if err == nil {
		SendRespond(w, r, code, nil)
		return
	}
	SendRespond(w, r, code, map[string]string{"error": err.Error()})
}

func SendRespond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
