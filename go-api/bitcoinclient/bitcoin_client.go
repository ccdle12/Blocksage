package bitcoinclient

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

// BitcoinClient is the struct responsible for communicating to the Bitcoin Node
type BitcoinClient struct {
	Client          *http.Client
	BitcoinNodeAddr string
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

// GetNetworkInfo returns a string response from the Bitcoin Node, requesting information about the node's conection to the network
func (b *BitcoinClient) GetNetworkInfo() (string, error) {
	return b.getResponse(&RPCBody{"getnetworkinfo", []string{}})
}

// GetBlock returns a string response from the Bitcoin Node, requesting information of a Block according the blockHash passed as an argument
func (b *BitcoinClient) GetBlock(blockHash string) (string, error) {
	return b.getResponse(&RPCBody{"getblock", []string{blockHash}})
}

// GetTransaction returns a string response from the Bitcoin Node, requesting information on a Transaction according to the TransactionHash passed as an argument
func (b *BitcoinClient) GetTransaction(txHash string) (string, error) {
	response, err := b.getResponse(&RPCBody{"getrawtransaction", []string{txHash}})
	if err != nil {
		return err.Error(), err
	}

	return b.decodeRawTransaction(response)
}

// decodeRawTransaction returns a string response from the Bitcoin Node, it takes in a rawTx string and decodes it to human readable format
func (b *BitcoinClient) decodeRawTransaction(rawTx string) (string, error) {
	var rpcResponse RPCBitcoinResponse
	json.Unmarshal([]byte(rawTx), &rpcResponse)

	return b.getResponse(&RPCBody{"decoderawtransaction", []string{*rpcResponse.Result}})
}

// getResponse is the generic function that will take in a *RPCBody struct and get a reponse from the Bitcoin Node
func (b *BitcoinClient) getResponse(rpcBody *RPCBody) (string, error) {
	req, err := b.createRequest(rpcBody)
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

	// Bitcoin Node returns an empty string if authentication fails and also returns a JSON for RPC requests
	rpcError := b.AuthenticateRPCResponse(resStr)
	if rpcError != nil {
		return err.Error(), err
	}

	return resStr, nil
}
