// +build integration

package nodeclient

import (
	"github.com/ccdle12/Blocksage/crawler/injector"
	"github.com/stretchr/testify/suite"
	"testing"
)

// ===========================================================
// Testing Suite
// ===========================================================
type NodeClientIntegrationSuite struct {
	suite.Suite
	nodeClient Controller
}

// This gets run automatically by `go test` so we call `suite.Run` inside it
func TestSuiteIntegrationNodeClient(t *testing.T) {
	// This is what actually runs our suite
	suite.Run(t, new(NodeClientIntegrationSuite))
}

// Testing Lifecycle Hooks
func (suite *NodeClientIntegrationSuite) SetupTest() {
	suite.nodeClient = New(injector.DefaultHTTPClient(), injector.BTCDomain(), injector.BTCUsername(), injector.BTCPassword())
}

// ===========================================================
// Integration Tests
// ===========================================================
// TestRequestBlockFromNode will send a request to the Node and request a block.
func (suite *NodeClientIntegrationSuite) TestRequestBlockFromNode() {
	suite.NotNil(suite.nodeClient, "Node Client should not be nil")

	block, err := suite.nodeClient.GetBlock("000000000000000000170b3da1e60f139c603e659b378753023c96c275169eea")

	suite.NoError(err, "There should be no errors when requesting a block")
	suite.NotNil(block, "Block should not be nil")
	suite.NotNil(block.MerkleRoot, "Block MerkleRoot should not be nil")
}

// TestRequestTxFromNode will send a request to the Node and request a transaction.
func (suite *NodeClientIntegrationSuite) TestRequestTxFromNode() {
	suite.NotNil(suite.nodeClient, "Node Client should not be nil")

	// Test retrieving transaction data
	tx, err := suite.nodeClient.GetTransaction("a1d89b5b43fcfe0dfc0b28847d701ff1ef11c6714a54eb7c7b92f10c973565d4")
	suite.NoError(err, "There should be no errors when requesting a block")
	suite.NotNil(tx, "Transaction should not be nil")
	suite.Equal("a1d89b5b43fcfe0dfc0b28847d701ff1ef11c6714a54eb7c7b92f10c973565d4", tx.TXID)
	suite.Equal(0, tx.Locktime)
	suite.Equal(3, len(tx.Vout))

	// Test retrieve transaction output data
	txOutput3 := tx.Vout[2]
	suite.Equal(0.80339718, txOutput3.Value)
	suite.Equal(2, txOutput3.N)

	// Test retrieve ScriptPubKey from the transaction output
	txOutSPK := txOutput3.ScriptPubKey
	suite.Equal("OP_DUP OP_HASH160 2f41d05b45b21a8bddb4aeaeaa61505c14178e17 OP_EQUALVERIFY OP_CHECKSIG", txOutSPK.Asm)
}
