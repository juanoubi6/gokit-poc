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

type ErrorResponse struct {
	Errors    []string `json:"errors"`
	HttpCode  int      `json:"httpCode"`
	Timestamp string   `json:"timestamp"`
}

func EncodeJSONError(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("Calling EncodeJSONError without an error.")
	}

	w.Header().Set(ContentType, ApplicationJSON)
	w.WriteHeader(http.StatusBadRequest)
	_ = json.NewEncoder(w).Encode(ErrorResponse{
		Errors:    []string{err.Error()},
		HttpCode:  http.StatusBadRequest,
		Timestamp: time.Now().Format(time.RFC3339),
	})
}
