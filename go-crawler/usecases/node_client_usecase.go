package usecases

import (
	"github.com/ccdle12/Blocksage/go-crawler/models"
)

// NodeClient is the interface for the DBClient Usecase.
type NodeClient interface {
	SendNodeRequest(*models.NodeRequest) (*models.NodeResponse, error)
}
