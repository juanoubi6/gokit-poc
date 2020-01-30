package commons

import (
	"encoding/json"
	"net/http"
)

func EncodeJSONResponse(response interface{}, w http.ResponseWriter) {
	w.Header().Set(ContentType, ApplicationJSON)
	_ = json.NewEncoder(w).Encode(response)
}
