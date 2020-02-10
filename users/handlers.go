package users

import (
	"context"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"gokit-poc/commons"
	"gokit-poc/security"
	"net/http"
)

var decoder = schema.NewDecoder()

func AddHTTPHandlersToRouter(router *mux.Router, endpoints Endpoints) {
	println("Adding user service routes")
	// Options added here:
	// ServerErrorEncoder: handles decoding errors
	opts := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(commons.EncodeJSONError),
		httptransport.ServerBefore(security.AuthTokenToContext()),
	}

	subRouter := router.PathPrefix("/users").Subrouter()

	subRouter.Methods(http.MethodPost).Path("").Handler(httptransport.NewServer(
		security.AccountAuthorizationMiddleware()(endpoints.CreateUser),
		decodeCreateUserRequest,
		encodeCreateUserResponse,
		opts...,
	))

	subRouter.Methods(http.MethodGet).Path("").Handler(httptransport.NewServer(
		security.AccountAuthorizationMiddleware()(endpoints.GetUsers),
		decodeGetUsersRequest,
		encodeGetUsersResponse,
		opts...,
	))

}

func decodeCreateUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request CreateUserRequest
	if err := commons.EncodeAndValidate(r.Body, &request); err != nil {
		return nil, err
	}

	return request, nil
}

func encodeCreateUserResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	switch res := response.(type) {
	case commons.BusinessError:
		commons.EncodeJSONError(ctx, res, w)
	case CreateUserResponse:
		commons.EncodeJSONResponse("User created", http.StatusCreated, response, w)
	}

	return nil
}

func decodeGetUsersRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request GetUsersRequest
	if err := decoder.Decode(&request, r.URL.Query()); err != nil {
		return nil, err
	}

	return request, nil
}

func encodeGetUsersResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	switch res := response.(type) {
	case commons.BusinessError:
		commons.EncodeJSONError(ctx, res, w)
	case GetUsersResponse:
		commons.EncodeJSONResponse("Users returned", http.StatusOK, response, w)
	}

	return nil
}
