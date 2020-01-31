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

func (repo *userRepository) CreateUser(_ context.Context, user *models.User) (*models.User, error) {
	if err := repo.db.Create(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
