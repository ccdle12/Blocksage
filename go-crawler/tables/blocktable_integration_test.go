// +build integration

package blocktable

import (
	"testing"

	"github.com/ccdle12/Blocksage/go-crawler/dbhandler"
)

func TestCreateTable(t *testing.T) {
	// Should create a table in the DB
	dbHandler, err := dbhandler.New("172.25.0.4", "5432", "postgres-dev", "12345", "dev", "postgres")
	defer dbHandler.Close()
	if err != nil {
		t.Errorf("Creating the dbHandler should NOT have resulted in an error")
	}

	blockTable, err := New(dbHandler)
	if err != nil {
		t.Errorf("Block table should have been initalised correctly and NOT returned an error")
	}

	err = blockTable.CreateTable()
	if err != nil {
		t.Errorf("Table should have been initialised correctly and err should NOT be nil")
	}
}

func TestInsertRowToTable(t *testing.T) {
	dbHandler, err := dbhandler.New("172.25.0.4", "5432", "postgres-dev", "12345", "dev", "postgres")
	defer dbHandler.Close()
	if err != nil {
		t.Errorf("Creating the dbHandler should NOT have resulted in an error")
	}

	blockTable, err := New(dbHandler)
	if err != nil {
		t.Errorf("Block table should have been initalised correctly and NOT returned an error")
	}

	err = blockTable.CreateTable()
	if err != nil {
		t.Errorf("Table should have been initialised correctly and err should NOT be nil")
	}

	blockRow := BlockRow{
		ID:   1,
		Type: "block",
	}

	err = blockTable.InsertBlock(blockRow)
	if err != nil {
		t.Errorf("Block table should inserted a new block to the table")
	}
}
func TestReadRowFromTable(t *testing.T) {
	// TODO: Add a test to catch an unrecognised host
	dbHandler, err := dbhandler.New("172.25.0.4", "5432", "postgres-dev", "12345", "dev", "postgres")
	defer dbHandler.Close()
	if err != nil {
		t.Errorf("Creating the dbHandler should NOT have resulted in an error")
	}

	blockTable, err := New(dbHandler)
	if err != nil {
		t.Errorf("Block table should have been initalised correctly and NOT returned an error")
	}

	err = blockTable.CreateTable()
	if err != nil {
		t.Errorf("Table should have been initialised correctly and err should NOT be nil")
	}

	blockRow := BlockRow{
		ID:   1,
		Type: "new row",
	}

	err = blockTable.InsertBlock(blockRow)
	if err != nil {
		t.Errorf("Block table should insert a new row to the table")
	}

	blocks, err := blockTable.GetBlocks()
	if err != nil {
		t.Errorf("Failed to read blocks from DB table")
	}

	t.Log("Blocks returned: ", blocks)
	if len(blocks) == 0 {
		t.Errorf("Failed to write and then read any blocks")
	}
}
