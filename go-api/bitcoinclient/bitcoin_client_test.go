package bitcoinclient

import (
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"
)

var (
	btcMainDomain = os.Getenv("BTC_MAIN_DOMAIN")
	bitcoinClient = &BitcoinClient{
		Client:          &http.Client{Timeout: time.Duration(5 * time.Second)},
		BitcoinNodeAddr: fmt.Sprintf("http://%s:8332", btcMainDomain),
	}

	blockHash = "000000000019d6689c085ae165831e934ff763ae46a2a6c172b3f1b60a8ce26f"
	txHash    = "b1fea52486ce0c62bb442b530a3f0132b826c74e473d1f2c220bfa78111c5082"
)

func TestGetNetworkInfo(t *testing.T) {
	// Shoudl get network info
	res, err := bitcoinClient.GetNetworkInfo()
	if err != nil {
		t.Errorf("GetNetworkInfo has failed, error returned: %v", err.Error())
	}

	if res == "" {
		t.Errorf("GetNetworkInfo has failed, Received: %v", res)
	}

	rpcBitcoinResponse := bitcoinClient.createRPCBitcoinResponse(res)

	if rpcBitcoinResponse.Error.Code != 0 {
		t.Errorf("GetNetworkInfo has returned an error, Received: %v", rpcBitcoinResponse.Error.Message)
	}
}
func TestGetBlock(t *testing.T) {
	// Shoudl get block hash
	res, err := bitcoinClient.GetBlock(blockHash)
	if err != nil {
		t.Errorf("GetBlock has failed, error returned: %v", err.Error())
	}

	if res == "" {
		t.Errorf("GetBlock has failed, Received: %v", res)
	}

	rpcBitcoinResponse := bitcoinClient.createRPCBitcoinResponse(res)

	if rpcBitcoinResponse.Error.Code != 0 {
		t.Errorf("GetBlock has returned an error, Received: %v", rpcBitcoinResponse.Error.Message)
	}
}
func TestGetTransaction(t *testing.T) {
	// Should get transaction
	res, err := bitcoinClient.GetTransaction(txHash)
	if err != nil {
		t.Errorf("GetTransaction has failed, error returned: %v", err.Error())
	}

	if res == "" {
		t.Errorf("GetTransaction has failed, Received: %v", res)
	}

	rpcBitcoinResponse := bitcoinClient.createRPCBitcoinResponse(res)

	if rpcBitcoinResponse.Error.Code != 0 {
		t.Errorf("GetTransaction has returned an error, Received: %v", rpcBitcoinResponse.Error.Message)
	}
}
