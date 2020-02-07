package authentications

import (
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
