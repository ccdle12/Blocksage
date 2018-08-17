// +build unit

package blocktable

import (
	"testing"

	"github.com/ccdle12/Blocksage/go-crawler/dbhandler"
)

func TestCreateBlockTableStruct(t *testing.T) {
	dbHandler, err := dbhandler.New("172.25.0.4", "5432", "postgres-dev", "12345", "dev", "postgres")
	if err != nil {
		t.Errorf("Creating the dbHandler should NOT have resutled in an error")
	}

	blockTable, err := New(dbHandler)

	if err != nil {
		t.Errorf("Block table should have been initalised correctly and NOT returned an error")
	}

	if blockTable == nil {
		t.Errorf("Block Table should have been initialised correctly and NOT be nil")
	}
}

func TestErrorPassingNil(t *testing.T) {
	blockTable, err := New(nil)

	if err != ErrPassingNil {
		t.Errorf("ErrPassingNil should have been thrown since we are passing a zero value of DBHandler")
	}

	if blockTable != nil {
		t.Errorf("Block Table should have been nil since we failed to pass a valid DBHandler")
	}

}
