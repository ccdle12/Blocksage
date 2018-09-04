// +build unit

package usecases

import (
	"github.com/ccdle12/Blocksage/go-crawler/models"
	"github.com/ccdle12/Blocksage/go-crawler/test-utils"
	"github.com/ccdle12/Blocksage/go-crawler/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestCreateBody will test the function to create a body as a *bytes.Buffer to send to the node.
func TestCreateBody(t *testing.T) {
	assert := assert.New(t)

	// Initializing the concrete class of BitcoinClient Usercase, this will allow testing of un-exported functions
	b := &BitcoinClient{}
	body, err := b.createBody(testutils.NodeBody)

	assert.NoError(err, "Error should be nil.")
	assert.NotNil(body, "Should be able to initialize body.")
}

// TestCreateRequest will test the function to create a request as a *http.Request to send to the node.
func TestCreateRequest(t *testing.T) {
	assert := assert.New(t)

	b := &BitcoinClient{}
	req, err := b.createRequest(testutils.NodeReq)

	assert.NoError(err, "Error should be nil.")
	assert.NotNil(req, "Should be able to initialize request.")

	// req should be well formed with the parameters passed
	assert.EqualValues(req.Method, "POST", "The request method should be POST")
	assert.EqualValues(req.Host, "123.45.67.8:8332", "Request host should be the passed domain")
}

// TestMockSendNodeRequest will test that we can mock a response and send it back.
func TestMockSendNodeRequest(t *testing.T) {
	assert := assert.New(t)

	// Initializing the Node Usercase
	nodeClient := NewBitcoinClient()

	// blockResponse is a mock response
	blockResponse := testutils.NodeResCorrectBlockNoTx0
	server := testutils.TestServer(blockResponse)
	defer server.Close()

	nodeReq := utils.NodeRequest(server.Client(), server.URL, testutils.Username, testutils.Password, testutils.GetBlock, testutils.GetBlockParams)
	res, err := nodeClient.SendNodeRequest(nodeReq)

	// err should be nil
	assert.NoError(err, "Should be able to send a node request to the Node Usecase")
	assert.EqualValues(res.ID, "", "Should return the same response")
	assert.EqualValues(res.Error, models.NodeError{Code: 0, Message: ""}, "Should return the zero value defaults for NodeError struct")
}

// TestGetNodeResponse will test that SendNodeRequest will return a NodeResponse model.
func TestGetNodeResponse(t *testing.T) {
	assert := assert.New(t)

	// Initializing the Node Usercase
	nodeClient := NewBitcoinClient()

	// blockResponse is a mock response
	blockResponse := testutils.NodeResCorrectBlockNoTx0
	server := testutils.TestServer(blockResponse)
	defer server.Close()

	nodeReq := utils.NodeRequest(server.Client(), server.URL, testutils.Username, testutils.Password, testutils.GetBlock, testutils.GetBlockParams)
	res, err := nodeClient.SendNodeRequest(nodeReq)

	// err should be nil
	assert.NoError(err, "Should be able to call SendNodeRequest to the usecase")
	assert.NotNil(res, "Should not have nil for res")
	assert.NotNil(res.Result, "Should not have nil for the Result")
	assert.EqualValues(res.ID, "")
}

// TestMethodNotFoundError will test the client handling a Method Not Found Error from the Node.
func TestMethodNotFoundError(t *testing.T) {
	assert := assert.New(t)

	// Initializing the Node Usercase
	nodeClient := NewBitcoinClient()

	// blockResponse is a mock response
	nodeResponseError := testutils.NodeResErrorMethodNotFound
	server := testutils.TestServer(nodeResponseError)
	defer server.Close()

	nodeReq := utils.NodeRequest(server.Client(), server.URL, testutils.Username, testutils.Password, testutils.GetBlock, testutils.GetBlockParams)
	res, err := nodeClient.SendNodeRequest(nodeReq)

	assert.Error(err, "Should return an error")
	assert.EqualValues(err.Error(), "Method not found")
	assert.Nil(res, "Response should be nil since we have an error")
}

// TestMethodNotFoundError will test the client handling a Method Not Found Error from the Node.
func TestNoBlockParams(t *testing.T) {
	assert := assert.New(t)

	// Initializing the Node Usercase
	nodeClient := NewBitcoinClient()

	// blockResponse is a mock response
	nodeResponseError := testutils.NodeResNoBlockParams
	server := testutils.TestServer(nodeResponseError)
	defer server.Close()

	nodeReq := utils.NodeRequest(server.Client(), server.URL, testutils.Username, testutils.Password, testutils.GetBlock, []string{""})
	res, err := nodeClient.SendNodeRequest(nodeReq)

	assert.Error(err, "Should return an error")
	assert.NotEqual(err.Error(), "")
	assert.Nil(res, "Response should be nil since we have an error")
}
