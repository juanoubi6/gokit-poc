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
	Errors  *[]string
}

func (v ValidationError) Error() string {
	return v.Message
}

func EncodeJSONError(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("Calling EncodeJSONError without an error.")
	}

	var httpStatusCode int
	var errors []string
	var message string

	switch err.(type) {
	case AuthorizationError:
		httpStatusCode = http.StatusUnauthorized
		errors = []string{err.Error()}
		message = err.(AuthorizationError).Message
	case BusinessError:
		httpStatusCode = http.StatusBadRequest
		errors = []string{err.Error()}
		message = err.(BusinessError).Message
	case ValidationError:
		httpStatusCode = http.StatusBadRequest
		if errSlice := err.(ValidationError).Errors; errSlice != nil {
			message = "Validation errors found"
			errors = *errSlice
		} else {
			errors = []string{err.Error()}
			message = err.(ValidationError).Message
		}
	default:
		httpStatusCode = http.StatusInternalServerError
		errors = []string{err.Error()}
	}

	w.Header().Set(ContentType, ApplicationJSON)
	w.WriteHeader(httpStatusCode)
	_ = json.NewEncoder(w).Encode(GenericResponse{
		Success:   false,
		Message:   message,
		Errors:    errors,
		HttpCode:  httpStatusCode,
		Timestamp: time.Now().Format(time.RFC3339),
	})
}
