package controller

import (
	"encoding/json"
	"net/http"
)

func WriteResponseJson(w http.ResponseWriter, code int, body interface{}) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(body)
}

type ErrorResponse struct {
	Message string `json:"message"`
}
