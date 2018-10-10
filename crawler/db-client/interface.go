package dbclient

import (
	"github.com/ccdle12/Blocksage/crawler/models"
)

// Controller is an interface for all DB Controllers.
type Controller interface {
	WriteBlock(block *models.Block) error
	WriteTransaction(tx *models.Transaction) error
	WriteInput(txHash string, in models.TransactionInput) error
	Connect() error
	Close() error
}
