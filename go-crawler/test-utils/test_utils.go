package testutils

import (
	"github.com/ccdle12/Blocksage/go-crawler/models"
	"net/http"
	"net/http/httptest"
	"time"
)

// The following variables are global testing constants
const (
	// NodeClient variables
	Username        = "dev123"
	Password        = "secretpass"
	NodeAddress     = "http://123.45.67.8:8332"
	GetBlock        = "getblock"
	BlockHash0      = "000000000019d6689c085ae165831e934ff763ae46a2a6c172b3f1b60a8ce26f"
	IncorrectMethod = "getalltxs"

	// DBHandler variables
	DBHost     = "http://123.45.67.8:8332"
	DBPort     = "8332"
	DBUser     = "dev"
	DBPassword = "123"

	// DBName is the database to connect to
	DBName = "postgresdev"

	// DBType is the driver type for the DB
	DBType = "postgres"
)

var (
	Client = &http.Client{Timeout: time.Duration(5 * time.Second)}

	GetBlockParams = []string{"12345"}

	NodeHeaders = models.NodeHeaders{
		Address:  NodeAddress,
		Username: Username,
		Password: Password,
	}

	NodeBody = models.NodeReqBody{
		Method: GetBlock,
		Params: GetBlockParams,
	}

	NodeReq = &models.NodeRequest{
		Headers: NodeHeaders,
		Body:    NodeBody,
	}

	NodeResponse = &models.NodeResponse{
		Result: &models.Block{},
		Error:  models.NodeError{Code: 0, Message: ""},
		ID:     "",
	}

	DBConfig = &models.DBConfig{
		DBHost:     DBHost,
		DBPort:     DBPort,
		DBUser:     DBUser,
		DBPassword: DBPassword,
		DBName:     DBName,
		DBType:     DBType,
	}
)

// TestServer will return a http test server object. It takes in an expected response in the form of a string,
// it will return the expectedResponse, when sending requests to the test server.
// NOTE: call defer server.Close() right after initializing the test server.
func TestServer(expectedResponse string) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte(expectedResponse))
	}))
}
