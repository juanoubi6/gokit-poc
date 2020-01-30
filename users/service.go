package users

import (
	"context"
	"errors"
	"gokit-poc/models"
)

// UserService provides operations on users.
type UserService interface {
	CreateUser(context.Context, CreateUserRequest) (*models.User, error)
}

type Service struct{}

func UserServiceFactory() UserService {
	svc := &Service{}
	userService := InstrumentingMiddlewareDecorator(svc)
	userService = LoggingMiddlewareDecorator(userService)

	return userService
}

func (s *Service) CreateUser(ctx context.Context, req CreateUserRequest) (*models.User, error) {
	if req.Age == 0 {
		return nil, errors.New("age can't be 0")
	}

	return &models.User{
		Name:     "John",
		LastName: "Doe",
		Age:      23,
	}, nil
}
