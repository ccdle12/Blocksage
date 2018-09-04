package usecases

import (
	"github.com/ccdle12/Blocksage/go-crawler/models"
)

// NodeClientUsecase is the interface for the BitcoinClient Usecase.
type NodeClientUsecase interface {
	SendNodeRequest(*models.NodeRequest) (*models.NodeResponse, error)
}
