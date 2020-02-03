package commons

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
)

type BusinessError struct {
	Message string
}

func (b BusinessError) Error() string {
	return b.Message
}

type AuthorizationError struct {
	Message string
}

func (a AuthorizationError) Error() string {
	return a.Message
}

type ErrorResponse struct {
	Errors    []string `json:"errors"`
	HttpCode  int      `json:"httpCode"`
	Timestamp string   `json:"timestamp"`
}

func EncodeJSONError(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("Calling EncodeJSONError without an error.")
	}

	var status int

	switch err.(type) {
	case AuthorizationError:
		status = http.StatusUnauthorized
	case BusinessError:
		status = http.StatusBadRequest
	default:
		status = http.StatusInternalServerError
	}

	w.Header().Set(ContentType, ApplicationJSON)
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(ErrorResponse{
		Errors:    []string{err.Error()},
		HttpCode:  status,
		Timestamp: time.Now().Format(time.RFC3339),
	})
}
