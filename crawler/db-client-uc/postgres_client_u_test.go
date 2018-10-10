// +build unit

package dbuc

import (
	"github.com/ccdle12/Blocksage/crawler/test-utils"
	"github.com/stretchr/testify/suite"
	"testing"
)

// ===========================================================
// Testing Suite
// ===========================================================
type PostGresClientSuite struct {
	suite.Suite
	postgresClient *PostGresClient
}

// This gets run automatically by `go test` so we call `suite.Run` inside it
func TestSuiteUnitPostGresClient(t *testing.T) {
	// This is what actually runs our suite
	suite.Run(t, new(PostGresClientSuite))
}

// Testing Lifecycle Hooks
func (suite *PostGresClientSuite) SetupTest() {
	suite.postgresClient = NewPostGresClient(testutils.DBConfig)
}

// ===========================================================
// Unit Tests
// ===========================================================
// TestPostGresClientInit will test that we can initialize the handler.
func (suite *PostGresClientSuite) TestPostGresClientInit() {
	suite.NotNil(suite.postgresClient, "postgresHandler should have been initialized")
}

// TestConfigInit will test that the config was initialized.
func (suite *PostGresClientSuite) TestConfigInit() {
	suite.NotNil(suite.postgresClient.cfg, "postgresHandler should have initialized a DBConfig")
	suite.EqualValues(testutils.DBPort, suite.postgresClient.cfg.DBPort, "DBPort in cfg should be the same as the one in testutils")
}
