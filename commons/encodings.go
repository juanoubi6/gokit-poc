package commons

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"io"
	"net/http"
)

var validate = validator.New()

func EncodeJSONResponse(response interface{}, w http.ResponseWriter) {
	w.Header().Set(ContentType, ApplicationJSON)
	_ = json.NewEncoder(w).Encode(response)
}

func EncodeAndValidate(r io.Reader, container interface{}) error {
	if err := json.NewDecoder(r).Decode(&container); err != nil {
		return err
	}
	if err := validate.Struct(container); err != nil {
		return err
	}

	return nil
}
