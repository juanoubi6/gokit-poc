package users

import (
	"context"
	"encoding/json"
	httptransport "github.com/go-kit/kit/transport/http"
	"net/http"
)

func CreateUserHandler(svc UserService) *httptransport.Server {
	return httptransport.NewServer(
		MakeCreateUserEndpoint(svc),
		decodeCreateUserRequest,
		encodeCreateUserResponse,
	)
}

func decodeCreateUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeCreateUserResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
