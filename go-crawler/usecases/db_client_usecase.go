package usecases

// DBClient is the usecase/implementation interface for all DB usecases.
type DBClient interface {
	OpenConnection() error
	CloseConnection() error
}
