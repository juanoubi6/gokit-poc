package authentications

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"gokit-poc/models"
)

type Endpoints struct {
	SignUp endpoint.Endpoint
	Login  endpoint.Endpoint
}

func MakeEndpoints(svc AuthenticationService) Endpoints {
	return Endpoints{
		SignUp: MakeSignUpEndpoint(svc),
		Login:  MakeLoginEndpoint(svc),
	}
}

func MakeSignUpEndpoint(svc AuthenticationService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(SignUpRequest)
		result, err := svc.SignUp(ctx, req)
		if err != nil {
			return nil, err
		}

		return SignUpResponse{Account: result}, nil
	}
}

type SignUpRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=100"`
}

type SignUpResponse struct {
	Account *models.Account `json:"account"`
}

func MakeLoginEndpoint(svc AuthenticationService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(LoginRequest)
		result, err := svc.Login(ctx, req)
		if err != nil {
			return nil, err
		}

		return LoginResponse{Token: result}, nil
	}
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=100"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
