package usecases

// DBUsecase is the usecase/implementation interface for all DB usecases.
type DBUsecase interface {
	OpenConnection() error
	CloseConnection() error
}
