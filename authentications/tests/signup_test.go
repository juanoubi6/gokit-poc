package authentications

import (
	"bytes"
	"encoding/json"
	"gokit-poc/commons"
	"gokit-poc/models"
	"net/http"
	"net/http/httptest"
)

func (suite *AuthenticationsTestSuite) PrepareSignUpRequest(email, password string) (*httptest.ResponseRecorder, error) {
	reqBody := map[string]interface{}{
		"email":    email,
		"password": password,
	}
	jsonStr, err := json.Marshal(reqBody)

	req, err := http.NewRequest("POST", "/signup", bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, err
	}

	rr := httptest.NewRecorder()
	suite.TestRouter.ServeHTTP(rr, req)

	return rr, nil

}

func (suite *AuthenticationsTestSuite) TestSignUpReturns201OnCreatedAccount() {
	rr, err := suite.PrepareSignUpRequest("account@test.com", "validpassword")
	if err != nil {
		suite.Fail(err.Error())
	}

	suite.Equal(rr.Code, http.StatusCreated, "Expected to be the same")
}

func (suite *AuthenticationsTestSuite) TestSignUpReturns400OnRepeatedEmail() {
	recorder1, err := suite.PrepareSignUpRequest("repeteadEmailTest@test.com", "validpassword")
	if err != nil {
		suite.Fail(err.Error())
	}
	recorder2, err := suite.PrepareSignUpRequest("repeteadEmailTest@test.com", "validpassword")
	if err != nil {
		suite.Fail(err.Error())
	}

	// First request should do OK
	suite.Equal(recorder1.Code, http.StatusCreated, "Expected to be the same")

	// Second request should fail because the email is already in use
	suite.Equal(recorder2.Code, http.StatusBadRequest, "Expected to be the same")

	var resp commons.GenericResponse
	if err := json.Unmarshal(recorder2.Body.Bytes(), &resp); err != nil {
		suite.Fail(err.Error())
	}

	expectedMessage := "Email already in use"
	suite.Equal(expectedMessage, resp.Message, "Expected to be the same")
}

func (suite *AuthenticationsTestSuite) TestSignUpStoresAccountInDB() {
	testEmail := "testaccountcreated@test.com"
	rr, err := suite.PrepareSignUpRequest(testEmail, "validpassword")
	if err != nil {
		suite.Fail(err.Error())
	}

	suite.Equal(rr.Code, http.StatusCreated, "Account creation failed")

	account := models.Account{}
	if err = suite.db.Where("email = ?", testEmail).First(&account).Error; err != nil {
		suite.Fail("Account not created: " + err.Error())
	}

}
