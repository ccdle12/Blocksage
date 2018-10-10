package nodeuc

import (
	"github.com/ccdle12/Blocksage/crawler/models"
)

// Usecase is the interface for the detailed implementation of the Node.
type Usecase interface {
	SendNodeRequest(*models.NodeRequest) (*models.NodeResponse, error)
}
