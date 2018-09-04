package controllers

import (
	"github.com/ccdle12/Blocksage/go-crawler/models"
)

// NodeClientController is an interface for NodeClients.
type NodeClientController interface {
	GetBlock(blockHash string) (*models.Block, error)
}
