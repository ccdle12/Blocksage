// +build unit

package dbhandler

import (
	"testing"
)

func TestIncorrectArguments(t *testing.T) {
	// Should Fail due to passing empty string
	_, expectedError := New("172.25.0.4", "", "postgres-dev", "12345", "dev", "postgres")

	if expectedError == nil {
		t.Errorf("Expected an Error, this should NOT be nil")
	}

	if expectedError != ErrMissingArgument {
		t.Errorf("expectedError should be type ErrMethodRequestNotFound")
	}
}
func TestDBHandlerFailedDBType(t *testing.T) {
	// Should fail due to passing an unsupported DB type
	_, expectedError := New("172.25.0.4", "5432", "postgres-dev", "12345", "dev", "mongodb")

	if expectedError != ErrCreatingDBConnection {
		t.Errorf("DB was initalised with an unsupported DatabaseType and should have returned ErrCreatingDBConnection")
	}
}

func TestOpenDB(t *testing.T) {
	cfg := &Config{
		Host:         "172.25.0.4",
		Port:         "5432",
		User:         "postgres-dev",
		Password:     "12345",
		DatabaseName: "dev",
		DatabaseType: "postgres",
	}

	db, _ := openDBConnection(cfg)
	if db == nil {
		t.Errorf("db should have been initalised")
	}
}

func TestErrorCreatingDBConnection(t *testing.T) {
	// Should fail due to passing an unsupported DB type
	cfg := &Config{
		Host:         "172.25.0.4",
		Port:         "5432",
		User:         "postgres-dev",
		Password:     "12345",
		DatabaseName: "dev",
		DatabaseType: "mongodb",
	}

	_, err := openDBConnection(cfg)
	if err != ErrCreatingDBConnection {
		t.Errorf("DB was initalised with an unsupported DatabaseType and should have returned ErrCreatingDBConnection")
	}
}

func TestCloseDBConnection(t *testing.T) {
	// Should fail due to passing an unsupported DB type
	dbHandler, _ := New("172.25.0.4", "5432", "postgres-dev", "12345", "dev", "postgres")

	err := dbHandler.Close()

	if err != nil && dbHandler.Db != nil {
		t.Errorf("DB should be nil since the connection was closed")
	}
}
