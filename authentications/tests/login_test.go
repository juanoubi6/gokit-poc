package authentications

import (
	"bytes"
	"encoding/json"
	"gokit-poc/commons"
	"net/http"
	"net/http/httptest"
)

func (suite *AuthenticationsTestSuite) PrepareLoginRequest(email, password string) (*httptest.ResponseRecorder, error) {
	reqBody := map[string]interface{}{
		"email":    email,
		"password": password,
	}
	jsonStr, err := json.Marshal(reqBody)

	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, err
	}

	rr := httptest.NewRecorder()
	suite.TestRouter.ServeHTTP(rr, req)

	return rr, nil
}

func (suite *AuthenticationsTestSuite) TestLoginReturns200AndJWTToken() {
	// Create user before login
	singUpRR, err := suite.PrepareSignUpRequest("loginTest@test.co", "validpassword")
	if err != nil {
		suite.Fail(err.Error())
	}
	suite.Equal(singUpRR.Code, http.StatusCreated, "Expected to be the same")

	// Login
	loginRR, err := suite.PrepareLoginRequest("loginTest@test.co", "validpassword")
	if err != nil {
		suite.Fail(err.Error())
	}
	suite.Equal(loginRR.Code, http.StatusOK, "Expected to be the same")

	var resp commons.GenericResponse
	if err := json.Unmarshal(loginRR.Body.Bytes(), &resp); err != nil {
		suite.Fail(err.Error())
	}

	loginData, ok := resp.Data.(map[string]interface{})
	if !ok {
		suite.Fail("Invalid login response format in data field")
	}

	if _, ok := loginData["token"]; !ok {
		suite.Fail("JWT token not present in login response")
	}

}

func (suite *AuthenticationsTestSuite) TestLoginReturns400WhenEmailDoesNotExist() {
	// Login
	loginRR, err := suite.PrepareLoginRequest("notSignedUpEmail@test.co", "validpassword")
	if err != nil {
		suite.Fail(err.Error())
	}
	suite.Equal(loginRR.Code, http.StatusBadRequest, "Expected to be the same")

	var resp commons.GenericResponse
	if err := json.Unmarshal(loginRR.Body.Bytes(), &resp); err != nil {
		suite.Fail(err.Error())
	}

	expectedMessage := "Account not found"
	suite.Equal(expectedMessage, resp.Message, "Expected to be the same")

}

func (suite *AuthenticationsTestSuite) TestLoginReturns400OnInvalidPassword() {
	// Create user before login
	singUpRR, err := suite.PrepareSignUpRequest("invalidPasswordTest@test.co", "validpassword")
	if err != nil {
		suite.Fail(err.Error())
	}
	suite.Equal(singUpRR.Code, http.StatusCreated, "Expected to be the same")

	// Login
	loginRR, err := suite.PrepareLoginRequest("invalidPasswordTest@test.co", "invalidPassword")
	if err != nil {
		suite.Fail(err.Error())
	}
	suite.Equal(loginRR.Code, http.StatusBadRequest, "Expected to be the same")

	var resp commons.GenericResponse
	if err := json.Unmarshal(loginRR.Body.Bytes(), &resp); err != nil {
		suite.Fail(err.Error())
	}

	expectedMessage := "Invalid password"
	suite.Equal(expectedMessage, resp.Message, "Expected to be the same")

}
