package helpers

import (
	"encoding/json"
	"net/http"
)

type APIError struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

func WriteJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}

func WriteJSONError(w http.ResponseWriter, status int, errName, message string) {
	WriteJSON(w, status, APIError{
		Error:   errName,
		Message: message,
	})
}