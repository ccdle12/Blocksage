// +build integration

package dbhandler

import (
	"testing"
)

func TestPingDBConnection(t *testing.T) {
	cfg := &Config{
		Host:         "172.25.0.4",
		Port:         "5432",
		User:         "postgres-dev",
		Password:     "12345",
		DatabaseName: "dev",
		DatabaseType: "postgres",
	}

	db, _ := openDBConnection(cfg)
	err := pingDBConnection(db)
	if err != nil {
		t.Errorf("Connection failure, unable to ping the DB")
	}
}
