// Package injector provides behavior to inject certain structs and variables.
package injector

import (
	"github.com/ccdle12/Blocksage/go-crawler/models"
	"net/http"
	"os"
	"time"
)

var (
	defaultHTTPClient = &http.Client{Timeout: time.Duration(5 * time.Second)}
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
	return os.Getenv("USERNAME")
}

// BTCPassword will retrieve the environment variable for the username of the BTC Mainnet Node.
func BTCPassword() string {
	return os.Getenv("PASSWORD")
}
