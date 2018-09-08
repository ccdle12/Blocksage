package controllers

import (
	"github.com/ccdle12/Blocksage/go-crawler/models"
)

// DBClientController is an interface for DBClient.
type DBClientController interface {
	WriteBlock(block *models.Block) error
}
