package users

import (
	"context"
	"github.com/jinzhu/gorm"
	"gokit-poc/commons"
	"gokit-poc/models"
)

type UserRepository interface {
	CreateUser(context.Context, *models.User) (*models.User, error)
	GetUsers(context.Context, GetUsersRequest) ([]*models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (repo *userRepository) CreateUser(_ context.Context, user *models.User) (*models.User, error) {
	if err := repo.db.Create(&user).Error; err != nil {
		return nil, commons.BusinessError{err.Error()}
	}

	return user, nil
}

func (repo *userRepository) GetUsers(_ context.Context, req GetUsersRequest) ([]*models.User, error) {
	var users []*models.User
	db := repo.db

	if req.Name != "" {
		db = db.Where("name = ?", req.Name)
	}
	if req.LastName != "" {
		db = db.Where("last_name = ?", req.LastName)
	}
	if req.Age != 0 {
		db = db.Where("age = ?", req.Age)
	}

	if err := db.Find(&users).Error; err != nil {
		return nil, commons.BusinessError{err.Error()}
	}

	return users, nil
}
