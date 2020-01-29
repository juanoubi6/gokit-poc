package users

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"gokit-poc/models"
)

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
	Age      int    `json:"age"`
}

type CreateUserResponse struct {
	User *models.User
}
