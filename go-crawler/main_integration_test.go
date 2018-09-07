// +build integration

package main

import (
	"github.com/ccdle12/Blocksage/go-crawler/controllers"
	"github.com/ccdle12/Blocksage/go-crawler/test-utils"
	"github.com/stretchr/testify/suite"
	"testing"
)

// NOTE: Ping() in db usecases does NOT respect context timeouts, testing dbClient.Connect()
// will take longer than the other tests.
// ===========================================================
// Testing Suite
// ===========================================================
type MainIntegrationSuite struct {
	suite.Suite
}

// This gets run automatically by `go test` so we call `suite.Run` inside it
func TestSuiteMainIntegration(t *testing.T) {
	// This is what actually runs our suite
	suite.Run(t, new(MainIntegrationSuite))
}

// ===========================================================
// Unit Tests
// ===========================================================
// TestDBConnection will create a DBClientUsecase and test we can open a connection to the DB.
func (suite *MainIntegrationSuite) TestDBConnection() {
	dbClient, err := controllers.NewDBClient(
		controllers.DBHost(testutils.DBHost),
		controllers.DBPort(testutils.DBHost),
		controllers.DBName(testutils.DBName),
		controllers.DBUser(testutils.DBUser),
		controllers.DBPassword(testutils.DBPassword),
		controllers.DBType(testutils.DBType))

	suite.NoError(err, "Should not return errors when initializing dbClient")

	err = dbClient.Connect()
	defer dbClient.Close()

	// TODO (ccdle12): This needs to use injector.DBHost() since these addresses aren't real
	suite.NoError(err, "Should not return error when connecting to the DB")
	suite.NotNil(dbClient, "dbClient should not be nil")
}

// TestDBConnectionShouldFail will create a DBClientUsecase and test we can open a connection to the DB.
func (suite *MainIntegrationSuite) TestDBConnectionShouldFail() {
	dbClient, err := controllers.NewDBClient(
		controllers.DBHost(testutils.DBHost),
		controllers.DBPort(testutils.DBHost),
		controllers.DBName(testutils.DBName),
		controllers.DBUser(testutils.DBUser),
		controllers.DBPassword(testutils.DBPassword),
		controllers.DBType(testutils.DBType))

	suite.NoError(err, "Should not return errors when initializing dbClient")

	err = dbClient.Connect()
	defer dbClient.Close()

	suite.Error(err, "Should return error trying to connect to a DB that doesn't exist")
}
