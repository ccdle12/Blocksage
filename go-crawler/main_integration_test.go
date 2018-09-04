// +build integration

package main

import (
	"github.com/ccdle12/Blocksage/go-crawler/controllers"
	"github.com/ccdle12/Blocksage/go-crawler/test-utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

// NOTE: Ping() in db usecases does NOT respect context timeouts, testing dbHandler.Connect()
// will take longer than the other tests.

// TestDBConnection will create a DBHandler and test we can open a connection to the DB.
func TestDBConnection(t *testing.T) {
	assert := assert.New(t)
	// TODO (ccdle12): Look into VIPER and inject environment variables into utils
	dbHandler, err := controllers.NewDBHandler(testutils.DBHost, testutils.DBPort, testutils.DBUser,
		testutils.DBPassword, testutils.DBName, testutils.DBType)

	assert.NoError(err, "Should not return errors when initializing dbHandler")

	err = dbHandler.Connect()
	assert.NoError(err, "Should not return error when connecting to the DB")

	defer dbHandler.Close()

	assert.NotNil(dbHandler, "dbHandler should not be nil")
}

// TestDBConnectionShouldFail will create a DBHandler and test we can open a connection to the DB.
func TestDBConnectionShouldFail(t *testing.T) {
	assert := assert.New(t)

	dbHandler, err := controllers.NewDBHandler(testutils.IncorrectDBHost, testutils.DBPort, testutils.DBUser,
		testutils.DBPassword, testutils.DBName, testutils.DBType)

	assert.NoError(err, "Should not return errors when initializing dbHandler")

	err = dbHandler.Connect()
	assert.Error(err, "Should return error trying to connect to a DB that doesn't exist")
}
