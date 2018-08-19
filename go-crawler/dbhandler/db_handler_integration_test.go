// +build integration

package dbhandler

import (
	"testing"
)

// TODO (ccdle12): Change all tests to t.Fatalf(...)
func TestPingDBConnection(t *testing.T) {
	dbHandler, err := New("172.25.0.4", "5432", "postgres-dev", "12345", "dev", "postgres")
	defer dbHandler.Close()
	if err != nil {
		t.Fatalf("Unable to create DB Connection")
	}

	if err == ErrFailedToPingDB {
		t.Fatalf("Unable to ping connection to DB")
	}
}

func TestOpenCloseDBConnection(t *testing.T) {
	dbHandler, err := New("172.25.0.4", "5432", "postgres-dev", "12345", "dev", "postgres")
	if err == ErrCreatingDBConnection {
		t.Fatalf("Unable to create DB Connection")
	}

	if err == ErrFailedToPingDB {
		t.Fatalf("Unable to ping connection to DB")
	}

	err = dbHandler.Close()
	if err != nil {
		t.Fatalf("DB connection was not correctly closed")
	}

	err = pingDBConnection(dbHandler.DB)
	if err != ErrFailedToPingDB {
		t.Fatalf("Ping should have failed because the connection is closed")
	}
}

func TestCreateTable(t *testing.T) {
	// Should create a table in the DB
	dbHandler, err := New("172.25.0.4", "5432", "postgres-dev", "12345", "dev", "postgres")
	defer dbHandler.Close()
	if err != nil {
		t.Fatalf("Unable to create DB Connection")
	}

	_, err = dbHandler.DB.Query("SELECT id FROM blocks ORDER BY id DESC LIMIT 1")
	if err != nil {
		t.Fatalf("Not able to read from the blocks table, table does not exist")
	}
}

// TODO (ccdle12): Update this test to read the genesis block
func TestGetBlocksFromTable(t *testing.T) {
	// TODO: Add a test to catch an unrecognised host
	dbHandler, err := New("172.25.0.4", "5432", "postgres-dev", "12345", "dev", "postgres")
	defer dbHandler.Close()
	if err != nil {
		t.Fatalf("Creating the dbHandler should NOT have resulted in an error")
	}

	blocks, err := dbHandler.GetBlocks()
	if err != nil {
		t.Fatalf("Failed to read blocks from DB table")
	}

	t.Log("Blocks returned: ", blocks)
	if len(blocks) == 0 {
		t.Fatalf("Failed to read any blocks")
	}
}

// TODO (ccdle12): Add test to get the highest 10 ids,
// // sort them descending and each one should successfully link through nextblockhashes
// func TestBlocksAreLinked(t *testing.T) {
// 	//1 . Create a dbHandler
// 	dbHandler, err := dbhandler.New("172.25.0.4", "5432", "postgres-dev", "12345", "dev", "postgres")
// 	defer dbHandler.Close()
// 	if err != nil {
// 		t.Fatalf("Creating the dbHandler should NOT have resulted in an error")
// 	}
// }
