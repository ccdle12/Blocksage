package nodeuc

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/ccdle12/Blocksage/go-crawler/models"
	"github.com/ccdle12/Blocksage/go-crawler/utils"
	"net/http"
)

// BitcoinClient is the struct that will handle the BitcoinClient Usecase implementation.
type BitcoinClient struct{}

// NewBitcoinClient is the constructor for BitcoinClient
func NewBitcoinClient() *BitcoinClient {
	return &BitcoinClient{}
}

// SendNodeRequest will run all un-exported functions needed to send a request to the
// Blockchain Node.
func (b *BitcoinClient) SendNodeRequest(nodeReq *models.NodeRequest) (*models.NodeResponse, error) {
	req, err := b.createRequest(nodeReq)
	if err != nil {
		return nil, err
	}

	res, err := b.sendRequest(nodeReq, req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	resByte, err := utils.ConvBodyToByte(res)
	if err != nil {
		return nil, err
	}

	nodeRes, err := utils.ConvByteToNodeRes(resByte)
	if err != nil {
		return nil, err
	}

	if err := b.handleNodeResponse(nodeRes); err != nil {
		return nil, err
	}

	return nodeRes, nil
}

// createRequest builds an *http.Request object to send the Blockchain Node.
func (b *BitcoinClient) createRequest(nodeReq *models.NodeRequest) (*http.Request, error) {
	body, err := b.createBody(nodeReq.Body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", nodeReq.Headers.Address, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(nodeReq.Headers.Username, nodeReq.Headers.Password)

	return req, nil
}

// createBody marshals the *models.NodeReqBody to JSON format.
func (b *BitcoinClient) createBody(body models.NodeReqBody) (*bytes.Buffer, error) {
	bodyJSON, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	return bytes.NewBuffer(bodyJSON), nil
}

// sendRequest, sends the http request to the Blockchain Node.
func (b *BitcoinClient) sendRequest(nodeReq *models.NodeRequest, req *http.Request) (*http.Response, error) {
	res, err := nodeReq.Client.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// handleNodeResponse will check the NodeResponse for an error. If there Error.Code is not 0, then there was
// an error. We return the error message from the NodeResponse.
func (b *BitcoinClient) handleNodeResponse(nodeRes *models.NodeResponse) error {
	if nodeRes.Error.Code != 0 {
		return errors.New(nodeRes.Error.Message)
	}

	return nil
}
