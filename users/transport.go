package users

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"gokit-poc/models"
)

func GetEndpoints(svc UserService) []endpoint.Endpoint {
	return []endpoint.Endpoint{
		MakeCreateUserEndpoint(svc),
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
	Name     string `json:"name"`
	LastName string `json:"lastName"`
	Age      string `json:"age"`
}

type CreateUserResponse struct {
	User models.User
}
