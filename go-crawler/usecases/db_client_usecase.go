package usecases

import (
	"github.com/ccdle12/Blocksage/go-crawler/models"
)

// DBClient is the usecase/implementation interface for all DB usecases.
type DBClient interface {
	OpenConnection() error
	CloseConnection() error
	InsertBlock(b *models.Block) error
}
