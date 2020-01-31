package authentications

import (
	"context"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"gokit-poc/commons"
	"net/http"
)

func NewHTTPHandler(router *mux.Router, endpoints Endpoints) {
	println("Adding routes")
	// Options added here:
	// ServerErrorEncoder: handles decoding errors
	opts := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(commons.EncodeJSONError),
	}

	router.Methods(http.MethodPost).Path("/signup").Handler(httptransport.NewServer(
		endpoints.SignUp,
		decodeSignUpRequest,
		encodeSignUpResponse,
		opts...,
	))

	router.Methods(http.MethodPost).Path("/login").Handler(httptransport.NewServer(
		endpoints.Login,
		decodeLoginRequest,
		encodeLoginResponse,
		opts...,
	))

}

func decodeSignUpRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request SignUpRequest
	if err := commons.EncodeAndValidate(r.Body, &request); err != nil {
		return nil, err
	}

	return request, nil
}

func encodeSignUpResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	switch res := response.(type) {
	case commons.BusinessError:
		commons.EncodeJSONError(ctx, res, w)
	case SignUpResponse:
		commons.EncodeJSONResponse(response, w)
	}

	return nil
}

func decodeLoginRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request LoginRequest
	if err := commons.EncodeAndValidate(r.Body, &request); err != nil {
		return nil, err
	}

	return request, nil
}

func encodeLoginResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	switch res := response.(type) {
	case commons.BusinessError:
		commons.EncodeJSONError(ctx, res, w)
	case LoginResponse:
		commons.EncodeJSONResponse(response, w)
	}

	return nil
}
