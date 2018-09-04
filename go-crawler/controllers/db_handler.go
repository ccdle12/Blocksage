package controllers

import (
	"github.com/ccdle12/Blocksage/go-crawler/models"
	"github.com/ccdle12/Blocksage/go-crawler/usecases"
	"github.com/ccdle12/Blocksage/go-crawler/utils"
)

// DBHandler is a struct that manages a connection to the DB and provides functions to read/write to the DB
type DBHandler struct {
	cfg     *models.DBConfig
	usecase usecases.DBUsecase
}

// NewDBHandler will return an instance of the DBHandler that will read/write from different
func NewDBHandler(host, port, user, password, dbName, dbType string) (*DBHandler, error) {
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

	usecase := usecases.NewPostGresHandler(cfg)

	return &DBHandler{
		cfg:     cfg,
		usecase: usecase,
	}, nil
}

// TODO (ccdle12): Needs Integration Testing
// Connect will request the usecase to open a connection to the DB.
func (d *DBHandler) Connect() error {
	if err := d.usecase.OpenConnection(); err != nil {
		return err
	}

	return nil
}

// TODO (ccdle12): Needs Integration Testing
// Close will request the usecase to close the connection to the DB.
func (d *DBHandler) Close() error {
	if err := d.usecase.CloseConnection(); err != nil {
		return err
	}

	return nil
}
