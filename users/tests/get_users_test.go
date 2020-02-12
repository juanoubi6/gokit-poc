package users

import (
	"fmt"
	"gokit-poc/commons"
	"gokit-poc/models"
	"gokit-poc/users"
	"net/http"
	"net/http/httptest"
)

func (suite *UsersTestSuite) PrepareGetUsersRequest(jwt string, name string, lastName string, age int) (*httptest.ResponseRecorder, error) {
	url := fmt.Sprintf("/users?name=%v&lastName=%v&age=%v", name, lastName, age)

	req, err := http.NewRequest("GET", url, nil)
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

func (suite *UsersTestSuite) TestGetUsersReturns200() {
	jwt := suite.CreateJWTTokenForUser(1, "someEmail@test.com")
	rr, err := suite.PrepareGetUsersRequest(jwt, "", "", 0)
	if err != nil {
		suite.Fail(err.Error())
	}

	suite.Equal(rr.Code, http.StatusOK, "Expected to be the same")

	_, err = suite.ParseResponseBodyToGenericResponse(rr.Body.Bytes())
	if err != nil {
		suite.Fail("Invalid response body")
	}

}

func (suite *UsersTestSuite) TestGetUsersWithQueryParamsReturnsSpecifiedUser() {
	if err := commons.GlobalDB.Save(&models.User{Name: "TestQueryParam", LastName: "TestLastNameQP"}).Error; err != nil {
		suite.Fail(err.Error())
	}

	jwt := suite.CreateJWTTokenForUser(1, "someEmail@test.com")
	rr, err := suite.PrepareGetUsersRequest(jwt, "TestQueryParam", "", 0)
	if err != nil {
		suite.Fail(err.Error())
	}

	suite.Equal(rr.Code, http.StatusOK, "Expected to be the same")

	var getUsersResp users.GetUsersResponse
	if err := suite.ParseResponseDataToStruct(rr.Body.Bytes(), &getUsersResp); err != nil {
		suite.Fail("Unexpected error when parsing response data")
	}

	suite.Len(getUsersResp.Users, 1, "Expected only 1 result")
	suite.Equal(getUsersResp.Users[0].LastName, "TestLastNameQP", "Expected to be the same")
}

func (suite *UsersTestSuite) TestGetUsersReturns401IfJWTIsMissing() {
	rr, err := suite.PrepareGetUsersRequest("", "", "", 0)
	if err != nil {
		suite.Fail(err.Error())
	}

	suite.Equal(rr.Code, http.StatusUnauthorized, "Expected to be the same")
}
