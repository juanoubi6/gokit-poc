package users

import (
	"context"
	"encoding/json"
	httptransport "github.com/go-kit/kit/transport/http"
	"gokit-poc/commons"
	"net/http"
)

// Options added here:
// ServerErrorEncoder: handles decoding errors
var opts = []httptransport.ServerOption{
	httptransport.ServerErrorEncoder(commons.EncodeJSONError),
}

func CreateUserHandler(svc UserService) *httptransport.Server {
	return httptransport.NewServer(
		MakeCreateUserEndpoint(svc),
		decodeCreateUserRequest,
		encodeCreateUserResponse,
		opts...,
	)
}

func decodeCreateUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}

	return request, nil
}

func encodeCreateUserResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	switch res := response.(type) {
	case commons.BusinessError:
		commons.EncodeJSONError(ctx, res, w)
	case CreateUserResponse:
		commons.EncodeJSONResponse(response, w)
	}

	return nil
}
