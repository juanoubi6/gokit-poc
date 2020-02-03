package security

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	kitJWT "github.com/go-kit/kit/auth/jwt"
	"github.com/go-kit/kit/endpoint"
	kitHTTP "github.com/go-kit/kit/transport/http"
	"gokit-poc/commons"
	"gokit-poc/models"
	"net/http"
	"strings"
	"time"
)

const (
	Bearer              string = "bearer"
	AuthorizationHeader string = "Authorization"
	JWTSecretKey        string = "notSoSecret"
	JWTTokenContextKey  string = "authToken"
	JWTClaimsContextKey string = "accountClaims"
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

		return context.WithValue(ctx, JWTTokenContextKey, token)
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
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, newClaims)

	return token.SignedString([]byte(JWTSecretKey))
}

// Rewriting implementation of NewParser from go-kit JWT tools.
func NewParser(keyFunc jwt.Keyfunc, method jwt.SigningMethod, newClaims kitJWT.ClaimsFactory) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			// tokenString is stored in the context from the transport handlers.
			tokenString, ok := ctx.Value(JWTTokenContextKey).(string)
			if !ok {
				return nil, commons.AuthorizationError{"Authorization token was not provided"}
			}
			// Parse takes the token string and a function for looking up the
			// key. The latter is especially useful if you use multiple keys
			// for your application.  The standard is to use 'kid' in the head
			// of the token to identify which key to use, but the parsed token
			// (head and claims) is provided to the callback, providing
			// flexibility.
			token, err := jwt.ParseWithClaims(tokenString, newClaims(), func(token *jwt.Token) (interface{}, error) {
				// Don't forget to validate the alg is what you expect:
				if token.Method != method {
					return nil, commons.AuthorizationError{"Authorization token signing method did not match"}
				}

				return keyFunc(token)
			})
			if err != nil {
				if e, ok := err.(*jwt.ValidationError); ok {
					switch {
					case e.Errors&jwt.ValidationErrorMalformed != 0:
						// Token is malformed
						return nil, commons.AuthorizationError{"Authorization token is malformed"}
					case e.Errors&jwt.ValidationErrorExpired != 0:
						// Token is expired
						return nil, commons.AuthorizationError{"Authorization token expired"}
					case e.Errors&jwt.ValidationErrorNotValidYet != 0:
						// Token is not active yet
						return nil, commons.AuthorizationError{"Authorization token is not valid yet"}
					case e.Inner != nil:
						// report e.Inner
						return nil, commons.AuthorizationError{e.Inner.Error()}
					}
					// We have a ValidationError but have no specific Go kit error for it.
					// Fall through to return original error.
				}
				return nil, commons.AuthorizationError{err.Error()}
			}

			if !token.Valid {
				return nil, commons.AuthorizationError{"Authorization token is invalid"}
			}

			ctx = context.WithValue(ctx, JWTClaimsContextKey, token.Claims)

			return next(ctx, request)
		}
	}
}

func AccountAuthorizationMiddleware() endpoint.Middleware {
	kf := func(tok *jwt.Token) (interface{}, error) {
		return []byte(JWTSecretKey), nil
	}
	claimFactory := func() jwt.Claims {
		return &AccountClaims{}
	}

	return NewParser(kf, jwt.SigningMethodHS256, claimFactory)
}
