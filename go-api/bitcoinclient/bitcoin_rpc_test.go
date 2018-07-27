// +build integration

package bitcoinclient

import (
	"testing"
)

func TestCreateRPCBitcoinResponse(t *testing.T) {
	// Should Create a RPCBitcoinResponse
	res, err := bitcoinClient.GetTransaction(txHash)
	if err != nil {
		t.Errorf("GetTransaction has failed, error returned: %v", err.Error())
	}

	if res == "" {
		t.Errorf("GetTransaction has failed, Received: %v", res)
	}

	rpcBitcoinResponse := bitcoinClient.createRPCBitcoinResponse(res)

	if rpcBitcoinResponse == nil {
		t.Errorf("CreateRPCBitcoinResponse did not instantiate a struct")
	}
}
