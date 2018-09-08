package usecases

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/ccdle12/Blocksage/go-crawler/models"
	"github.com/ccdle12/Blocksage/go-crawler/utils"

	// Using the blank identifier in order to solely provide the side-effects of the package.
	// Essentially the side effect is calling the `init()`
	// method of `lib/pq`: func init () {  sql.Register("postgres", &Driver{} }
	_ "github.com/lib/pq"
)

// TestPostGresClient struct is the concrete implementation of the TestPostgres
// usecase.
type TestPostGresClient struct {
	cfg *models.DBConfig
	db  *sql.DB
}

// TestNewPostGresClient will return a new instance of the TestPostGresClient struct.
// This struct is essentially a mirror of PostgresClient, this will allow the dev
// to test reading/writing to the DB without corrupting the integrity of the main
// tables.
func TestNewPostGresClient(cfg *models.DBConfig) *TestPostGresClient {
	return &TestPostGresClient{
		cfg: cfg,
	}
}

// OpenConnection will create a connection to the DB.
func (p *TestPostGresClient) OpenConnection() error {
	if err := p.connect(); err != nil {
		return err
	}

	if err := p.ping(); err != nil {
		return err
	}

	return nil
}

func (p *TestPostGresClient) connect() error {
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

func (p *TestPostGresClient) ping() error {
	if err := p.db.Ping(); err != nil {
		return errors.New("TestPostGresClient failed to ping db")
	}

	return nil
}

// CloseConnection will close the connection to the DB.
func (p *TestPostGresClient) CloseConnection() error {
	if p.db == nil {
		return nil
	}

	if err := p.db.Close(); err != nil {
		return errors.New("TestPostGresClient failed to close db connection")
	}

	return nil
}

// InsertBlock will write a Block to the DB
func (p *TestPostGresClient) InsertBlock(b *models.Block) error {
	_, err := p.db.Exec("INSERT INTO testblocks (hash, strippedsize, size, weight, height, version, versionHex, merkleroot, foreignkeytx, time, mediantime, nonce, bits, difficulty, chainwork, nextblockhash) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16) ON CONFLICT (hash) DO NOTHING;",
		b.Hash, b.Strippedsize, b.Size, b.Weight, b.Height, b.Version,
		b.VersionHex, b.MerkleRoot, b.Hash, b.Time, b.MedianTime, b.Nonce, b.Bits, b.Difficulty, b.Chainwork, b.NextBlockHash)

	if err != nil {
		return utils.FailedToInsertToDB(err)
	}

	return nil
}
