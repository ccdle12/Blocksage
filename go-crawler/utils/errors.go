package utils

import (
	"errors"
)

// These Errors will be injected throughout the app
// TODO (ccdle12): Move to its own go file?
var (
	ErrPassingEmptyString        = errors.New("Cannot pass empty string as an argument")
	ErrFailedToOpenDBConnection  = errors.New("Could not open connection to the DB")
	ErrFailedToPingDB            = errors.New("Could not Ping DB connection")
	ErrFailedToCloseDBConnection = errors.New("Failed to close DB connection")
	ErrFailedToInsertToDB        = errors.New("Failed to insert row to the DB")
)
