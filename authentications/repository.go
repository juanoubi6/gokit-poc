package authentications

import (
	"context"
	"errors"
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
	existingAccount, err := repo.GetAccountByEmail(context.Background(), account.Email)
	if err != nil {
		return nil, commons.BusinessError{err.Error()}
	}
	if existingAccount != nil {
		return nil, commons.BusinessError{"Email already in use"}
	}
	if err := repo.db.Create(&account).Error; err != nil {
		return nil, errors.New("unexpected error when creating account: " + err.Error())
	}

	return account, nil
}

func (repo *authenticationRepository) GetAccountByEmail(_ context.Context, email string) (*models.Account, error) {
	var account models.Account
	err := repo.db.Where("email = ?", email).First(&account).Error

	if gorm.IsRecordNotFoundError(err) {
		return nil, nil
	}

	if err != nil {
		return nil, commons.BusinessError{err.Error()}
	}

	return &account, nil
}
