// +build integration

package controllers

import (
	"github.com/ccdle12/Blocksage/go-crawler/injector"
	"github.com/ccdle12/Blocksage/go-crawler/test-utils"
	"github.com/stretchr/testify/suite"
	"testing"
)

// ===========================================================
// Testing Suite
// ===========================================================
type DBClientIntegrationSuite struct {
	suite.Suite
	dbClient     *DBClient
	testDBClient *DBClient
}

// This gets run automatically by `go test` so we call `suite.Run` inside it
func TestSuiteIntegrationDBClient(t *testing.T) {
	// This is what actually runs our suite
	suite.Run(t, new(DBClientIntegrationSuite))
}

func (suite *DBClientIntegrationSuite) SetupTest() {
	// DBClient
	dbClient, err := NewDBClient(
		DBPort(injector.PostgresPort()),
		DBName(injector.PostgresDBName()),
		DBUser(injector.PostgresUserName()),
		DBHost(injector.PostgresDomain()),
		DBPassword(injector.PostgresPassword()),
		DBType("postgres"),
		PostgresClient())

	suite.NoError(err, "There should be no error")
	suite.NotNil(dbClient, "DBClient is not nil when attempting to initialize")

	suite.dbClient = dbClient

	// TestDBClient
	testDBClient, err := NewDBClient(
		DBPort(injector.PostgresPort()),
		DBName(injector.PostgresDBName()),
		DBUser(injector.PostgresUserName()),
		DBHost(injector.PostgresDomain()),
		DBPassword(injector.PostgresPassword()),
		DBType("postgres"),
		TestPostgresClient())

	suite.NoError(err, "There should be no error")
	suite.NotNil(testDBClient, "DBClient is not nil when attempting to initialize")

	suite.testDBClient = testDBClient
}

// ===========================================================
// Integration Tests
// ===========================================================
// TestConnectDBClient will test that the a connection to the DB can be created by the DBClient.
func (suite *DBClientIntegrationSuite) TestConnectDBClient() {
	err := suite.dbClient.Connect()
	defer suite.dbClient.Close()

	suite.NoError(err, "There should be no error when connecting to the DB.")
}

// TestConnectTestDBClient will test that the a connection to the Test DB can be created by the DBClient.
func (suite *DBClientIntegrationSuite) TestConnectTestDBClient() {
	err := suite.testDBClient.Connect()
	defer suite.testDBClient.Close()

	suite.NoError(err, "There should be no error when connecting to the Test DB.")
}

// TestWriteBlockToTestDB will test that client can write to the DB.
func (suite *DBClientIntegrationSuite) TestWriteBlockToTestDB() {
	err := suite.testDBClient.Connect()
	defer suite.testDBClient.Close()

	suite.NoError(err, "There should be no error when connecting to the Test DB.")

	err = suite.testDBClient.WriteBlock(testutils.Block538770)
	suite.NoError(err, "There should be no error when writing the block to the DB.")
}
