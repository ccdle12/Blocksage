// +build integration

package dbuc

import (
	"github.com/ccdle12/Blocksage/go-crawler/test-utils"
	"github.com/stretchr/testify/suite"
	"testing"
)

// ===========================================================
// Testing Suite
// ===========================================================
type PostGresClientIntegrationSuite struct {
	suite.Suite
	postgresClient *PostGresClient
}

// This gets run automatically by `go test` so we call `suite.Run` inside it
func TestSuiteIntegrationPostGresClient(t *testing.T) {
	// This is what actually runs our suite
	suite.Run(t, new(PostGresClientIntegrationSuite))
}

// Testing Lifecycle Hooks
func (suite *PostGresClientIntegrationSuite) SetupTest() {
	suite.postgresClient = NewPostGresClient(testutils.DBConfig)
}

// ===========================================================
// Integration Tests
// ===========================================================
// TestPostGresClientInit will test that we can initialize the handler.
func (suite *PostGresClientIntegrationSuite) TestPostGresClientInit() {
	suite.NotNil(suite.postgresClient, "postgresClient should have been initialized")
}

// TestPostGresClientConnect will test that we can connect to the DB.
func (suite *PostGresClientIntegrationSuite) TestPostGresClientConnect() {
	err := suite.postgresClient.OpenConnection()
	defer suite.postgresClient.CloseConnection()

	suite.NoError(err, "There should be no error opening a connection")
	suite.NotNil(suite.postgresClient.db, "The db in postgresclient should not be nil")

	err = suite.postgresClient.ping()
	suite.NoError(err, "Should be able to ping db")
}
