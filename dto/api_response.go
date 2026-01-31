package dto

import (
	"encoding/json"
	"net/http"
)

type APIResponse struct {
	Message string      `json:"message"`
	Data    any `json:"data,omitempty"`
}

// Untuk response sukses
func WriteJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

// Untuk response error
func WriteError(w http.ResponseWriter, status int, message string) {
	WriteJSON(w, status, APIResponse{
		Message: message,
	})
}