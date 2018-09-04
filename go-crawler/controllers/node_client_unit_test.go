// +build unit

package controllers

import (
	"github.com/ccdle12/Blocksage/go-crawler/test-utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestInitializingNodeClient will test that the Node Client class will be
// initialized correctly.
func TestInitializingNodeClient(t *testing.T) {
	nodeClient := NewNodeClient(testutils.Client, testutils.NodeAddress, testutils.Username, testutils.Password)

	// nodeClient should be initialized
	assert.NotNil(t, nodeClient, "Should be able to initialize NodeClient.")
}

// TestGetBlock will call GetBlock() and return a Block Struct. The test will be using a mock test server.
func TestMockSendNodeRequest(t *testing.T) {
	assert := assert.New(t)

	// Create Test Server and pass it to NewNodeClient()
	server := testutils.TestServer(testutils.NodeResCorrectBlockNoTx0)
	defer server.Close()

	nodeClient := NewNodeClient(server.Client(), server.URL, testutils.Username, testutils.Password)
	assert.EqualValues(nodeClient.address, server.URL)

	block, err := nodeClient.GetBlock("0000000000000000001ca03d9e1dd30d2cf49e44ba1569c8819e56cef88b67d4")

	assert.NoError(err, "There should be no error when getting NodeResCorrectBlockNoTx0")
	assert.NotNil(block, "Returned block should not be nil")
	assert.EqualValues("000000000000000000000000000000000000000002eb51495ec06b0a5427f048", block.Chainwork, "Chain Work should equal the chainwork found in block-respones.go")
	assert.EqualValues(3, block.Confirmations, "Confirmations should equal confirmations found in block-respones.go")
	assert.EqualValues(6727225469722.534, block.Difficulty, "Difficulty should be a float64 that matches the model in block_responses.go")
}

// TestSendRequestToMalformedServer will attempt to send a request to a server that is offline or not online at all.
func TestSendRequestToMalformedServer(t *testing.T) {
	assert := assert.New(t)

	// Create Test Server and pass it to NewNodeClient()
	server := testutils.TestServer(testutils.NodeResCorrectBlockNoTx0)
	defer server.Close()

	nodeClient := NewNodeClient(server.Client(), "http://localhost:3421", testutils.Username, testutils.Password)
	_, err := nodeClient.GetBlock("0000000000000000001ca03d9e1dd30d2cf49e44ba1569c8819e56cef88b67d4")

	assert.Error(err, "There should be an error")
}
