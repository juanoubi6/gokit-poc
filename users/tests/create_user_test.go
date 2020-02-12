package users

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
)

func (suite *UsersTestSuite) PrepareCreateUserRequest(jwt string, name string, lastName string, age int) (*httptest.ResponseRecorder, error) {
	reqBody := map[string]interface{}{
		"name":     name,
		"lastName": lastName,
		"age":      age,
	}
	jsonStr, err := json.Marshal(reqBody)

	req, err := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, err
	}

	if jwt != "" {
		req.Header.Set("Authorization", "bearer "+jwt)
	}

	rr := httptest.NewRecorder()
	suite.TestRouter.ServeHTTP(rr, req)

	return rr, nil
}

func (suite *UsersTestSuite) TestCreateUserReturns201OnCreatedUser() {
	jwt := suite.CreateJWTTokenForUser(1, "someEmail@test.com")
	rr, err := suite.PrepareCreateUserRequest(jwt, "TestName", "TestLastName", 20)
	if err != nil {
		suite.Fail(err.Error())
	}

	suite.Equal(rr.Code, http.StatusCreated, "Expected to be the same")

	_, err = suite.ParseResponseBodyToGenericResponse(rr.Body.Bytes())
	if err != nil {
		suite.Fail("Invalid response body")
	}
}

func (suite *UsersTestSuite) TestCreateUserReturns401IfJWTIsMissing() {
	rr, err := suite.PrepareCreateUserRequest("", "TestName", "TestLastName", 20)
	if err != nil {
		suite.Fail(err.Error())
	}

	suite.Equal(rr.Code, http.StatusUnauthorized, "Expected to be the same")
}
