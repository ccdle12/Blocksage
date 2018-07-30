package dbhandler

import (
	"database/sql"
	"errors"
	"fmt"
	// Using the blank identifier in order to solely provide the side-effects of the package.
	// Essentially the side effect is calling the `init()`
	// method of `lib/pq`: func init () {  sql.Register("postgres", &Driver{} }
	_ "github.com/lib/pq"
)

var (
	// ErrMissingArgument is an error that catches an incorrect argument passed to a function parameter
	ErrMissingArgument = errors.New("Error: Missing Argument, make sure all arguments are passed and not empty strings")

	// ErrCreatingDBConnection is an error that catches a failed attempt to create a DB connection
	ErrCreatingDBConnection = errors.New("Error: There was an error when creating the DB connection")

	// ErrFailedToCloseDB is an error that catches a failed attempt to close a DB connection
	ErrFailedToCloseDB = errors.New("Error: There was an error when closing the DB connection")

	// ErrFailedToPingDB is an error that catches a failed attempt to ping the DB
	ErrFailedToPingDB = errors.New("Error: There was an error when pinging the DB")
)

// DBHandler is a struct that allows a connection to the DB and holds the config for the connection
type DBHandler struct {
	Db  *sql.DB
	cfg *Config
}

// Config is a struct that holds the information needed to open a connection to a DB
type Config struct {
	Host     string
	Port     string
	User     string
	Password string

	// DatabaseName to connect to (must have been created prior)
	DatabaseName string

	// DatabaseType is used to create the driver to a certain DB types
	DatabaseType string
}

// New returns an instance of the DBHandler Struct
func New(host, port, user, password, dbName, dbType string) (*DBHandler, error) {
	if host == "" || port == "" || user == "" || password == "" || dbName == "" || dbType == "" {
		return nil, ErrMissingArgument
	}

	cfg := &Config{
		Host:         host,
		Port:         port,
		User:         user,
		Password:     password,
		DatabaseName: dbName,
		DatabaseType: dbType,
	}

	db, err := openDBConnection(cfg)
	if err != nil {
		return nil, err
	}

	if err := pingDBConnection(db); err != nil {
		return nil, err
	}

	d := &DBHandler{
		Db:  db,
		cfg: cfg,
	}

	return d, nil
}

func openDBConnection(cfg *Config) (*sql.DB, error) {
	db, err := sql.Open(cfg.DatabaseType, fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		cfg.User, cfg.Password, cfg.DatabaseName, cfg.Host, cfg.Port))

	if err != nil {
		return nil, ErrCreatingDBConnection
	}

	return db, nil
}

func pingDBConnection(db *sql.DB) error {
	if err := db.Ping(); err != nil {
		return ErrFailedToPingDB
	}

	return nil
}

// Close performs the release of any resources that `sql/database` DB pool created. This is usually meant
// to be used in the exiting of a program or `panic`ing.
func (d *DBHandler) Close() error {
	if d.Db == nil {
		return nil
	}

	if err := d.Db.Close(); err != nil {
		return ErrFailedToCloseDB
	}

	return nil
}
