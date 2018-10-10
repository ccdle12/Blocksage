package nodeclient

import (
	"github.com/ccdle12/Blocksage/go-crawler/models"
)

// Controller is an interface for NodeClients, to interact with
// a Blockchain Node.
type Controller interface {
	GetBlock(blockHash string) (*models.Block, error)
	GetTransaction(txHash string) (*models.Transaction, error)
}
