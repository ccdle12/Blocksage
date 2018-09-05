package controllers

import (
	"github.com/ccdle12/Blocksage/go-crawler/models"
	"github.com/ccdle12/Blocksage/go-crawler/usecases"
	"github.com/ccdle12/Blocksage/go-crawler/utils"
)

// DBClient is a struct that manages a connection to the DB and provides functions to read/write to the DB
type DBClient struct {
	cfg     *models.DBConfig
	usecase usecases.DBClient
}

// NewDBClient will return an instance of the DBClient that will read/write from different
// usecases.
func NewDBClient(host, port, user, password, dbName, dbType string) (*DBClient, error) {
	if utils.EmptyString(host, port, user, password, dbName) {
		return nil, utils.ErrPassingEmptyString
	}

	cfg := &models.DBConfig{
		DBHost:     host,
		DBPort:     port,
		DBUser:     user,
		DBPassword: password,
		DBName:     dbName,
		DBType:     dbType,
	}

	usecase := usecases.NewPostGresClient(cfg)

	return &DBClient{
		cfg:     cfg,
		usecase: usecase,
	}, nil
}

// TODO (ccdle12): Needs Integration Testing
// Connect will request the usecase to open a connection to the DB.
func (d *DBClient) Connect() error {
	if err := d.usecase.OpenConnection(); err != nil {
		return err
	}

	return nil
}

// TODO (ccdle12): Needs Integration Testing
// Close will request the usecase to close the connection to the DB.
func (d *DBClient) Close() error {
	if err := d.usecase.CloseConnection(); err != nil {
		return err
	}

	return nil
}
