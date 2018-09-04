package main

import (
	"github.com/ccdle12/Blocksage/go-crawler/controllers"
	"github.com/ccdle12/Blocksage/go-crawler/models"
	"github.com/ccdle12/Blocksage/go-crawler/test-utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestModelPackageExists will import the models package and create
// a Block object. If the Block object is initialized to all the zero
// values, then we are able to import the models package.
func TestModelPackageExists(t *testing.T) {
	assert := assert.New(t)
	block := &models.Block{}

	assert.NotNil(block, "Block should have been initialized using the models package import")
	assert.EqualValues(block.Bits, "", "block.Bits should be an empty string")
}

// TestControllersPackageExists will import the controllers package and create
// a NodeClient object. If the NodeClient object is initialized to all the zero
// values, then we are able to import the controllers package.
func TestControllersPackageExists(t *testing.T) {
	assert := assert.New(t)
	nodeClient := controllers.NewNodeClient(testutils.Client, testutils.NodeAddress, testutils.Username, testutils.Password)

	assert.NotNil(nodeClient, "Node Client should have been initialized using the controllers package import")
}

// TestReferenceByInterface will test whether NodeClient can be created and referenced
// using the interface.
func TestReferenceByInterface(t *testing.T) {
	assert := assert.New(t)
	var nodeClient controllers.NodeClientController

	nodeClient = controllers.NewNodeClient(testutils.Client, testutils.NodeAddress, testutils.Username, testutils.Password)
	assert.NotNil(nodeClient, "nodeClient was initialized and referenced using the interface")
}
