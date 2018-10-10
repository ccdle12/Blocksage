package models

import (
	"net/http"
)

// NodeRequest is a struct that packages all the information needed to make a request
// to the Node.
type NodeRequest struct {
	Client  *http.Client
	Headers NodeHeaders
	Body    NodeReqBody
}

// NodeResponse is a struct that contains the format for returned information from the Node.
type NodeResponse struct {
	Result interface{} `json:"result"`
	Error  NodeError   `json:"error"`
	ID     string      `json:"id"`
}

// NodeHeaders is a struct that will hold header information for the request to the node.
type NodeHeaders struct {
	Address  string
	Username string
	Password string
}

// NodeReqBody is a struct for the body when sending POST requests to the via the Node Client.
type NodeReqBody struct {
	Method string   `json:"method"`
	Params []string `json:"params"`
}

// NodeError is a struct that contains the format for errors returned from the Node.
type NodeError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
