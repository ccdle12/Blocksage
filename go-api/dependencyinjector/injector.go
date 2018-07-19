package dependencyinjector

import (
	"fmt"
	"github.com/ccdle12/Blocksage/go-api/api"
	"github.com/ccdle12/Blocksage/go-api/bitcoinclient"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"time"
)

// DependencyInjector is a struct that will return requested objects
type DependencyInjector struct{}

var (
	btcMainnetDomain = os.Getenv("BTC_MAINNET_DOMAIN_EXTERNAL")
	singletonAPI     = api.API{
		BitcoinClient: &bitcoinclient.BitcoinClient{
			Client:          &http.Client{Timeout: time.Duration(5 * time.Second)},
			BitcoinNodeAddr: fmt.Sprintf("http://%s:8332", btcMainnetDomain),
		},
		Router: mux.NewRouter(),
	}
)

// InjectAPI will return an initialised API struct
func (d *DependencyInjector) InjectAPI() *api.API {
	return &singletonAPI
}
