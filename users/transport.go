package users

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"gokit-poc/models"
)

type Endpoints struct {
	CreateUser endpoint.Endpoint
}

func MakeEndpoints(svc UserService) Endpoints {
	return Endpoints{
		CreateUser: MakeCreateUserEndpoint(svc),
	}
}

func MakeCreateUserEndpoint(svc UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateUserRequest)
		result, err := svc.CreateUser(ctx, req)
		if err != nil {
			return nil, err
		}

		return CreateUserResponse{User: result}, nil
	}
}

type CreateUserRequest struct {
	Name     string `json:"name" validate:"required"`
	LastName string `json:"lastName" validate:"required"`
	Age      int    `json:"age" validate:"required"`
}

type CreateUserResponse struct {
	User *models.User `json:"user"`
}
