package dbuc

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/ccdle12/Blocksage/crawler/models"
	"github.com/ccdle12/Blocksage/crawler/utils"

	// Using the blank identifier in order to solely provide the side-effects of the package.
	// Essentially the side effect is calling the `init()`
	// method of `lib/pq`: func init () {  sql.Register("postgres", &Driver{} }
	_ "github.com/lib/pq"
)

// PostGresClient struct is the concrete implementation of the Postgres
// usecase.
type PostGresClient struct {
	cfg *models.DBConfig
	db  *sql.DB
}

// NewPostGresClient will return a new instance of the PostGresClient struct.
func NewPostGresClient(cfg *models.DBConfig) *PostGresClient {
	return &PostGresClient{
		cfg: cfg,
	}
}

// OpenConnection will create a connection to the DB.
func (p *PostGresClient) OpenConnection() error {
	if err := p.connect(); err != nil {
		return err
	}

	if err := p.ping(); err != nil {
		return err
	}

	return nil
}

func (p *PostGresClient) connect() error {
	db, err := sql.Open(p.cfg.DBType, fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		p.cfg.DBUser, p.cfg.DBPassword, p.cfg.DBName, p.cfg.DBHost, p.cfg.DBPort))

	if err != nil {
		return utils.FailedDBConnection(err)
	}

	// assign self the DB
	p.db = db

	return nil
}

func (p *PostGresClient) ping() error {
	if err := p.db.Ping(); err != nil {
		return utils.ErrFailedToPingDB
	}

	return nil
}

// CloseConnection will close the connection to the DB.
func (p *PostGresClient) CloseConnection() error {
	if p.db == nil {
		return nil
	}

	if err := p.db.Close(); err != nil {
		return utils.ErrFailedToCloseDBConnection
	}

	return nil
}

// InsertBlock will write a Block to the DB
func (p *PostGresClient) InsertBlock(b *models.Block) error {
	// Check that the struct passed is not nil.
	if b == nil {
		return errors.New("Block is Null in DB Client usecase")
	}

	// TODO: (ccdle12) having trouble using Sprintf to create queries, need to make this more elegant
	var err error
	if p.cfg.Test {
		_, err = p.db.Exec("INSERT INTO testblocks (hash, strippedsize, size, weight, height, version, versionHex, merkleroot, time, mediantime, nonce, bits, difficulty, chainwork, nextblockhash) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15) ON CONFLICT (hash) DO NOTHING;",
			b.Hash, b.Strippedsize, b.Size, b.Weight, b.Height, b.Version, b.VersionHex, b.MerkleRoot, b.Time, b.MedianTime, b.Nonce,
			b.Bits, b.Difficulty, b.Chainwork, b.NextBlockHash)
	} else {
		_, err = p.db.Exec("INSERT INTO blocks (hash, strippedsize, size, weight, height, version, versionHex, merkleroot, time, mediantime, nonce, bits, difficulty, chainwork, nextblockhash) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15) ON CONFLICT (hash) DO NOTHING;",
			b.Hash, b.Strippedsize, b.Size, b.Weight, b.Height, b.Version, b.VersionHex, b.MerkleRoot, b.Time, b.MedianTime, b.Nonce,
			b.Bits, b.Difficulty, b.Chainwork, b.NextBlockHash)
	}
	if err != nil {
		return utils.FailedToInsertToDB(err)
	}

	return nil
}

// InsertTransaction will write a Transaction to the DB
func (p *PostGresClient) InsertTransaction(t *models.Transaction) error {
	// Check that the struct passed is not nil.
	if t == nil {
		return errors.New("Transaction is Null in DB Client usecase")
	}

	// TODO: (ccdle12) having trouble using Sprintf to create queries, need to make this more elegant
	var err error
	if p.cfg.Test {
		_, err = p.db.Exec("INSERT INTO testtransactions (blockhash, txid, hash, version, size, vsize, locktime) VALUES ($1, $2, $3, $4, $5, $6, $7) ON CONFLICT (hash) DO NOTHING;",
			t.Blockhash, t.TXID, t.Hash, t.Version, t.Size, t.Vsize, t.Locktime)
	} else {
		_, err = p.db.Exec("INSERT INTO transactions (blockhash, txid, hash, version, size, vsize, locktime) VALUES ($1, $2, $3, $4, $5, $6, $7) ON CONFLICT (hash) DO NOTHING;",
			t.Blockhash, t.TXID, t.Hash, t.Version, t.Size, t.Vsize, t.Locktime)
	}
	if err != nil {
		return utils.FailedToInsertToDB(err)
	}

	return nil
}
