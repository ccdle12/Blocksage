// +build integration

package controllers_test

import (
	"github.com/ccdle12/Blocksage/go-crawler/controllers"
	"github.com/ccdle12/Blocksage/go-crawler/test-utils"
	"github.com/stretchr/testify/suite"
	"os"
	"testing"
)

// ===========================================================
// Testing Suite
// ===========================================================
type NodeClientIntegrationSuite struct {
	suite.Suite
	nodeClient controllers.NodeClientController
}

// This gets run automatically by `go test` so we call `suite.Run` inside it
func TestSuiteIntegrationNodeClient(t *testing.T) {
	// This is what actually runs our suite
	suite.Run(t, new(NodeClientIntegrationSuite))
}

//TODO (ccdle12): Use OS.getenv here, but then import it into utils
// Testing Lifecycle Hooks
func (suite *NodeClientIntegrationSuite) SetupTest() {
	suite.nodeClient = controllers.NewNodeClient(testutils.Client, os.Getenv("BTC_MAIN_DOMAIN"), os.Getenv("USERNAME"), os.Getenv("PASSWORD"))
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
}
