package bitcoinclient

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// BitcoinRPC describes an expected set of behaviours to create/send requests and create a body
type BitcoinRPC interface {
	createRequest(s string) (*http.Request, error)
	createBody(s string) (*bytes.Buffer, error)
	sendRequest(r *http.Request) (*http.Response, error)
	convBodyToStr(r *http.Response) (string, error)
}

// methodBody is a struct for the body when sending POST requests to the Bitcoin Node RPC
type methodBody struct {
	Method string `json:"method"`
}

var (
	username = os.Getenv("USERNAME")
	password = os.Getenv("PASSWORD")
)

// createRequest builds a request to send the Bitcoin Node
func (b *BitcoinClient) createRequest(methodType string) (*http.Request, error) {
	body, err := b.createBody(methodType)
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

// createBody, creates a body specifying the RPC methods to request from the Bitcoin Node
func (b *BitcoinClient) createBody(methodType string) (*bytes.Buffer, error) {
	body := methodBody{methodType}
	bodyJSON, err := json.Marshal(body)
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
		log.Println(err)
		return "", ErrConvertingBodyToString
	}

	return string(bodyBytes), nil
}
