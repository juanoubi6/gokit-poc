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

type ValidationError struct {
	Message string
}

func (v ValidationError) Error() string {
	return v.Message
}

func EncodeJSONError(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("Calling EncodeJSONError without an error.")
	}

	var httpStatusCode int

	switch err.(type) {
	case AuthorizationError:
		httpStatusCode = http.StatusUnauthorized
	case BusinessError, ValidationError:
		httpStatusCode = http.StatusBadRequest
	default:
		httpStatusCode = http.StatusInternalServerError
	}

	w.Header().Set(ContentType, ApplicationJSON)
	w.WriteHeader(httpStatusCode)
	_ = json.NewEncoder(w).Encode(GenericResponse{
		Success:   false,
		Errors:    []string{err.Error()},
		HttpCode:  httpStatusCode,
		Timestamp: time.Now().Format(time.RFC3339),
	})
}
