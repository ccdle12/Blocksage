package indexer

import (
	"github.com/ccdle12/Blocksage/crawler/db-client"
	"github.com/ccdle12/Blocksage/crawler/injector"
	"github.com/ccdle12/Blocksage/crawler/node-client"
	"github.com/ccdle12/Blocksage/crawler/test-utils"
	"github.com/stretchr/testify/suite"
	"testing"
)

// ===========================================================
// Testing Suite
// ===========================================================
type IndexerIntegrationSuite struct {
	suite.Suite
}

// This gets run automatically by `go test` so we call `suite.Run` inside it.
func TestSuiteIndexerIntegration(t *testing.T) {
	// This is what actually runs our suite.
	suite.Run(t, new(IndexerIntegrationSuite))
}

// ===========================================================
// Unit Tests
// ===========================================================

// TestIndexerWrite will test that the indexer can write blocks and all
// subsequent information from the block.
func (suite *IndexerIntegrationSuite) TestCrawlReadWrite() {
	// Create a controller to write to the DB using the test tables.
	var testDBClient dbclient.Controller
	testDBClient, err := dbclient.New(
		dbclient.DBPort(injector.PostgresPort()),
		dbclient.DBName(injector.PostgresDBName()),
		dbclient.DBUser(injector.PostgresUserName()),
		dbclient.DBHost(injector.PostgresDomain()),
		dbclient.DBPassword(injector.PostgresPassword()),
		dbclient.PostgresClient(),
		dbclient.Test())

	// Connect to the DB.
	err = testDBClient.Connect()
	defer testDBClient.Close()
	suite.NoError(err, "There should be no error connecting to the DB")

	// Create a node Controller to communicate with a Blockchain Node.
	var nodeClient nodeclient.Controller
	nodeClient = nodeclient.New(
		injector.DefaultHTTPClient(),
		injector.BTCDomain(),
		injector.BTCUsername(),
		injector.BTCPassword())

	// Run the crawler.
	indexer := New(nodeClient, testDBClient)
	err = indexer.write(testutils.ReducedBlock538770)
	suite.NoError(err, "There should be no error when calling write on the indexer.")
}

// TestIndexerGetBlock will test whether the Indexer can make a call via the node controller.
func (suite *IndexerIntegrationSuite) TestIndexerGetBlock() {
	// Create a controller to write to the DB using the test tables.
	var testDBClient dbclient.Controller
	testDBClient, err := dbclient.New(
		dbclient.DBPort(injector.PostgresPort()),
		dbclient.DBName(injector.PostgresDBName()),
		dbclient.DBUser(injector.PostgresUserName()),
		dbclient.DBHost(injector.PostgresDomain()),
		dbclient.DBPassword(injector.PostgresPassword()),
		dbclient.PostgresClient(),
		dbclient.Test())

	// Connect to the DB.
	err = testDBClient.Connect()
	defer testDBClient.Close()
	suite.NoError(err, "There should be no error connecting to the DB")

	// Create a node Controller to communicate with a Blockchain Node.
	var nodeClient nodeclient.Controller
	nodeClient = nodeclient.New(
		injector.DefaultHTTPClient(),
		injector.BTCDomain(),
		injector.BTCUsername(),
		injector.BTCPassword())

	// Create an Indexer.
	indexer := New(nodeClient, testDBClient)
	block, err := indexer.getBlock("0000000000000000001ca03d9e1dd30d2cf49e44ba1569c8819e56cef88b67d4")
	suite.NoError(err, "There should be no error calling GetBlock()")
	suite.NotNil(block, "Block should not be nil")
	suite.Equal(block.Hash, "0000000000000000001ca03d9e1dd30d2cf49e44ba1569c8819e56cef88b67d4", "Retrived block should have the same block hash.")
}
