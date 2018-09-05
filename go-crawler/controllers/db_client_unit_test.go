// +build unit

package controllers

import (
	"github.com/ccdle12/Blocksage/go-crawler/test-utils"
	"github.com/ccdle12/Blocksage/go-crawler/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestInitDBClient will test that a DBClient can be initialized.
func TestInitDBClient(t *testing.T) {
	assert := assert.New(t)

	dbClient, err := NewDBClient(testutils.DBHost, testutils.DBPort, testutils.DBUser,
		testutils.DBPassword, testutils.DBName, testutils.DBType)

	assert.NoError(err, "There should be no error")
	assert.NotNil(dbClient, "DBClient is not nil when attempting to initialize")
}

// TestConfigInit will test that the cfg DBConfig struct was initialized by the
// constructor.
func TestConfigInit(t *testing.T) {
	assert := assert.New(t)

	dbClient, err := NewDBClient(testutils.DBHost, testutils.DBPort, testutils.DBUser,
		testutils.DBPassword, testutils.DBName, testutils.DBType)

	assert.NoError(err, "There should be no error")
	assert.NotNil(dbClient.cfg, "cfg is not nil when attempting to initialize")
	assert.EqualValues(dbClient.cfg.DBHost, testutils.DBHost, "Same host was initialized in cfg.DBHost")
}

// TestEmptyStringInit will test that we are passing an empty string as one of the parameters for
// initializing a DBClient.
func TestEmptyStringInit(t *testing.T) {
	assert := assert.New(t)

	dbClient, err := NewDBClient(testutils.DBHost, "", testutils.DBUser,
		testutils.DBPassword, testutils.DBName, testutils.DBType)

	assert.Error(err, "Error should be returned")
	assert.EqualValues(err, utils.ErrPassingEmptyString, "The error return ErrPassingEmptyString")
	assert.Nil(dbClient, "DBClient is not nil when attempting to initialize")
}

// TestUseCaseInit will test that the DB usecase was initialized.
func TestUseCaseInit(t *testing.T) {
	assert := assert.New(t)

	dbClient, err := NewDBClient(testutils.DBHost, testutils.DBPort, testutils.DBUser,
		testutils.DBPassword, testutils.DBName, testutils.DBType)

	assert.NoError(err, "No error should be returned")
	assert.NotNil(dbClient.usecase, "usecase is not nil and initialized")
}

// TODO (ccdle12): For Integration Tests
// TestDBClientConnection
// func TestDBClientConnection(t *testing.T) {
// 	assert := assert.New(t)

// 	dbClient, err := NewDBClient(testutils.DBHost, testutils.DBPort, testutils.DBUser,
// 		testutils.DBPassword, testutils.DBName, testutils.DBType)

// 	assert.NoError(err, "No error should be returned")
// 	assert.NotNil(dbClient.usecase, "usecase is not nil and initialized")

// 	dbClient.Connect()
// }

// TODO (ccdle12): For Integration Tests
// TestDBClientClose
// func TestDBClientClose(t *testing.T) {
// 	assert := assert.New(t)

// 	dbClient, err := NewDBClient(testutils.DBHost, testutils.DBPort, testutils.DBUser,
// 		testutils.DBPassword, testutils.DBName, testutils.DBType)

// 	assert.NoError(err, "No error should be returned")
// 	assert.NotNil(dbClient.usecase, "usecase is not nil and initialized")

// 	dbClient.Close()
// }
