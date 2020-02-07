package users

import (
	"bytes"
	"encoding/json"
	"gokit-poc/commons"
	"net/http"
	"net/http/httptest"
	"testing"
)

func PrepareCreateUserRequest(jwt string, name string, lastName string, age int) (*httptest.ResponseRecorder, error) {
	reqBody := map[string]interface{}{
		"name":     name,
		"lastName": lastName,
		"age":      age,
	}
	jsonStr, err := json.Marshal(reqBody)

	req, err := http.NewRequest("POST", "/user", bytes.NewBuffer(jsonStr))
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

func TestCreateUserReturns201OnCreatedUser(t *testing.T) {
	jwt := CreateJWTTokenForUser(1, "someEmail@test.com")
	rr, err := PrepareCreateUserRequest(jwt, "TestName", "TestLastName", 20)
	if err != nil {
		t.Fatal(err)
	}

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	var resp commons.GenericResponse
	if err := json.Unmarshal(rr.Body.Bytes(), &resp); err != nil {
		t.Fatal(err)
	}
}

func TestCreateUserReturns401IfJWTIsMissing(t *testing.T) {
	rr, err := PrepareCreateUserRequest("", "TestName", "TestLastName", 20)
	if err != nil {
		t.Fatal(err)
	}

	if status := rr.Code; status != http.StatusUnauthorized {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusUnauthorized)
	}
}