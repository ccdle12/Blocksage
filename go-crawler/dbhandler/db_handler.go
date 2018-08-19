package dbhandler

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/ccdle12/Blocksage/go-crawler/models"

	// Using the blank identifier in order to solely provide the side-effects of the package.
	// Essentially the side effect is calling the `init()`
	// method of `lib/pq`: func init () {  sql.Register("postgres", &Driver{} }
	"github.com/lib/pq"
)

var (
	// ErrMissingArgument is an error that catches an incorrect argument passed to a function parameter
	ErrMissingArgument = errors.New("Error: Missing Argument, make sure all arguments are passed and not empty strings")

	// ErrCreatingDBConnection is an error that catches a failed attempt to create a DB connection
	ErrCreatingDBConnection = errors.New("Error: There was an error when creating the DB connection")

	// ErrFailedToCloseDB is an error that catches a failed attempt to close a DB connection
	ErrFailedToCloseDB = errors.New("Error: There was an error when closing the DB connection")

	// ErrFailedToPingDB is an error that catches a failed attempt to ping the DB
	ErrFailedToPingDB = errors.New("Error: There was an error when pinging the DB")
)

// DBHandler is a struct that manages a connection to the DB and provides functions to read/write to the DB
type DBHandler struct {
	DB  *sql.DB
	cfg *config
}

// Config is a struct that holds the information needed to open a connection to a DB
type config struct {
	Host     string
	Port     string
	User     string
	Password string

	// DatabaseName to connect to (must have been created prior)
	DatabaseName string

	// DatabaseType is used to create the driver to a certain DB types
	DatabaseType string
}

// The following error variables are used to handle errors in Tables
var (
	ErrPassingNil  = errors.New("Error: Passing nil as argument")
	ErrCreateTable = errors.New("Error: Unable to create a table")
)

// New creates and returns an instance of the DBHandler Struct
func New(host, port, user, password, dbName, dbType string) (*DBHandler, error) {
	if host == "" || port == "" || user == "" || password == "" || dbName == "" || dbType == "" {
		return nil, ErrMissingArgument
	}

	cfg := &config{
		Host:         host,
		Port:         port,
		User:         user,
		Password:     password,
		DatabaseName: dbName,
		DatabaseType: dbType,
	}

	db, err := openDBConnection(cfg)
	if err != nil {
		return nil, err
	}

	if err := pingDBConnection(db); err != nil {
		return nil, err
	}

	d := &DBHandler{
		DB:  db,
		cfg: cfg,
	}

	d.createTables()

	return d, nil
}

func openDBConnection(cfg *config) (*sql.DB, error) {
	db, err := sql.Open(cfg.DatabaseType, fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		cfg.User, cfg.Password, cfg.DatabaseName, cfg.Host, cfg.Port))

	if err != nil {
		return nil, ErrCreatingDBConnection
	}

	return db, nil
}

func pingDBConnection(db *sql.DB) error {
	if err := db.Ping(); err != nil {
		return ErrFailedToPingDB
	}

	return nil
}

// Close performs the release of any resources that `sql/database` DB pool created. This is usually meant
// to be used in the exiting of a program or `panic`ing.
func (d *DBHandler) Close() error {
	if d.DB == nil {
		return nil
	}

	if err := d.DB.Close(); err != nil {
		return ErrFailedToCloseDB
	}

	fmt.Println("closing connection")
	return nil
}

// TODO (ccdle12): Update this function to call sub functions to create specific tables
// createTables creates all the tables in the DB
func (d *DBHandler) createTables() error {
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

	if _, err := d.DB.Exec(qry); err != nil {
		return ErrCreateTable
	}

	return nil
}

// InsertBlock creates and inserts a new row to the DB table
func (d *DBHandler) InsertBlock(receivedBlock *models.Block) error {
	// TODO (ccdle12):
	// This should be separated into different functions, WriteBlock() will handle the conditional statements
	// Move the actual DB execution to another function like InsertBlock()

	// 1. Get the last Block in the DB,
	// 2. If the lastBlockHash != receivedBlock.Hash then we've received a new block
	// 	a. Update a the last Block's 'nextblockhash' in the db with the hash of the receivedBlock.Hash
	// 	b. Write the new block into the DB
	// 3. Else it will do nothing
	fmt.Println("Insert Block called")
	lastBlockHash, lastNextBlockHash, lastBlockErr := d.GetLastBlockHashes()
	if lastBlockErr != nil {
		fmt.Println("Could not retrieve last receivedBlock")
		return lastBlockErr
	}
	if lastNextBlockHash != "" {
		fmt.Println("last receivedBlock: ", lastNextBlockHash)
	}
	fmt.Println(lastBlockHash)

	// TODO (ccdle12): checking if the received block is the next highest block
	// THIS IS BROKEN, highest block is writing the previous block as the nextblockhash
	fmt.Println("About to update next block hash")
	if lastBlockHash != receivedBlock.Hash {
		// TODO (ccdle12): Update the lastBlockHash's nextBlockHash with the receivedBlock.Hash
		fmt.Println("Updated next block hash")
		d.writeNextBlockHash(lastBlockHash, receivedBlock.Hash)
	}

	_, err := d.DB.Exec("INSERT INTO blocks (hash, confirmations, strippedsize, size, weight, height, version, versionHex, merkleroot, tx, time, mediantime, nonce, bits, difficulty, chainwork, nextblockhash) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17) ON CONFLICT (hash) DO NOTHING;",
		receivedBlock.Hash, receivedBlock.Confirmations, receivedBlock.Strippedsize, receivedBlock.Size, receivedBlock.Weight, receivedBlock.Height, receivedBlock.Version,
		receivedBlock.VersionHex, receivedBlock.MerkleRoot, pq.Array(receivedBlock.TX), receivedBlock.Time, receivedBlock.MedianTime, receivedBlock.Nonce, receivedBlock.Bits, receivedBlock.Difficulty,
		receivedBlock.Chainwork, receivedBlock.NextBlockHash)
	//TODO (ccdle12): Handle Error
	if err != nil {
		fmt.Println("Erorr when writing to DB: ", err.Error())
		return errors.New("Failed to insert row into DB")
	}

	return nil
}

