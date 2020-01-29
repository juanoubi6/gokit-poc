package users

import (
	"context"
	"gokit-poc/models"
)

// UserService provides operations on users.
type UserService interface {
	CreateUser(context.Context, CreateUserRequest) (models.User, error)
}

type Service struct{}

func CreateUserService() UserService {
	return &Service{}
}

func (s *Service) CreateUser(ctx context.Context, req CreateUserRequest) (models.User, error) {
	return models.User{
		Name:     "John",
		LastName: "Doe",
		Age:      23,
	}, nil
}
