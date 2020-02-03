package authentications

import (
	"context"
	"github.com/jinzhu/gorm"
	"gokit-poc/commons"
	"gokit-poc/models"
)

type AuthenticationRepository interface {
	CreateAccount(context.Context, *models.Account) (*models.Account, error)
	GetAccountByEmail(context.Context, string) (*models.Account, error)
}

type authenticationRepository struct {
	db *gorm.DB
}

func NewAuthenticationRepository(db *gorm.DB) AuthenticationRepository {
	return &authenticationRepository{db}
}

func (repo *authenticationRepository) CreateAccount(_ context.Context, account *models.Account) (*models.Account, error) {
	if err := repo.db.Create(&account).Error; err != nil {
		return nil, commons.BusinessError{err.Error()}
	}

	return account, nil
}

func (repo *authenticationRepository) GetAccountByEmail(_ context.Context, email string) (*models.Account, error) {
	var account models.Account
	if err := repo.db.Where("email = ?", email).First(&account).Error; err != nil {
		return nil, commons.BusinessError{err.Error()}
	}

	return &account, nil
}
