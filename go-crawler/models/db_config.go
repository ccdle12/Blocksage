package models

// DBConfig is a struct that holds all the relevant information needed for a
// Database Handler/Controller to create a connection with a Database.
type DBConfig struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string

	// DBName to connect to (must have been created prior)
	DBName string

	// DBType is used to create the driver to a certain DB types
	DBType string
}
