package dbuc

import (
	"github.com/ccdle12/Blocksage/go-crawler/models"
)

// Usecase is the usecase/implementation interface for all DB usecases.
type Usecase interface {
	OpenConnection() error
	CloseConnection() error
	InsertBlock(b *models.Block) error
	InsertTransaction(t *models.Transaction) error
}
