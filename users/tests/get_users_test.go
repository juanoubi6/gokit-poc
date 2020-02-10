package users

import (
	"encoding/json"
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

	var resp commons.GenericResponse
	if err := json.Unmarshal(rr.Body.Bytes(), &resp); err != nil {
		t.Fatal(err)
	}

}

func TestGetUsersWithQueryParamsReturnsSpecifiedUser(t *testing.T) {
	if err := commons.GlobalDB.Save(&models.User{Name: "TestQueryParam", LastName: "TestLastNameQP"}).Error; err != nil {
		t.Fatal(err)
	}

	jwt := CreateJWTTokenForUser(1, "someEmail@test.com")
	rr, err := PrepareGetUsersRequest(jwt, "", "", 0)
	if err != nil {
		t.Fatal(err)
	}

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	var resp commons.GenericResponse
	if err := json.Unmarshal(rr.Body.Bytes(), &resp); err != nil {
		t.Fatal(err)
	}

	var getUsersResp users.GetUsersResponse

	dec := json.NewDecoder(rr.Body)
	dec.DisallowUnknownFields()
	if err := dec.Decode(&getUsersResp); err != nil {
		t.Fatal(err)
	}

}
