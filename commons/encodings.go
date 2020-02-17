package commons

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"io"
	"net/http"
	"strings"
	"time"
)

var validate = validator.New()

type GenericResponse struct {
	Success   bool        `json:"success"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
	Errors    []string    `json:"errors"`
	HttpCode  int         `json:"httpCode"`
	Timestamp string      `json:"timestamp"`
}

func EncodeJSONResponse(message string, httpCode int, response interface{}, w http.ResponseWriter) {
	w.Header().Set(ContentType, ApplicationJSON)
	w.WriteHeader(httpCode)
	_ = json.NewEncoder(w).Encode(GenericResponse{
		Success:   true,
		Message:   message,
		Data:      response,
		HttpCode:  httpCode,
		Timestamp: time.Now().Format(time.RFC3339),
	})
}

func EncodeAndValidate(r io.Reader, container interface{}) error {
	if err := json.NewDecoder(r).Decode(&container); err != nil {
		return ValidationError{Message: err.Error(), Errors: nil}
	}
	if err := validate.Struct(container); err != nil {
		if fieldErrors, ok := err.(validator.ValidationErrors); ok {
			var errorTags []string
			for _, fieldError := range fieldErrors {
				errorTags = append(errorTags, strings.Join([]string{fieldError.Field(), fieldError.Tag(), fieldError.Param()}, " "))
			}
			return ValidationError{Message: err.Error(), Errors: &errorTags}
		} else {
			return ValidationError{Message: err.Error(), Errors: nil}
		}
	}

	return nil
}
