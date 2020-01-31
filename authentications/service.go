package authentications

import (
	"context"
	"errors"
	"github.com/jinzhu/gorm"
	"gokit-poc/models"
	"golang.org/x/crypto/bcrypt"
)

// Authentication provides login and sign up operations.
type AuthenticationService interface {
	SignUp(context.Context, SignUpRequest) (*models.Account, error)
	Login(context.Context, LoginRequest) (string, error)
}

type Service struct {
	repository AuthenticationRepository
}

func AuthenticationServiceFactory(db *gorm.DB) AuthenticationService {
	repo := NewAuthenticationRepository(db)
	svc := &Service{repo}
	authenticationService := InstrumentingMiddlewareDecorator(svc)
	authenticationService = LoggingMiddlewareDecorator(authenticationService)

	return svc
}

func (s *Service) SignUp(ctx context.Context, req SignUpRequest) (*models.Account, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.MinCost)
	if err != nil {
		return nil, errors.New("unexpected error when hashing password: " + err.Error())
	}

	account := &models.Account{
		Email:    req.Email,
		Password: string(hash),
	}

	return s.repository.CreateAccount(ctx, account)
}

func (s *Service) Login(ctx context.Context, req LoginRequest) (string, error) {
	account, err := s.repository.GetAccountByEmail(ctx, req.Email)
	if err != nil {
		return "", err
	}

	if account == nil {
		return "", errors.New("account not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(req.Password)); err != nil {
		return "", errors.New("invalid password")
	}

	return GenerateAccountAuthorizationToken(account), nil
}

func GenerateAccountAuthorizationToken(account *models.Account) string {
	return "token from " + account.Email
}
