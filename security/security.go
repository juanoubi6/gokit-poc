package security

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	kitJWT "github.com/go-kit/kit/auth/jwt"
	kitHTTP "github.com/go-kit/kit/transport/http"
	"gokit-poc/models"
	"net/http"
	"strings"
	"time"
)

const (
	Bearer              string = "bearer"
	AuthorizationHeader string = "Authorization"
	JWTSecretKey        string = "notSoSecret"
)

type AccountClaims struct {
	Id    uint
	Email string
	jwt.StandardClaims
}

// AuthTokenToContext moves a JWT from request header to context
func AuthTokenToContext() kitHTTP.RequestFunc {
	return func(ctx context.Context, r *http.Request) context.Context {
		token, ok := extractTokenFromAuthorizationHeader(r.Header.Get(AuthorizationHeader))
		if !ok {
			return ctx
		}

		return context.WithValue(ctx, kitJWT.JWTTokenContextKey, token)
	}
}

func extractTokenFromAuthorizationHeader(val string) (token string, ok bool) {
	authHeaderParts := strings.Split(val, " ")
	if len(authHeaderParts) != 2 || !strings.EqualFold(authHeaderParts[0], Bearer) {
		return "", false
	}

	return authHeaderParts[1], true
}

func CreateAccountJWT(account *models.Account) (string, error) {
	newClaims := AccountClaims{
		Id:    account.ID,
		Email: account.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(5 * time.Minute).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, newClaims)

	return token.SignedString([]byte(JWTSecretKey))
}
