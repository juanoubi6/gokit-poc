package authentications

import (
	"bytes"
	"encoding/json"
	"gokit-poc/commons"
	"net/http"
	"net/http/httptest"
	"testing"
)

func PrepareLoginRequest(email, password string) (*httptest.ResponseRecorder, error) {
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
	TestingRouter.ServeHTTP(rr, req)

	return rr, nil
}

func TestLoginReturns200AndJWTToken(t *testing.T) {
	// Create user before login
	singUpRR, err := PrepareSignUpRequest("loginTest@test.co", "validpassword")
	if err != nil {
		t.Fatal(err)
	}
	if status := singUpRR.Code; status != http.StatusCreated {
		t.Errorf("Sign up returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	// Login
	loginRR, err := PrepareLoginRequest("loginTest@test.co", "validpassword")
	if err != nil {
		t.Fatal(err)
	}
	if status := loginRR.Code; status != http.StatusOK {
		t.Errorf("Login returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	var resp commons.GenericResponse
	if err := json.Unmarshal(loginRR.Body.Bytes(), &resp); err != nil {
		t.Fatal(err)
	}

	loginData, ok := resp.Data.(map[string]interface{})
	if !ok {
		t.Errorf("Invalid login response format in data field")
	}

	if _, ok := loginData["token"]; !ok {
		t.Errorf("JWT token not present in login response")
	}

}

func TestLoginReturns400WhenEmailDoesNotExist(t *testing.T) {
	// Login
	loginRR, err := PrepareLoginRequest("notSignedUpEmail@test.co", "validpassword")
	if err != nil {
		t.Fatal(err)
	}
	if status := loginRR.Code; status != http.StatusBadRequest {
		t.Errorf("Login returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}

	var resp commons.GenericResponse
	if err := json.Unmarshal(loginRR.Body.Bytes(), &resp); err != nil {
		t.Fatal(err)
	}

	expectedMessage := "Account not found"
	if resp.Message != expectedMessage {
		t.Errorf("Response returned a different message: got %v want %v", resp.Message, expectedMessage)
	}

}

func TestLoginReturns400OnInvalidPassword(t *testing.T) {
	// Create user before login
	singUpRR, err := PrepareSignUpRequest("invalidPasswordTest@test.co", "validpassword")
	if err != nil {
		t.Fatal(err)
	}
	if status := singUpRR.Code; status != http.StatusCreated {
		t.Errorf("Sign up returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	// Login
	loginRR, err := PrepareLoginRequest("invalidPasswordTest@test.co", "invalidPassword")
	if err != nil {
		t.Fatal(err)
	}
	if status := loginRR.Code; status != http.StatusBadRequest {
		t.Errorf("Login returned wrong status code: got %v want %v", status, http.StatusBadRequest)
	}

	var resp commons.GenericResponse
	if err := json.Unmarshal(loginRR.Body.Bytes(), &resp); err != nil {
		t.Fatal(err)
	}

	expectedMessage := "Invalid password"
	if resp.Message != expectedMessage {
		t.Errorf("Response returned a different message: got %v want %v", resp.Message, expectedMessage)
	}

}
