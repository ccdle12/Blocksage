// +build integration

package dbhandler

import (
	"testing"
)

func TestPingDBConnection(t *testing.T) {
	// TODO: Initalise a config in argument?
	dbHandler, err := New("172.25.0.4", "5432", "postgres-dev", "12345", "dev", "postgres")
	defer dbHandler.Close()
	if err != nil {
		t.Errorf("Fail, was unable to create DB Connection")
	}

	if err == ErrFailedToPingDB {
		t.Errorf("Fail unable to ping connection to DB")
	}
}

func TestOpenCloseDBConnection(t *testing.T) {
	// TODO: Initalise a config in argument?
	dbHandler, err := New("172.25.0.4", "5432", "postgres-dev", "12345", "dev", "postgres")
	if err == ErrCreatingDBConnection {
		t.Errorf("Fail, was unable to create DB Connection")
	}

	if err == ErrFailedToPingDB {
		t.Errorf("Fail unable to ping connection to DB")
	}

	err = dbHandler.Close()
	if err != nil {
		t.Errorf("DB connection was not correctly closed")
	}

	err = pingDBConnection(dbHandler.DB)
	if err != ErrFailedToPingDB {
		t.Errorf("Ping should have failed because the connection is closed")
	}
}
