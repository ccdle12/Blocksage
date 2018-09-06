package injector

import (
	"github.com/ccdle12/Blocksage/go-crawler/models"
	"net/http"
	"time"
)

// These variable are used throughout the project and can be referend here for dependency injection.
// TODO (ccdle12): Move to an injector file?
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
