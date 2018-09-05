// +build integration

package main

import (
	"github.com/ccdle12/Blocksage/go-crawler/controllers"
	"github.com/ccdle12/Blocksage/go-crawler/test-utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

// NOTE: Ping() in db usecases does NOT respect context timeouts, testing dbClient.Connect()
// will take longer than the other tests.

// TestDBConnection will create a DBClientUsecase and test we can open a connection to the DB.
func TestDBConnection(t *testing.T) {
	assert := assert.New(t)
	// TODO (ccdle12): Look into VIPER and inject environment variables into utils
	dbClient, err := controllers.NewDBClient(testutils.DBHost, testutils.DBPort, testutils.DBUser,
		testutils.DBPassword, testutils.DBName, testutils.DBType)

	assert.NoError(err, "Should not return errors when initializing dbClient")

	err = dbClient.Connect()
	assert.NoError(err, "Should not return error when connecting to the DB")

	defer dbClient.Close()

	assert.NotNil(dbClient, "dbClient should not be nil")
}

// TestDBConnectionShouldFail will create a DBClientUsecase and test we can open a connection to the DB.
func TestDBConnectionShouldFail(t *testing.T) {
	assert := assert.New(t)

	dbClient, err := controllers.NewDBClient(testutils.IncorrectDBHost, testutils.DBPort, testutils.DBUser,
		testutils.DBPassword, testutils.DBName, testutils.DBType)

	assert.NoError(err, "Should not return errors when initializing dbClient")

	err = dbClient.Connect()
	assert.Error(err, "Should return error trying to connect to a DB that doesn't exist")
}
