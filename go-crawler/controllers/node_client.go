package controllers

import (
	"github.com/ccdle12/Blocksage/go-crawler/models"
	"github.com/ccdle12/Blocksage/go-crawler/usecases"
	"github.com/ccdle12/Blocksage/go-crawler/utils"
	"net/http"
)

// NodeClient is a struct that will handle the connection to the Bitcoin Nodes.
type NodeClient struct {
	client   *http.Client
	address  string
	username string
	password string
	usecase  usecases.NodeClientUsecase
}

// NewNodeClient is the constructor for the NodeClient and will return an instance of
// the NodeClient struct.
func NewNodeClient(client *http.Client, address, username, password string) *NodeClient {
	return &NodeClient{
		client:   client,
		address:  address,
		username: username,
		password: password,
		usecase:  usecases.NewBitcoinClient(),
	}
}

// GetBlock calls the usecase to send a request according to the method and params,
// the returned values are then converted into a Block Model.
func (n NodeClient) GetBlock(blockHash string) (*models.Block, error) {
	nodeReq := utils.NodeRequest(n.client, n.address, n.username, n.password, "getblock", []string{blockHash})

	res, err := n.usecase.SendNodeRequest(nodeReq)
	if err != nil {
		return nil, err
	}

	block, err := utils.ConvNodeResToBlock(res)
	if err != nil {
		return nil, err
	}

	return block, nil
}