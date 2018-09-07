package utils

import (
	"encoding/json"
	"github.com/ccdle12/Blocksage/go-crawler/models"
	"io/ioutil"
	"net/http"
	"strings"
)

// ConvBodyToByte is a function that will convert the body of a http.Response to a []byte.
func ConvBodyToByte(res *http.Response) ([]byte, error) {
	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return bodyBytes, nil
}

// ConvByteToNodeRes is a function that will convert a []byte to a NodeResponse.
func ConvByteToNodeRes(resByte []byte) (*models.NodeResponse, error) {
	var nodeRes *models.NodeResponse
	if err := json.Unmarshal(resByte, &nodeRes); err != nil {
		return nil, err
	}

	return nodeRes, nil
}

// ConvNodeResToBlock is a function that will convert a NodeRespone to a Block.
func ConvNodeResToBlock(nodeRes *models.NodeResponse) (*models.Block, error) {
	out, err := json.Marshal(nodeRes.Result)
	if err != nil {
		return nil, err
	}

	var block *models.Block
	if err := json.Unmarshal(out, &block); err != nil {
		return nil, err
	}

	return block, nil
}

// EmptyString will return a bool if a string is empty (zero value)
func EmptyString(input ...string) bool {

	// Iterate over all inputs, if one of the inputs is an empty string,
	// exit and return true
	for _, s := range input {
		if s == "" {
			return true
		}
	}

	return false
}

// FormatAddress will receive an address and format it, to be used in the project.
func FormatAddress(address string) string {
	if strings.HasPrefix(address, "http://") {
		return address
	}

	return "http://" + address
}
