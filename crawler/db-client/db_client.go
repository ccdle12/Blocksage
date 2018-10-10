package dbclient

import (
	"errors"
	"github.com/ccdle12/Blocksage/crawler/db-client-uc"
	"github.com/ccdle12/Blocksage/crawler/models"
	"github.com/ccdle12/Blocksage/crawler/utils"
)

// Options is used in "functional options" pattern, essentially a Builder Pattern.
// It allows the Client constructor arguments to be passed in any order and optionally.
type Options struct {
	dbClient *Client
}

// Option is a type of function that is used to pass arguments to the constructor.
type Option func(*Options)

// Client is a struct that manages a connection to the DB and provides functions to read/write to the DB.
type Client struct {
	cfg     *models.DBConfig
	usecase dbuc.Usecase
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

// PostgresClient is a functional paramater to initialize the Postgres Client usecase in config.
func PostgresClient() Option {
	return func(args *Options) {
		args.dbClient.usecase = dbuc.NewPostGresClient(args.dbClient.cfg)
		args.dbClient.cfg.DBType = "postgres"
	}
}

// Test is a functional parameter to use the client in test mode, meaning all data will be written to the
// test tables.
func Test() Option {
	return func(args *Options) {
		args.dbClient.cfg.Test = true
	}
}

// New will return an instance of the Client that can read/write using different
// db usecases.
func New(setters ...Option) (*Client, error) {
	// Init args variable as an Options struct. This struct will hold a dbClient (concrete implementation).
	args := &Options{
		dbClient: &Client{
			cfg: &models.DBConfig{},
		},
	}

	// Loop over each option (setters) passed to the constructor and pass the args
	// struct to each optional function. This will init each variable in the dbClient config.
	for _, setter := range setters {
		setter(args)
	}

	// Make a check to make sure all the configs are not empty strings.
	cfg := args.dbClient.cfg
	if utils.EmptyString(cfg.DBHost, cfg.DBName, cfg.DBPassword, cfg.DBPort, cfg.DBType, cfg.DBUser) {
		return nil, utils.ErrPassingEmptyString
	}

	return args.dbClient, nil
}

// Connect will request the usecase to open a connection to the DB.
func (d *Client) Connect() error {
	if err := d.usecase.OpenConnection(); err != nil {
		return err
	}

	return nil
}

// Close will request the usecase to close the connection to the DB.
func (d *Client) Close() error {
	if err := d.usecase.CloseConnection(); err != nil {
		return err
	}

	return nil
}

// WriteBlock will request the usecase to write a block to the DB.
func (d *Client) WriteBlock(block *models.Block) error {
	if block == nil {
		return errors.New("Block is nil")
	}

	if err := d.usecase.InsertBlock(block); err != nil {
		return err
	}

	return nil
}

// WriteTransaction will request the usecase to write a transaction to the DB.
func (d *Client) WriteTransaction(tx *models.Transaction) error {
	if tx == nil {
		return errors.New("Block is nil")
	}

	if err := d.usecase.InsertTransaction(tx); err != nil {
		return err
	}

	return nil
}

// TODO (ccdle12)
func (d *Client) WriteInput(txHash string, in models.TransactionInput) error {
	if err := d.usecase.InsertInput(txHash, in); err != nil {
		return err
	}

	return nil
}
