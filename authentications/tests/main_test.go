package authentications

import (
	"encoding/json"
	"github.com/stretchr/testify/suite"
	"gokit-poc/builder"
	"gokit-poc/commons"
	"net/http"
	"testing"
)

type AuthenticationsTestSuite struct {
	suite.Suite
	TestRouter http.Handler
}

// Runs before the suite tests are run
func (suite *AuthenticationsTestSuite) SetupSuite() {
	db := builder.CreateDatabase(commons.TestDatabaseUri)
	suite.TestRouter = builder.BuildAppRouter(db)
}

func TestAuthenticationsTestSuite(t *testing.T) {
	suite.Run(t, new(AuthenticationsTestSuite))
}

func (suite *AuthenticationsTestSuite) ParseResponseBodyToGenericResponse(responseBody []byte) (*commons.GenericResponse, error) {
	var resp commons.GenericResponse
	if err := json.Unmarshal(responseBody, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
