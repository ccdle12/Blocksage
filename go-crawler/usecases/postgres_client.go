package usecases

import (
	"database/sql"
	"fmt"
	"github.com/ccdle12/Blocksage/go-crawler/models"
	"github.com/ccdle12/Blocksage/go-crawler/utils"

	// Using the blank identifier in order to solely provide the side-effects of the package.
	// Essentially the side effect is calling the `init()`
	// method of `lib/pq`: func init () {  sql.Register("postgres", &Driver{} }
	"github.com/lib/pq"
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

// TODO (ccdle12): This will need it's own integration tests
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
		return utils.ErrFailedToOpenDBConnection
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

// TODO (ccdle12): This will need it's own integration tests
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

// TODO (ccdle12): Fix InsertBlock,
// 1. make height the primary key
// 2. remove confirmations
// 3. create a table for tx and write txs to them
// InsertBlock will write a Block to the DB
func (p *PostGresClient) InsertBlock(b *models.Block) error {
	_, err := p.db.Exec("INSERT INTO blocks (hash, strippedsize, size, weight, height, version, versionHex, merkleroot, tx, time, mediantime, nonce, bits, difficulty, chainwork, nextblockhash) VALUES ($1, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17) ON CONFLICT (hash) DO NOTHING;",
		b.Hash, b.Strippedsize, b.Size, b.Weight, b.Height, b.Version,
		b.VersionHex, b.MerkleRoot, pq.Array(b.TX), b.Time, b.MedianTime, b.Nonce, b.Bits, b.Difficulty, b.Chainwork, b.NextBlockHash)

	if err != nil {
		return utils.ErrFailedToInsertToDB
	}

	return nil
}
