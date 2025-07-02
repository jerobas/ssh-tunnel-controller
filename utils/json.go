package utils

import (
	"encoding/json"
	"net/http"
)

// JSON writes a response with the given status and payload
func JSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

// Error writes a JSON error response
func Error(w http.ResponseWriter, status int, message string) {
	JSON(w, status, map[string]string{"error": message})
}

// Success writes a JSON success message
func Success(w http.ResponseWriter, status int, message string) {
	JSON(w, status, map[string]string{"message": message})
}
