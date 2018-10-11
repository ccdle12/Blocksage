// +build integration

package main

import (
	"github.com/ccdle12/Blocksage/crawler/db-client"
	"github.com/ccdle12/Blocksage/crawler/injector"
	"github.com/ccdle12/Blocksage/crawler/node-client"
	"github.com/ccdle12/Blocksage/crawler/test-utils"
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
// TestDBConnection will create a DBClient and test we can open a connection to the DB.
func (suite *MainIntegrationSuite) TestDBConnection() {
	// Create a concrete implementation of the DBClient
	dbClient, err := dbclient.New(
		dbclient.DBPort(injector.PostgresPort()),
		dbclient.DBName(injector.PostgresDBName()),
		dbclient.DBUser(injector.PostgresUserName()),
		dbclient.DBHost(injector.PostgresDomain()),
		dbclient.DBPassword(injector.PostgresPassword()),
		dbclient.PostgresClient())
	suite.NoError(err, "Should not return errors when initializing dbClient")

	// Create the connection to the DB
	err = dbClient.Connect()
	defer dbClient.Close()
	suite.NoError(err, "Should not return error when connecting to the DB")
	suite.NotNil(dbClient, "dbClient should not be nil")
}

// TestWriteBlockToTestDB will test that the client can write to the DB.
func (suite *MainIntegrationSuite) TestWriteBlockToTestDB() {
	// Create a DBClient that will be using the test tables.
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
	suite.NoError(err, "There should be no error when connecting to the Test DB.")

	// Write a block to the DB.
	block := testutils.Block506664
	err = testDBClient.WriteBlock(block)
	suite.NoError(err, "There should be no error writing block to db")
}

// TestGetTransation will test that the node client can retrieve a transaction.
func (suite *MainIntegrationSuite) TestGetTransation() {
	// Create a controller to communicate with a Node.
	var nodeClient nodeclient.Controller
	nodeClient = nodeclient.New(
		injector.DefaultHTTPClient(),
		injector.BTCDomain(),
		injector.BTCUsername(),
		injector.BTCPassword())

	// Retrieve transactions from a block.
	block := testutils.Block506664
	txs := block.TX

	// call GetTransactions using the hash of the transactions. This should return
	// details about the transaction.
	tx, err := nodeClient.GetTransaction(txs[3])
	suite.NoError(err, "There should be no error getting transaction 3 from the node.")
	suite.NotNil(tx, "Tx retrieved should not be nil")
	suite.Equal("24438ff6f5e41d70d6b6e9414e965a8e0a9171cfc4333c596640e64694c0292e", tx.Hash)

	// call GetTransactions using the hash of transaction 4. This should return
	// details about the transaction.
	tx, err = nodeClient.GetTransaction(txs[4])
	suite.NoError(err, "There should be no error getting transaction 4 from the node.")
	suite.NotNil(tx, "Tx retrieved should not be nil")
	suite.Equal("b1bdf0b2ac321b00545598043c0554cf6b5c6fa61ce36b30cf7b3addcccc0bf3", tx.Hash)
}

// TestWriteTransaction will test that the db client write a block to the DB.
func (suite *MainIntegrationSuite) TestWriteBlock() {
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
	suite.NoError(err, "There should be no error when init testDBClient")

	// Create connection to the DB.
	err = testDBClient.Connect()
	defer testDBClient.Close()
	suite.NoError(err, "There should be no error connecting to the DB")

	// Retrieve blocks and transactions.
	block := testutils.Block506664
	// txs := block.TX

	// Write the Block to the db.
	err = testDBClient.WriteBlock(block)
	suite.NoError(err, "Should be able to write block to the db")
}

// TestWriteMultipleTXs will test the flow for iterating over a slice of hashes and writing them.
func (suite *MainIntegrationSuite) TestWriteMultipleTXs() {
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

	// Retrieve a Block
	block := testutils.Block506664
	txs := block.TX

	// Write the Block to the db
	err = testDBClient.WriteBlock(block)
	suite.NoError(err, "Should be able to write block to the db")

	// Write 10 transactions to the DB.
	for i, hash := range txs {
		// Break early, this block has +1000 plus txs.
		if i == 10 {
			break
		}

		// Request GetTransaction from the node.
		tx, err := nodeClient.GetTransaction(hash)
		suite.NoError(err, "There should be no error when connecting to the Test DB.")

		// Add blockhash to transaction.
		tx.Blockhash = block.Hash

		// Write the retrieved transaction to the transaction table.
		err = testDBClient.WriteTransaction(tx)
		suite.NoError(err, "There should be no error writing transaction")

		// Loop over each input in the transaction.
		for _, input := range tx.Vin {
			// Write the transaction inputs to the inputs table.
			err = testDBClient.WriteInput(tx.Hash, input)
			suite.NoError(err, "There should be no error writing inputs to the db.")
		}

		// Loop over each output in the transaction.
		for _, output := range tx.Vout {
			// Write the transaction outputs to the outputs table.
			err = testDBClient.WriteOutput(tx.Hash, output)
			suite.NoError(err, "There should be no error writing outputs to the db.")
		}
	}
}

// Test: how to test it?
// [block{next: x}, block{next: nil}, block{next: y}]
// 1. call Crawl(hash)
// 2. call GetBlock(hash)
// 3. If utils.EmptyString(block.NextBlockHash)
// 4. sleep 1 minute
// 6. Go Back to step 2.
// 7. Else call WriteBlock(block)
// 8. call Crawl(block.nextBlockHash)
