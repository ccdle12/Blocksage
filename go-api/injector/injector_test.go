// +build unit

package injector

import (
	"testing"
	"time"
)

var (
	inj           = DependencyInjector{}
	API           = inj.InjectMainnetAPI()
	Router        = inj.InjectRouter()
	BitcoinClient = inj.InjectBitcoinClient()
)

func TestInjectMainnetAPI(t *testing.T) {
	if API.BitcoinClient == nil {
		t.Errorf("API did not instantiate the BitcoinClient Field")
	}

	if API.BitcoinClient.Client.Timeout != time.Duration(5*time.Second) {
		t.Errorf("API Bitcoin Client did not set the Timeout time correctly. Received: %v, It Should Be: %v", API.BitcoinClient.Client.Timeout, time.Duration(5*time.Second))
	}
}

func TestInjectRouter(t *testing.T) {
	if Router == nil {
		t.Errorf("Injector did not instantiate Router")
	}
}

func TestInjectBitcoinClient(t *testing.T) {
	if BitcoinClient == nil {
		t.Errorf("Injector did not instantiate BitcoinClient")
	}

	if BitcoinClient.Client.Timeout != time.Duration(5*time.Second) {
		t.Errorf("API Bitcoin Client did not set the Timeout time correctly. Received: %v, It Should Be: %v", API.BitcoinClient.Client.Timeout, time.Duration(5*time.Second))
	}
}
