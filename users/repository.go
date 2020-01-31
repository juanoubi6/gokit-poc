package users

import (
	"context"
	"github.com/jinzhu/gorm"
	"gokit-poc/models"
)

type UserRepository interface {
	CreateUser(context.Context, *models.User) (*models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

//TODO: Add error handling
func (repo *userRepository) CreateUser(_ context.Context, user *models.User) (*models.User, error) {
	repo.db.Create(&user)

	return user, nil
}
