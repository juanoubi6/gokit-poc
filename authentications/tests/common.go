package authentications

import (
	"gokit-poc/builder"
	"gokit-poc/commons"
	"net/http"
)

var TestingRouter http.Handler

func CreateTestingRouter() {
	db := builder.CreateDatabase(commons.TestDatabaseUri)
	TestingRouter = builder.BuildAppRouter(db)
}
