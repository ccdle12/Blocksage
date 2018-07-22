package bitcoinclient

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// RPCBody is a struct for the body when sending POST requests to the Bitcoin Node RPC
type RPCBody struct {
	Method string   `json:"method"`
	Params []string `json:"params"`
}

// RPCBitcoinResponse is a struct that contains the format for returned information from the Bitcoin Node
type RPCBitcoinResponse struct {
	Result *string  `json:"result"`
	Error  rpcError `json:"error"`
	ID     *string  `json:"id"`
}

// rpcError is a struct the contains the format for errors returned from the Bitcoin Node
type rpcError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// TODO: The authentication variables will need to be changed according to testnet and mainnet
var (
	username = os.Getenv("USERNAME")
	password = os.Getenv("PASSWORD")
)

// createRequest builds a request to send the Bitcoin Node
func (b *BitcoinClient) createRequest(rpcBody *RPCBody) (*http.Request, error) {
	body, err := b.createBody(rpcBody)
	if err != nil {
		log.Println(ErrCreatingBody)
		return nil, ErrCreatingBody
	}

	req, err := http.NewRequest("POST", b.BitcoinNodeAddr, body)
	if err != nil {
		log.Println(ErrCreatingRequest)
		return nil, ErrCreatingRequest
	}

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(username, password)

	return req, nil
}

// createBody marshals the rpcBody to JSON format
func (b *BitcoinClient) createBody(rpcBody *RPCBody) (*bytes.Buffer, error) {
	bodyJSON, err := json.Marshal(rpcBody)
	if err != nil {
		log.Println(ErrCreatingBody)
		return nil, ErrCreatingBody
	}

	return bytes.NewBuffer(bodyJSON), nil
}

// sendRequest, sends the http request to the Bitcoin Node
func (b *BitcoinClient) sendRequest(req *http.Request) (*http.Response, error) {
	res, err := b.Client.Do(req)
	if err != nil {
		log.Println(ErrUnresponsive)
		return nil, ErrUnresponsive
	}

	return res, nil
}

// convBodyToStr, converts the whole body of a response to string
func (b *BitcoinClient) convBodyToStr(res *http.Response) (string, error) {
	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(ErrConvertingBodyToString)
		return "", ErrConvertingBodyToString
	}

	return string(bodyBytes), nil
}
