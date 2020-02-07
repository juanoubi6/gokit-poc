package authentications

import (
	"bytes"
	"encoding/json"
	"gokit-poc/commons"
	"gokit-poc/models"
	"net/http"
	"net/http/httptest"
	"testing"
)

func PrepareSignUpRequest(email, password string) (*httptest.ResponseRecorder, error) {
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
	TestingRouter.ServeHTTP(rr, req)

	return rr, nil

}

func TestSignUpReturns201OnCreatedAccount(t *testing.T) {
	rr, err := PrepareSignUpRequest("account@test.com", "validpassword")
	if err != nil {
		t.Fatal(err)
	}

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}
}

func TestSignUpReturns400OnRepeatedEmail(t *testing.T) {
	recorder1, err := PrepareSignUpRequest("repeteadEmailTest@test.com", "validpassword")
	if err != nil {
		t.Fatal(err)
	}
	recorder2, err := PrepareSignUpRequest("repeteadEmailTest@test.com", "validpassword")
	if err != nil {
		t.Fatal(err)
	}

	// First request should do OK
	if status := recorder1.Code; status != http.StatusCreated {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	// Second request should fail because the email is already in use
	if status := recorder2.Code; status != http.StatusBadRequest {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	var resp commons.GenericResponse
	if err := json.Unmarshal(recorder2.Body.Bytes(), &resp); err != nil {
		t.Fatal(err)
	}

	expectedMessage := "Email already in use"
	if resp.Message != expectedMessage {
		t.Errorf("Response returned a different message: got %v want %v", resp.Message, expectedMessage)
	}
}

func TestSignUpStoresAccountInDB(t *testing.T) {
	testEmail := "testaccountcreated@test.com"
	rr, err := PrepareSignUpRequest(testEmail, "validpassword")
	if err != nil {
		t.Fatal(err)
	}

	if rr.Code != http.StatusCreated {
		t.Errorf("Account creation failed: " + rr.Body.String())
	}

	account := models.Account{}
	if err = commons.GlobalDB.Where("email = ?", testEmail).First(&account).Error; err != nil {
		t.Errorf("Account not created: " + err.Error())
	}

}
