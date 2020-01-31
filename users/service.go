package users

import (
	"context"
	"errors"
	"github.com/jinzhu/gorm"
	"gokit-poc/models"
)

// UserService provides operations on users.
type UserService interface {
	CreateUser(context.Context, CreateUserRequest) (*models.User, error)
}

type Service struct {
	repository UserRepository
}

func UserServiceFactory(db *gorm.DB) UserService {
	repo := NewUserRepository(db)
	svc := &Service{repo}
	userService := InstrumentingMiddlewareDecorator(svc)
	userService = LoggingMiddlewareDecorator(userService)

	return svc
}

func (s *Service) CreateUser(ctx context.Context, req CreateUserRequest) (*models.User, error) {
	if req.Age == 0 {
		return nil, errors.New("age can't be 0")
	}

	user := &models.User{
		Name:     req.Name,
		LastName: req.LastName,
		Age:      req.Age,
	}

	return s.repository.CreateUser(ctx, user)
}
