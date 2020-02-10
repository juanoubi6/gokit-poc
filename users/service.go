package users

import (
	"context"
	"github.com/jinzhu/gorm"
	"gokit-poc/commons"
	"gokit-poc/models"
)

// UserService provides operations on users.
type UserService interface {
	CreateUser(context.Context, CreateUserRequest) (*models.User, error)
	GetUsers(context.Context, GetUsersRequest) ([]*models.User, error)
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
	if req.Age == 100 {
		return nil, commons.BusinessError{"Age can't be 100"}
	}

	user := &models.User{
		Name:     req.Name,
		LastName: req.LastName,
		Age:      req.Age,
	}

	return s.repository.CreateUser(ctx, user)
}

func (s *Service) GetUsers(ctx context.Context, req GetUsersRequest) ([]*models.User, error) {
	return s.repository.GetUsers(ctx, req)
}
