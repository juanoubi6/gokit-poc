package users

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/suite"
	"gokit-poc/builder"
	"gokit-poc/commons"
	"gokit-poc/models"
	"gokit-poc/security"
	"net/http"
	"testing"
)

type UsersTestSuite struct {
	db *gorm.DB
	suite.Suite
	TestRouter http.Handler
}

// Runs before the suite tests are run
func (suite *UsersTestSuite) SetupSuite() {
	suite.db = builder.CreateDatabase(commons.TestDatabaseUri)
	suite.TestRouter = builder.BuildAppRouter(suite.db)
}

func (suite *UsersTestSuite) CreateJWTTokenForUser(id uint, email string) string {
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

func (suite *UsersTestSuite) ParseResponseDataToStruct(responseBody []byte, container interface{}) error {
	response, err := suite.ParseResponseBodyToGenericResponse(responseBody)
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

func (suite *UsersTestSuite) ParseResponseBodyToGenericResponse(responseBody []byte) (*commons.GenericResponse, error) {
	var resp commons.GenericResponse
	if err := json.Unmarshal(responseBody, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}

func TestUsersTestSuite(t *testing.T) {
	suite.Run(t, new(UsersTestSuite))
}
