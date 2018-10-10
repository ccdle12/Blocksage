// +build integration

package dbclient

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
	dbClient     Controller
	testDBClient Controller
}

// This gets run automatically by `go test` so we call `suite.Run` inside it
func TestSuiteIntegrationDBClient(t *testing.T) {
	// This is what actually runs our suite
	suite.Run(t, new(DBClientIntegrationSuite))
}

func (suite *DBClientIntegrationSuite) SetupTest() {
	// DBClient
	dbClient, err := New(
		DBPort(injector.PostgresPort()),
		DBName(injector.PostgresDBName()),
		DBUser(injector.PostgresUserName()),
		DBHost(injector.PostgresDomain()),
		DBPassword(injector.PostgresPassword()),
		PostgresClient())

	suite.NoError(err, "There should be no error")
	suite.NotNil(dbClient, "DBClient is not nil when attempting to initialize")

	suite.dbClient = dbClient

	// TestDBClient
	testDBClient, err := New(
		DBPort(injector.PostgresPort()),
		DBName(injector.PostgresDBName()),
		DBUser(injector.PostgresUserName()),
		DBHost(injector.PostgresDomain()),
		DBPassword(injector.PostgresPassword()),
		PostgresClient(),
		Test())

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

// TestDBClientWriteBlock will test that we write a block to the DB.
func (suite *DBClientIntegrationSuite) TestDBClientWriteBlock() {
	// Connect to the DB.
	err := suite.testDBClient.Connect()
	defer suite.testDBClient.Close()

	suite.NoError(err, "There should be no error opening a connection")

	// Write Block to DB.
	block := testutils.Block538770
	err = suite.testDBClient.WriteBlock(block)
	suite.NoError(err, "There should be no error writing a block")
}

// TestPostGresClientWriteTx will test that we write a TX to the DB.
func (suite *DBClientIntegrationSuite) TestPostGresClientWriteTx() {
	err := suite.testDBClient.Connect()
	defer suite.testDBClient.Close()

	// Use Block538770 as the block to reference this tx
	block := testutils.Block538770

	// Retrieve same tx and add the blockhash of Block538770
	tx := testutils.SampleTX
	tx.Blockhash = block.Hash

	// Write the tx to the db
	err = suite.testDBClient.WriteTransaction(tx)
	suite.NoError(err, "There should be no error writing tx")

}
