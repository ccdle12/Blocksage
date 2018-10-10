// +build unit

package nodeuc

import (
	"github.com/ccdle12/Blocksage/go-crawler/injector"
	"github.com/ccdle12/Blocksage/go-crawler/models"
	"github.com/ccdle12/Blocksage/go-crawler/test-utils"
	"github.com/stretchr/testify/suite"
	"testing"
)

// ===========================================================
// Testing Suite
// ===========================================================
type BitcoinClientSuite struct {
	suite.Suite
	bitcoinClient *BitcoinClient
	nodeClient    Usecase
}

// This gets run automatically by `go test` so we call `suite.Run` inside it
func TestSuiteUnitBitcoinClient(t *testing.T) {
	// This is what actually runs our suite
	suite.Run(t, new(BitcoinClientSuite))
}

// Testing Lifecycle Hooks
func (suite *BitcoinClientSuite) SetupTest() {
	// Initializing the concrete class of BitcoinClient Usercase, this will allow testing of un-exported functions
	suite.bitcoinClient = &BitcoinClient{}

	// Initializing the concrete class referenced as the interface
	suite.nodeClient = NewBitcoinClient()
}

// ===========================================================
// Unit Tests
// ===========================================================
// TestCreateBody will test the function to create a body as a *bytes.Buffer to send to the node.
func (suite *BitcoinClientSuite) TestCreateBody() {
	body, err := suite.bitcoinClient.createBody(testutils.NodeBody)

	suite.NoError(err, "Error should be nil.")
	suite.NotNil(body, "Should be able to initialize body.")
}

// TestCreateRequest will test the function to create a request as a *http.Request to send to the node.
func (suite *BitcoinClientSuite) TestCreateRequest() {
	req, err := suite.bitcoinClient.createRequest(testutils.NodeReq)

	suite.NoError(err, "Error should be nil.")
	suite.NotNil(req, "Should be able to initialize request.")

	// req should be well formed with the parameters passed
	suite.EqualValues(req.Method, "POST", "The request method should be POST")
	suite.EqualValues(req.Host, "123.45.67.8:8332", "Request host should be the passed domain")
}

// TestMockSendNodeRequest will test that we can mock a response and send it back.
func (suite *BitcoinClientSuite) TestMockSendNodeRequest() {
	// blockResponse is a mock response
	blockResponse := testutils.NodeResCorrectBlockNoTx0
	server := testutils.TestServer(blockResponse)
	defer server.Close()

	nodeReq := injector.NodeRequest(server.Client(), server.URL, testutils.Username, testutils.Password, testutils.GetBlock, testutils.GetBlockParams)
	res, err := suite.nodeClient.SendNodeRequest(nodeReq)

	// err should be nil
	suite.NoError(err, "Should be able to send a node request to the Node Usecase")
	suite.EqualValues(res.ID, "", "Should return the same response")
	suite.EqualValues(res.Error, models.NodeError{Code: 0, Message: ""}, "Should return the zero value defaults for NodeError struct")
}

// TestGetNodeResponse will test that SendNodeRequest will return a NodeResponse model.
func (suite *BitcoinClientSuite) TestGetNodeResponse() {
	// blockResponse is a mock response
	blockResponse := testutils.NodeResCorrectBlockNoTx0
	server := testutils.TestServer(blockResponse)
	defer server.Close()

	nodeReq := injector.NodeRequest(server.Client(), server.URL, testutils.Username, testutils.Password, testutils.GetBlock, testutils.GetBlockParams)
	res, err := suite.nodeClient.SendNodeRequest(nodeReq)

	// err should be nil
	suite.NoError(err, "Should be able to call SendNodeRequest to the usecase")
	suite.NotNil(res, "Should not have nil for res")
	suite.NotNil(res.Result, "Should not have nil for the Result")
	suite.EqualValues(res.ID, "")
}

// TestMethodNotFoundError will test the client handling a Method Not Found Error from the Node.
func (suite *BitcoinClientSuite) TestMethodNotFoundError() {
	// blockResponse is a mock response
	nodeResponseError := testutils.NodeResErrorMethodNotFound
	server := testutils.TestServer(nodeResponseError)
	defer server.Close()

	nodeReq := injector.NodeRequest(server.Client(), server.URL, testutils.Username, testutils.Password, testutils.GetBlock, testutils.GetBlockParams)
	res, err := suite.nodeClient.SendNodeRequest(nodeReq)

	suite.Error(err, "Should return an error")
	suite.EqualValues(err.Error(), "Method not found")
	suite.Nil(res, "Response should be nil since we have an error")
}

// TestNoBlockParams will test the client handling sending a node request without params.
func (suite *BitcoinClientSuite) TestNoBlockParams() {
	// blockResponse is a mock response
	nodeResponseError := testutils.NodeResNoBlockParams
	server := testutils.TestServer(nodeResponseError)
	defer server.Close()

	nodeReq := injector.NodeRequest(server.Client(), server.URL, testutils.Username, testutils.Password, testutils.GetBlock, []string{""})
	res, err := suite.nodeClient.SendNodeRequest(nodeReq)

	suite.Error(err, "Should return an error")
	suite.NotEqual(err.Error(), "")
	suite.Nil(res, "Response should be nil since we have an error")
}
