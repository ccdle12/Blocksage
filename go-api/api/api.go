package api

import (
	"fmt"
	"github.com/ccdle12/Blocksage/go-api/bitcoinclient"
	"github.com/gorilla/mux"
	"net/http"
)

// API struct is responsible for connecting the end user with the below functions for each route specified in main
type API struct {
	BitcoinClient *bitcoinclient.BitcoinClient
	Router        *mux.Router
}

// NetworkInfo is responsible for returning information about the Bitcoin Nodes connection to the network
func (a *API) NetworkInfo(w http.ResponseWriter, r *http.Request) {
	response, err := a.BitcoinClient.GetNetworkInfo()
	if err != nil {
		handledWriter := a.BitcoinClient.HandleStatusCodeError(w, err)
		fmt.Fprintf(handledWriter, err.Error())
		return
	}

	fmt.Fprintf(w, response)
}
