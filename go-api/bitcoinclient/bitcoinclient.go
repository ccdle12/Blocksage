package bitcoinclient

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// BitcoinClient is the struct that builds the requests, sends the requests and validates
// responses from the Bitcoin Node RPC
type BitcoinClient struct {
	Client          *http.Client
	BitcoinNodeAddr string
	bitcoinRPC      bitcoinRPC
}

func init() {
	// Log Print the line number of where the error occurred
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Write error log prints to log.txt
	nf, err := os.Create("./bitcoinclient/log.txt")
	if err != nil {
		fmt.Println("Non-Crucial Internal Error - Cannot create log.txt file for bitcoinclient.go")
		return
	}

	log.SetOutput(nf)
}

// GetNetworkInfo returns a string response from the Bitcoin Node requesting "getnetworkinfo"
// "getnetworkinfo" returns information about the node's conection to the network
func (b *BitcoinClient) GetNetworkInfo() (string, error) {
	req, err := b.createRequest("getnetworkinfo")
	if err != nil {
		return err.Error(), err
	}

	res, err := b.sendRequest(req)
	if err != nil {
		return err.Error(), err
	}
	defer res.Body.Close()

	resStr, err := b.convBodyToStr(res)
	if err != nil {
		return err.Error(), err
	}

	// Bitcoin Node returns an empty string if authentication fails and also returns a JSON for RPC errors
	rpcError := b.AuthenticateRPCResponse(resStr)
	if rpcError != nil {
		return rpcError.Error(), rpcError
	}

	return resStr, nil
}