// GetLastBlockHashes retrieves the block hash and the next block hash of the highest block
func (d *DBHandler) GetLastBlockHashes() (string, string, error) {
	fmt.Println("GetLastBlockHashes() called")
	rows, err := d.DB.Query("SELECT hash, nextblockhash FROM blocks ORDER BY id DESC LIMIT 1")
	if err != nil {
		//TODO (ccdle12): Handle Error
		fmt.Printf("Error from SELECT query: \n", err)
		return "", "", err
	}
	defer rows.Close()

	var hash string
	var nextBlockHash string

	for rows.Next() {
		if scanErr := rows.Scan(&hash, &nextBlockHash); scanErr != nil {
			fmt.Printf("Error scanning receivedBlock to struct: \n", scanErr)
			//TODO (ccdle12): Handle Error
			return "", "", scanErr
		}
	}

	fmt.Println("Hash of highest block: ", hash)
	fmt.Println("nextBlockHash of highest block: ", nextBlockHash)

	return hash, nextBlockHash, nil
}

func (d *DBHandler) writeNextBlockHash(lastBlockHash, receivedBlockHash string) {
	fmt.Println("Received Block Hash: ", receivedBlockHash)
	fmt.Println("Last Block Hash: ", lastBlockHash)
	_, err := d.DB.Exec("UPDATE blocks SET nextblockhash = $1 WHERE hash = $2", receivedBlockHash, lastBlockHash)
	if err != nil {
		//TODO (ccdle12): Handle Error
		fmt.Println(err)
	}
}

// type Block struct {
// 	Hash          string   `json:"hash"`
// 	Confirmations int      `json:"confirmations"`
// 	Strippedsize  int      `json:"strippedsize"`
// 	Size          int      `json:"size"`
// 	Weight        int      `json:"weight"`
// 	Height        int      `json:"height"`
// 	Version       int      `json:"version"`
// 	VersionHex    string   `json:"versionHex"`
// 	MerkleRoot    string   `json:"merkleroot"`
// 	TX            []string `json:"tx"`
// 	Time          int      `json:"time"`
// 	MedianTime    int      `json:"mediantime"`
// 	Nonce         int      `json:"nonce"`
// 	Bits          string   `json:"bits"`
// 	Difficulty    float64  `json:"difficulty"`
// 	Chainwork     string   `json:"chainwork"`
// 	NextBlockHash string   `json:"nextblockhash"`
// }

// GetBlocks returns the rows from the DB as BlockRow structs
func (d *DBHandler) GetBlocks() ([]models.Block, error) {
	// Retrieve all rows from the blocks table
	rows, err := d.DB.Query("SELECT * FROM blocks")
	if err != nil {
		fmt.Printf("Error from SELECT query: \n", err)
		return nil, err
	}
	defer rows.Close()

	// Intialise a models.Block slice ready to receive all the rows init as Blocks
	var blocks []models.Block
	for rows.Next() {
		receivedBlock := models.Block{}

		if err := rows.Scan(&receivedBlock.ID, &receivedBlock.Hash, &receivedBlock.Confirmations, &receivedBlock.Strippedsize, &receivedBlock.Size, &receivedBlock.Weight, &receivedBlock.Height, &receivedBlock.Version,
			&receivedBlock.VersionHex, &receivedBlock.MerkleRoot, &receivedBlock.TX, &receivedBlock.Time, &receivedBlock.MedianTime, &receivedBlock.Nonce, &receivedBlock.Bits, &receivedBlock.Difficulty,
			&receivedBlock.Chainwork, &receivedBlock.NextBlockHash); err != nil {
			fmt.Println("Err returned from scan: ", err)
			return nil, err
		}
		fmt.Println("Appending receviedBlock into blocks")
		blocks = append(blocks, receivedBlock)
	}

	return blocks, nil
}
