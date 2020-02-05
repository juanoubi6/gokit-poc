package authentications

import (
	"bytes"
	"encoding/json"
	"gokit-poc/commons"
	"gokit-poc/models"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	CreateTestingRouter()
	code := m.Run()
	os.Exit(code)
}

func TestSignUpReturns201OnCreatedAccount(t *testing.T) {
	reqBody := map[string]interface{}{
		"email":    "account@test.com",
		"password": "validpassword",
	}
	jsonStr, err := json.Marshal(reqBody)

	req, err := http.NewRequest("POST", "/signup", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	TestingRouter.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}
}

func TestSignUpStoresAccountInDB(t *testing.T) {
	testEmail := "testaccountcreated@test.com"

	reqBody := map[string]interface{}{
		"email":    testEmail,
		"password": "validpassword",
	}
	jsonStr, err := json.Marshal(reqBody)

	req, err := http.NewRequest("POST", "/signup", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	TestingRouter.ServeHTTP(rr, req)

	if rr.Code != http.StatusCreated {
		t.Errorf("Account creation failed: " + rr.Body.String())
	}

	account := models.Account{}
	if err = commons.GlobalDB.Where("email = ?", testEmail).First(&account).Error; err != nil {
		t.Errorf("Account not created: " + err.Error())
	}

}
