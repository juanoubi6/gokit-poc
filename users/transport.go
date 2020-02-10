package users

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"gokit-poc/models"
)

type Endpoints struct {
	CreateUser endpoint.Endpoint
	GetUsers   endpoint.Endpoint
}

func MakeEndpoints(svc UserService) Endpoints {
	return Endpoints{
		CreateUser: MakeCreateUserEndpoint(svc),
		GetUsers:   MakeGetUsersEndpoint(svc),
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

func MakeGetUsersEndpoint(svc UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetUsersRequest)
		result, err := svc.GetUsers(ctx, req)
		if err != nil {
			return nil, err
		}

		return GetUsersResponse{Users: result}, nil
	}
}

type CreateUserRequest struct {
	Name     string `json:"name" validate:"required"`
	LastName string `json:"lastName" validate:"required"`
	Age      int    `json:"age" validate:"gt=0"`
}

type CreateUserResponse struct {
	User *models.User `json:"user"`
}

type GetUsersRequest struct {
	Name     string `schema:"name"`
	LastName string `schema:"lastName"`
	Age      int    `schema:"age"`
}

type GetUsersResponse struct {
	Users []*models.User `json:"users"`
}
