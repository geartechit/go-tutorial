package httputil

import (
	"encoding/json"
	"go-tutorial/internal/dto"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(payload)
}

func WriteSuccess[T any](w http.ResponseWriter, statusCode int, data T) {
	WriteJSON(w, statusCode, dto.BaseResponse[T]{
		Success: true,
		Data:    data,
	})
}

func WriteError(w http.ResponseWriter, statusCode int, message string, details []string) {
	WriteJSON(w, statusCode, dto.ErrorResponse{
		Success: false,
		Message: message,
		Details: details,
	})
}
