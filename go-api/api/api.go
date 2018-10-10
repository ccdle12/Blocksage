package api

// "github.com/ccdle12/Blocksage/go-api/bitcoinclient"

// API struct is responsible for connecting the requests from the client side with functions specific to each route
// type API struct {
// 	BitcoinClient *bitcoinclient.BitcoinClient
// }

// // NetworkInfo is responsible for returning information about the Bitcoin Nodes connection to the network
// func (a *API) NetworkInfo(w http.ResponseWriter, r *http.Request) {
// 	response, err := a.BitcoinClient.GetNetworkInfo()
// 	if err != nil {
// 		handledWriter := a.BitcoinClient.HandleStatusCodeError(w, err)
// 		fmt.Fprintf(handledWriter, err.Error())
// 		return
// 	}

// 	fmt.Fprintf(w, response)
// }

// // Blocks is responsible for returning information about a certain Block in the Bitcoin Blockchain
// func (a *API) Blocks(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	response, err := a.BitcoinClient.GetBlock(vars["blockhash"])
// 	if err != nil {
// 		handledWriter := a.BitcoinClient.HandleStatusCodeError(w, err)
// 		fmt.Fprintf(handledWriter, err.Error())
// 		return
// 	}

// 	fmt.Fprintf(w, response)
// }

// // Transactions is responsible for returning information about a certain Transaction in the Bitcoin Blockchain
// func (a *API) Transactions(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	response, err := a.BitcoinClient.GetTransaction(vars["txhash"])
// 	if err != nil {
// 		handledWriter := a.BitcoinClient.HandleStatusCodeError(w, err)
// 		fmt.Fprintf(handledWriter, err.Error())
// 		return
// 	}

// 	fmt.Fprintf(w, response)
// }
