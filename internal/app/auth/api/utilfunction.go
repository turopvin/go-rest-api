package api

import (
	"encoding/json"
	"net/http"
)

func sendError(w http.ResponseWriter, r *http.Request, code int, err error) {
	if err == nil {
		sendRespond(w, r, code, nil)
		return
	}
	sendRespond(w, r, code, map[string]string{"error": err.Error()})
}

func sendRespond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
