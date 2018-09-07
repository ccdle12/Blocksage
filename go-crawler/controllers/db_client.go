package controllers

import (
	"github.com/ccdle12/Blocksage/go-crawler/models"
	"github.com/ccdle12/Blocksage/go-crawler/usecases"
	"github.com/ccdle12/Blocksage/go-crawler/utils"
)

// Options is used in "functional options" pattern, essentially a Builder Pattern.
// It allows the NewDBClient constructor arguments to be passed in any order.
type Options struct {
	dbClient *DBClient
}

// Option is a type of function that is used to pass arguments to the constructor.
type Option func(*Options)

// DBClient is a struct that manages a connection to the DB and provides functions to read/write to the DB.
type DBClient struct {
	cfg     *models.DBConfig
	usecase usecases.DBClient
}

// DBHost is functional paramater to initialize DBHost in config.
func DBHost(host string) Option {
	return func(args *Options) {
		args.dbClient.cfg.DBHost = host
	}
}

// DBPort is functional paramater to initialize DBPort in config.
func DBPort(port string) Option {
	return func(args *Options) {
		args.dbClient.cfg.DBPort = port
	}
}

// DBUser is functional paramater to initialize DBUser in config.
func DBUser(user string) Option {
	return func(args *Options) {
		args.dbClient.cfg.DBUser = user
	}
}

// DBPassword is functional paramater to initialize DBPassword in config.
func DBPassword(pw string) Option {
	return func(args *Options) {
		args.dbClient.cfg.DBPassword = pw
	}
}

// DBName is functional paramater to initialize DBName in config.
func DBName(name string) Option {
	return func(args *Options) {
		args.dbClient.cfg.DBName = name
	}
}

// DBType is functional paramater to initialize DBType in config.
func DBType(dbType string) Option {
	return func(args *Options) {
		args.dbClient.cfg.DBType = dbType
	}
}

// NewDBClient will return an instance of the DBClient that will read/write from different
// usecases.
func NewDBClient(setters ...Option) (*DBClient, error) {
	// Init args which is Options and inits a DBClient.
	args := &Options{
		dbClient: &DBClient{
			cfg: &models.DBConfig{},
		},
	}

	// Loop over each argument (optional function) passed to the constructor and pass the args
	// struct to each argument (optional function). This will init each variable in the DBClients
	// config.
	for _, setter := range setters {
		setter(args)
	}

	// Make a check to make sure all the configs are not empty strings.
	cfg := args.dbClient.cfg
	if utils.EmptyString(cfg.DBHost, cfg.DBName, cfg.DBPassword, cfg.DBPort, cfg.DBType, cfg.DBUser) {
		return nil, utils.ErrPassingEmptyString
	}

	// Init the usecase for the dbClient.
	args.dbClient.usecase = usecases.NewPostGresClient(args.dbClient.cfg)

	return args.dbClient, nil
}

// TODO (ccdle12): Needs Integration Testing
// Connect will request the usecase to open a connection to the DB.
func (d *DBClient) Connect() error {
	if err := d.usecase.OpenConnection(); err != nil {
		return err
	}

	return nil
}

// TODO (ccdle12): Needs Integration Testing
// Close will request the usecase to close the connection to the DB.
func (d *DBClient) Close() error {
	if err := d.usecase.CloseConnection(); err != nil {
		return err
	}

	return nil
}
