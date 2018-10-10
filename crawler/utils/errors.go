package utils

import (
	"errors"
	"fmt"
)

// These Errors will be injected throughout the app
// TODO (ccdle12): Move to its own go file?
var (
	ErrPassingEmptyString        = errors.New("Cannot pass empty string as an argument")
	ErrFailedToOpenDBConnection  = errors.New("Could not open connection to the DB")
	ErrFailedToPingDB            = errors.New("Could not Ping DB connection")
	ErrFailedToCloseDBConnection = errors.New("Failed to close DB connection")
	ErrFailedToInsertToDB        = errors.New("[Failed to insert row to the DB]")
)

// FailedDBConnection will take the error from the DB, pre-pend a simple message with the complex message.
func FailedDBConnection(err error) error {
	return formatErr(ErrFailedToOpenDBConnection, err)
}

// FailedToInsertToDB will take the error from the DB, pre-pend a simple message with the complex message.
func FailedToInsertToDB(err error) error {
	return formatErr(ErrFailedToInsertToDB, err)
}

func formatErr(simpleErr, err error) error {
	message := fmt.Sprintf(simpleErr.Error() + ": " + err.Error())

	return errors.New(message)
}
