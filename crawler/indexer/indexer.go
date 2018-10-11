package indexer

import (
	"github.com/ccdle12/Blocksage/crawler/db-client"
	"github.com/ccdle12/Blocksage/crawler/models"
	"github.com/ccdle12/Blocksage/crawler/node-client"
)

// Indexer is the struct that will hold the logic for the crawling process.
// It is responsible for making calls via the client controller to the node
// and writing the returned results to the db via the db controller.
type Indexer struct {
	node nodeclient.Controller
	db   dbclient.Controller
}

// New is the constructor for the Indexer Struct.
func New(node nodeclient.Controller, db dbclient.Controller) *Indexer {
	return &Indexer{
		node: node,
		db:   db,
	}
}

// GetBlock calls the Node Client to retrieve a Block according to a blockhash.
func (i *Indexer) GetBlock(hash string) (*models.Block, error) {
	// Get a block.
	block, err := i.node.GetBlock(hash)
	if err != nil {
		return nil, err
	}

	return block, nil
}

// Write will receive a Block and write all subsequent information from the block
// to the db, including transactions, inputs and outputs.
func (i *Indexer) Write(block *models.Block) error {
	// Write the received block to the db.
	if err := i.db.WriteBlock(block); err != nil {
		return err
	}

	// Write all transactions from the block to the db.
	for _, hash := range block.TX {
		// Call the node to get a transaction details from the has.
		tx, err := i.node.GetTransaction(hash)
		if err != nil {
			return err
		}

		// Add blockhash to transaction, this is used as the Foreign Key in
		// the transaction table.
		tx.Blockhash = block.Hash

		// Write the retrieved transaction to the transaction table.
		if err = i.db.WriteTransaction(tx); err != nil {
			return err
		}

		// Loop over each input in the transaction.
		for _, input := range tx.Vin {
			// Write the transaction inputs to the inputs table.
			if err = i.db.WriteInput(tx.Hash, input); err != nil {
				return err
			}
		}

		// Loop over each output in the transaction.
		for _, output := range tx.Vout {
			// Write the transaction outputs to the outputs table.
			if err = i.db.WriteOutput(tx.Hash, output); err != nil {
				return err
			}
		}
	}

	return nil
}
