package authentications

import (
	"encoding/json"
	"gokit-poc/builder"
	"gokit-poc/commons"
	"net/http"
	"os"
	"testing"
)

var TestingRouter http.Handler

func TestMain(m *testing.M) {
	CreateTestingRouter()
	code := m.Run()
	os.Exit(code)
}

func CreateTestingRouter() {
	db := builder.CreateDatabase(commons.TestDatabaseUri)
	TestingRouter = builder.BuildAppRouter(db)
}

func ParseResponseBodyToGenericResponse(responseBody []byte) (*commons.GenericResponse, error) {
	var resp commons.GenericResponse
	if err := json.Unmarshal(responseBody, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
