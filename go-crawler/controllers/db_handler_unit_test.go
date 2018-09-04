// +build unit

package controllers

import (
	"github.com/ccdle12/Blocksage/go-crawler/test-utils"
	"github.com/ccdle12/Blocksage/go-crawler/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestInitDBHandler will test that a DBHandler can be initialized.
func TestInitDBHandler(t *testing.T) {
	assert := assert.New(t)

	dbHandler, err := NewDBHandler(testutils.DBHost, testutils.DBPort, testutils.DBUser,
		testutils.DBPassword, testutils.DBName, testutils.DBType)

	assert.NoError(err, "There should be no error")
	assert.NotNil(dbHandler, "DBHandler is not nil when attempting to initialize")
}

// TestConfigInit will test that the cfg DBConfig struct was initialized by the
// constructor.
func TestConfigInit(t *testing.T) {
	assert := assert.New(t)

	dbHandler, err := NewDBHandler(testutils.DBHost, testutils.DBPort, testutils.DBUser,
		testutils.DBPassword, testutils.DBName, testutils.DBType)

	assert.NoError(err, "There should be no error")
	assert.NotNil(dbHandler.cfg, "cfg is not nil when attempting to initialize")
	assert.EqualValues(dbHandler.cfg.DBHost, testutils.DBHost, "Same host was initialized in cfg.DBHost")
}

// TestEmptyStringInit will test that we are passing an empty string as one of the parameters for
// initializing a DBHandler.
func TestEmptyStringInit(t *testing.T) {
	assert := assert.New(t)

	dbHandler, err := NewDBHandler(testutils.DBHost, "", testutils.DBUser,
		testutils.DBPassword, testutils.DBName, testutils.DBType)

	assert.Error(err, "Error should be returned")
	assert.EqualValues(err, utils.ErrPassingEmptyString, "The error return ErrPassingEmptyString")
	assert.Nil(dbHandler, "DBHandler is not nil when attempting to initialize")
}

// TestUseCaseInit will test that the DB usecase was initialized.
func TestUseCaseInit(t *testing.T) {
	assert := assert.New(t)

	dbHandler, err := NewDBHandler(testutils.DBHost, testutils.DBPort, testutils.DBUser,
		testutils.DBPassword, testutils.DBName, testutils.DBType)

	assert.NoError(err, "No error should be returned")
	assert.NotNil(dbHandler.usecase, "usecase is not nil and initialized")
}

// TODO (ccdle12): For Integration Tests
// TestDBHandlerConnection
// func TestDBHandlerConnection(t *testing.T) {
// 	assert := assert.New(t)

// 	dbHandler, err := NewDBHandler(testutils.DBHost, testutils.DBPort, testutils.DBUser,
// 		testutils.DBPassword, testutils.DBName, testutils.DBType)

// 	assert.NoError(err, "No error should be returned")
// 	assert.NotNil(dbHandler.usecase, "usecase is not nil and initialized")

// 	dbHandler.Connect()
// }

// TODO (ccdle12): For Integration Tests
// TestDBHandlerClose
// func TestDBHandlerClose(t *testing.T) {
// 	assert := assert.New(t)

// 	dbHandler, err := NewDBHandler(testutils.DBHost, testutils.DBPort, testutils.DBUser,
// 		testutils.DBPassword, testutils.DBName, testutils.DBType)

// 	assert.NoError(err, "No error should be returned")
// 	assert.NotNil(dbHandler.usecase, "usecase is not nil and initialized")

// 	dbHandler.Close()
// }
