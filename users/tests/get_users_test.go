package users

import (
	"fmt"
	"gokit-poc/commons"
	"gokit-poc/models"
	"gokit-poc/users"
	"net/http"
	"net/http/httptest"
	"testing"
)

func PrepareGetUsersRequest(jwt string, name string, lastName string, age int) (*httptest.ResponseRecorder, error) {
	url := fmt.Sprintf("/users?name=%v&lastName=%v&age=%v", name, lastName, age)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	if jwt != "" {
		req.Header.Set("Authorization", "bearer "+jwt)
	}

	rr := httptest.NewRecorder()
	TestingRouter.ServeHTTP(rr, req)

	return rr, nil
}

func TestGetUsersReturns200(t *testing.T) {
	jwt := CreateJWTTokenForUser(1, "someEmail@test.com")
	rr, err := PrepareGetUsersRequest(jwt, "", "", 0)
	if err != nil {
		t.Fatal(err)
	}

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	_, err = ParseResponseBodyToGenericResponse(rr.Body.Bytes())
	if err != nil {
		t.Errorf("Invalid response body")
	}

}

func TestGetUsersWithQueryParamsReturnsSpecifiedUser(t *testing.T) {
	if err := commons.GlobalDB.Save(&models.User{Name: "TestQueryParam", LastName: "TestLastNameQP"}).Error; err != nil {
		t.Fatal(err)
	}

	jwt := CreateJWTTokenForUser(1, "someEmail@test.com")
	rr, err := PrepareGetUsersRequest(jwt, "TestQueryParam", "", 0)
	if err != nil {
		t.Fatal(err)
	}

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	var getUsersResp users.GetUsersResponse
	if err := ParseResponseDataToStruct(rr.Body.Bytes(), &getUsersResp); err != nil {
		t.Fatal("Unexpected error when parsing response data")
	}

	if len(getUsersResp.Users) > 1 {
		t.Errorf("Expected only 1 result,got %v", len(getUsersResp.Users))
	}

	if getUsersResp.Users[0].LastName != "TestLastNameQP" {
		t.Errorf("Expected user last name to be  TestLastNameQP,got %v", getUsersResp.Users[0].LastName)
	}

}
