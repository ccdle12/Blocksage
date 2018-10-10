package nodeclient

import (
	"github.com/ccdle12/Blocksage/go-crawler/injector"
	"github.com/ccdle12/Blocksage/go-crawler/models"
	"github.com/ccdle12/Blocksage/go-crawler/node-client-uc"
	"github.com/ccdle12/Blocksage/go-crawler/utils"
	"net/http"
)

// Client is a struct that will handle the connection to the Bitcoin Nodes.
type Client struct {
	client   *http.Client
	address  string
	username string
	password string
	usecase  nodeuc.Usecase
}

// New is the constructor for the Client and will return an instance of
// the Client struct.
func New(client *http.Client, address, username, password string) *Client {
	return &Client{
		client:   client,
		address:  utils.FormatAddress(address),
		username: username,
		password: password,
		usecase:  nodeuc.NewBitcoinClient(),
	}
}

// GetBlock calls the usecase to send a request according to the method and params,
// the returned values are then converted into a Block Model.
func (n *Client) GetBlock(blockHash string) (*models.Block, error) {
	nodeReq := injector.NodeRequest(n.client, n.address, n.username, n.password, "getblock", []string{blockHash})

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

// GetTransaction calls the usecase to send a request according to the method and params,
// the returned values are then converted into a Transaction Model.
func (n *Client) GetTransaction(txHash string) (*models.Transaction, error) {

	// Send request to get raw transaction
	nodeReq := injector.NodeRequest(n.client, n.address, n.username, n.password, "getrawtransaction", []string{txHash})
	res, err := n.usecase.SendNodeRequest(nodeReq)
	if err != nil {
		return nil, err
	}

	// Send request to decode the raw tx
	nodeReq = injector.NodeRequest(n.client, n.address, n.username, n.password, "decoderawtransaction", []string{res.Result.(string)})
	res, err = n.usecase.SendNodeRequest(nodeReq)
	if err != nil {
		return nil, err
	}

	tx, err := utils.ConvNodeResToTransaction(res)
	if err != nil {
		return nil, err
	}

	return tx, nil
}
