package blocktable

import (
	"errors"
	"fmt"

	"github.com/ccdle12/Blocksage/go-crawler/dbhandler"
	"github.com/ccdle12/Blocksage/go-crawler/models"

	// Using the blank identifier in order to solely provide the side-effects of the package.
	// Essentially the side effect is calling the `init()`
	// method of `lib/pq`: func init () {  sql.Register("postgres", &Driver{} }
	"github.com/lib/pq"
)

var (
	// ErrPassingNil is an error that catches passing an arugment as nil
	ErrPassingNil = errors.New("Error: Passing nil as argument")

	// ErrCreateTable is an error that catches an error when creating a Table
	ErrCreateTable = errors.New("Error: Unable to create a table")
)

// BlockTable is a struct that holds a DBHandler and specifies the behavior when reading/writing tables for Blocks
type BlockTable struct {
	DBHandler *dbhandler.DBHandler
}

// New creates and returns an instance of the BlockTable struct
func New(dbHandler *dbhandler.DBHandler) (*BlockTable, error) {
	if dbHandler == nil {
		return nil, ErrPassingNil
	}

	b := &BlockTable{
		DBHandler: dbHandler,
	}

	return b, nil
}

// CreateTable creates a table via the DBHandler
func (b *BlockTable) CreateTable() error {
	const qry = `
		CREATE TABLE IF NOT EXISTS blocks (
			id serial PRIMARY KEY,
			hash text UNIQUE NOT NULL,
			confirmations bigint NOT NULL,
			strippedsize bigint NOT NULL,
			size bigint NOT NULL,
			weight bigint NOT NULL,
			height bigint NOT NULL,
			version bigint NOT NULL,
			versionHex text NOT NULL,
			merkleroot text UNIQUE NOT NULL,
			tx text[] NOT NULL,
			time bigint NOT NULL,
			mediantime bigint NOT NULL,
			nonce bigint NOT NULL,
			bits text NOT NULL,
			difficulty double precision NOT NULL,
			chainwork text NOT NULL,
			nextblockhash text
		)
	`

	if _, err := b.DBHandler.DB.Exec(qry); err != nil {
		return ErrCreateTable
	}

	return nil
}

// InsertBlock creates and inserts a new row to the DB table
func (b *BlockTable) InsertBlock(receivedBlock *models.Block) error {
	// TODO (ccdle12):
	// This should be separated into different functions, WriteBlock() will handle the conditional statements
	// Move the actual DB execution to another function like InsertBlock()

	// 1. Get the last Block in the DB,
	// 2. If the lastBlockHash != receivedBlock.Hash then we've received a new block
	// 	a. Update a the last Block's 'nextblockhash' in the db with the hash of the receivedBlock.Hash
	// 	b. Write the new block into the DB
	// 3. Else it will do nothing
	fmt.Println("Insert Block called")
	lastBlockHash, nextBlockHash, lastBlockErr := b.GetLastBlock()
	if lastBlockErr != nil {
		fmt.Println("Could not retrieve last receivedBlock")
		return lastBlockErr
	}
	if nextBlockHash != "" {
		fmt.Println("last receivedBlock: ", nextBlockHash)
	}
	fmt.Println(lastBlockHash)

	_, err := b.DBHandler.DB.Exec("INSERT INTO blocks (hash, confirmations, strippedsize, size, weight, height, version, versionHex, merkleroot, tx, time, mediantime, nonce, bits, difficulty, chainwork, nextblockhash) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17) ON CONFLICT (hash) DO NOTHING;",
		receivedBlock.Hash, receivedBlock.Confirmations, receivedBlock.Strippedsize, receivedBlock.Size, receivedBlock.Weight, receivedBlock.Height, receivedBlock.Version,
		receivedBlock.VersionHex, receivedBlock.MerkleRoot, pq.Array(receivedBlock.TX), receivedBlock.Time, receivedBlock.MedianTime, receivedBlock.Nonce, receivedBlock.Bits, receivedBlock.Difficulty,
		receivedBlock.Chainwork, receivedBlock.NextBlockHash)
	if err != nil {
		fmt.Println("Erorr when writing to DB: ", err.Error())
		return errors.New("Failed to insert row into DB")
	}

	return nil
}

// GetLastBlock retrieves the highest block in the DB
func (b *BlockTable) GetLastBlock() (string, string, error) {
	rows, err := b.DBHandler.DB.Query("SELECT hash, nextblockhash FROM blocks WHERE nextblockhash is NULL")
	if err != nil {
		fmt.Printf("Error from SELECT query: \n", err)
		return "", "", err
	}
	defer rows.Close()

	var hash string
	var nextBlockHash string

	for rows.Next() {
		if scanErr := rows.Scan(&hash, &nextBlockHash); scanErr != nil {
			fmt.Printf("Error scanning receivedBlock to struct: \n", scanErr)
			return "", "", scanErr
		}
	}

	return hash, nextBlockHash, nil
}

// GetBlocks returns the rows from the DB as BlockRow structs
// func (b *BlockTable) GetBlocks() ([]BlockRow, error) {
// 	rows, err := b.DBHandler.DB.Query("SELECT * FROM blocks")
// 	if err != nil {
// 		fmt.Printf("Error from SELECT query: \n", err)
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var blocks []BlockRow
// 	for rows.Next() {
// 		receivedBlock := BlockRow{}

// 		if err := rows.Scan(&receivedBlock.ID, &receivedBlock.Type); err != nil {
// 			fmt.Println("Err returned from scan")
// 			return nil, err
// 		}

// 		blocks = append(blocks, receivedBlock)
// 	}

// 	return blocks, nil
// }
