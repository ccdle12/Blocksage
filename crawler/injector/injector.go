// Package injector provides behavior to inject certain structs and variables.
package injector

import (
	"github.com/ccdle12/Blocksage/crawler/models"
	"net/http"
	"os"
	"strings"
	"time"
)

var (
	defaultHTTPClient = &http.Client{Timeout: time.Duration(20 * time.Second)}
	Block1Hash        = "000000006a625f06636b8bb6ac7b960a8d03705d1ace08b1a19da3fdcc99ddbd"
)

// NodeRequest will create a NodeRequest, needed to communicate client details from the
// controller to the usecase.
func NodeRequest(client *http.Client, nodeAddress, username, password, method string, params []string) *models.NodeRequest {
	return &models.NodeRequest{
		Client: client,
		Headers: models.NodeHeaders{
			Address:  nodeAddress,
			Username: username,
			Password: password,
		},
		Body: models.NodeReqBody{
			Method: method,
			Params: params,
		},
	}
}

// DefaultHTTPClient will inject the var defaultClient, this maybe modified to pass in a time in seconds to create an *httpClient
func DefaultHTTPClient() *http.Client {
	return defaultHTTPClient
}

// BTCDomain will retrieve the environment variable for the address of the BTC Mainnet Node.
func BTCDomain() string {
	return os.Getenv("BTC_MAIN_DOMAIN")
}

// BTCUsername will retrieve the environment variable for the password of the BTC Mainnet Node.
func BTCUsername() string {
	return os.Getenv("BTC_USERNAME")
}

// BTCPassword will retrieve the environment variable for the username of the BTC Mainnet Node.
func BTCPassword() string {
	return os.Getenv("BTC_PASSWORD")
}

// PostgresDBName will retrieve the environment variable for the DB Name of the Postgres DB.
func PostgresDBName() string {
	return os.Getenv("POSTGRES_DB_NAME")
}

// PostgresUserName will retrieve the environment variable for the Username of the Postgres DB.
func PostgresUserName() string {
	return os.Getenv("POSTGRES_USER")
}

// PostgresPassword will retrieve the environment variable for the Password of the Postgres DB.
func PostgresPassword() string {
	return os.Getenv("POSTGRES_PASSWORD")
}

// PostgresDomain will retrieve the environment variable for the Postgres Domain of the Postgres DB.
func PostgresDomain() string {
	domain := splitDBDomain()[0]

	return domain
}

// PostgresPort will retrieve the environment variable for the Postgres Port of the Postgres DB.
func PostgresPort() string {
	host := splitDBDomain()[1]

	return host
}

// splitDBDomain will split the DB into to items into a []string{domain, host}.
// The reasoning behind splitting the environment variable is when deployed using k8's,
// we are going to use DNS to link each service. The DNS route the end points as <host:domain>.
func splitDBDomain() []string {
	address := os.Getenv("POSTGRES_DB_DOMAIN")

	return strings.Split(address, ":")
}
