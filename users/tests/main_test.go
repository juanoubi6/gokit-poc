package users

import (
	"gokit-poc/builder"
	"gokit-poc/commons"
	"gokit-poc/models"
	"gokit-poc/security"
	"net/http"
	"os"
	"testing"
)

var TestingRouter http.Handler

func TestMain(m *testing.M) {
	CreateTestingRouter()
	code := m.Run()
	os.Exit(code)
}

func CreateTestingRouter() {
	db := builder.CreateDatabase(commons.TestDatabaseUri)
	TestingRouter = builder.BuildAppRouter(db)
}

func CreateJWTTokenForUser(id uint, email string) string {
	account := models.Account{
		ID:    id,
		Email: email,
	}

	jwt, err := security.CreateAccountJWT(&account)
	if err != nil {
		panic("Unexpected error when creating JWT: " + err.Error())
	}

	return jwt
}
