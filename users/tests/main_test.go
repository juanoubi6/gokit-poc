package users

import (
	"encoding/json"
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

func ParseResponseDataToStruct(responseBody []byte, container interface{}) error {
	response, err := ParseResponseBodyToGenericResponse(responseBody)
	if err != nil {
		return err
	}

	jsonData, err := json.Marshal(response.Data.(map[string]interface{}))
	if err != nil {
		return err
	}

	if err := json.Unmarshal(jsonData, &container); err != nil {
		return err
	}

	return nil
}

func ParseResponseBodyToGenericResponse(responseBody []byte) (*commons.GenericResponse, error) {
	var resp commons.GenericResponse
	if err := json.Unmarshal(responseBody, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
