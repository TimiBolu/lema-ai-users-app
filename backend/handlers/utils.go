package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

type SuccessResponse[T any] struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    T      `json:"data,omitempty"`
}

type ErrorResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func NewSuccessResponse[T any](message string, data T) SuccessResponse[T] {
	return SuccessResponse[T]{
		Success: true,
		Message: message,
		Data:    data,
	}
}

func NewErrorResponse(message string) ErrorResponse {
	return ErrorResponse{
		Success: false,
		Message: message,
	}
}

func (r SuccessResponse[T]) JSON() ([]byte, error) {
	return json.Marshal(r)
}

func (r ErrorResponse) JSON() ([]byte, error) {
	return json.Marshal(r)
}

func SendSuccessResponse[T any](w http.ResponseWriter, statusCode int, message string, data T) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := NewSuccessResponse(message, data)
	json.NewEncoder(w).Encode(response)
}

func SendErrorResponse(w http.ResponseWriter, statusCode int, message string, logger *logrus.Logger,
	err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := NewErrorResponse(message)
	json.NewEncoder(w).Encode(response)

	logger.WithError(err).Error(message)
}
