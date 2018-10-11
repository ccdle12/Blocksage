// +build unit

package main

import (
	"github.com/ccdle12/Blocksage/crawler/db-client"
	"github.com/ccdle12/Blocksage/crawler/indexer"
	"github.com/ccdle12/Blocksage/crawler/injector"
	"github.com/ccdle12/Blocksage/crawler/models"
	"github.com/ccdle12/Blocksage/crawler/node-client"
	"github.com/ccdle12/Blocksage/crawler/test-utils"
	"github.com/stretchr/testify/suite"
	"testing"
)

// ===========================================================
// Testing Suite
// ===========================================================
type MainUnitSuite struct {
	suite.Suite
	block                *models.Block
	nodeClient           *nodeclient.Client
	nodeClientController nodeclient.Controller
}

// This gets run automatically by `go test` so we call `suite.Run` inside it
func TestSuiteMainUnitSuite(t *testing.T) {
	// This is what actually runs our suite
	suite.Run(t, new(MainUnitSuite))
}

// Testing Lifecycle Hooks
func (suite *MainUnitSuite) SetupTest() {
	suite.block = &models.Block{}
	suite.nodeClient = nodeclient.New(injector.DefaultHTTPClient(), testutils.NodeAddress, testutils.Username, testutils.Password)
	suite.nodeClientController = nodeclient.New(injector.DefaultHTTPClient(), testutils.NodeAddress, testutils.Username, testutils.Password)
}

// ===========================================================
// Unit Tests
// ===========================================================
// TestModelPackageExists will import the models package and create
// a Block object. If the Block object is initialized to all the zero
// values, then we are able to import the models package.
func (suite *MainUnitSuite) TestModelPackageExists() {
	suite.NotNil(suite.block, "Block should have been initialized using the models package import")
	suite.EqualValues(suite.block.Bits, "", "block.Bits should be an empty string")
}

// TestControllersPackageExists will import the controllers package and create
// a Client object. If the Client object is initialized to all the zero
// values, then we are able to import the controllers package.
func (suite *MainUnitSuite) TestControllersPackageExists() {
	suite.NotNil(suite.nodeClient, "Node Client should have been initialized using the controllers package import")
}

// TestReferenceByInterface will test whether Client can be created and referenced
// using the interface.
func (suite *MainUnitSuite) TestReferenceByInterface() {
	suite.NotNil(suite.nodeClientController, "nodeClient was initialized and referenced using the interface")
}

// TestInitIndexer will test that the Crawler can be initalised from main.
func (suite *MainUnitSuite) TestInitIndexer() {
	// Create a node Controller to communicate with a Blockchain Node.
	var nodeClient nodeclient.Controller
	nodeClient = nodeclient.New(
		injector.DefaultHTTPClient(),
		injector.BTCDomain(),
		injector.BTCUsername(),
		injector.BTCPassword())

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
	suite.NoError(err, "There should be no error when create a db client.")

	// Create an Indexer.
	indexer := indexer.New(nodeClient, testDBClient)
	suite.NotNil(indexer, "Indexer should be initialised and not nil")
}
