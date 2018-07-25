package injector

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
	btcMainDomain = os.Getenv("BTC_MAIN_DOMAIN")
	mainnetAPI    = api.API{}
	router        = mux.NewRouter()
)

// InjectMainnetAPI will return an initialised API struct
func (d *DependencyInjector) InjectMainnetAPI() *api.API {
	mainnetAPI.BitcoinClient = d.InjectBitcoinClient()
	return &mainnetAPI
}

// InjectRouter will return the mux Router
func (d *DependencyInjector) InjectRouter() *mux.Router {
	return router
}

// InjectBitcoinClient will create and return a BitcoinClient struct
func (d *DependencyInjector) InjectBitcoinClient() *bitcoinclient.BitcoinClient {
	return &bitcoinclient.BitcoinClient{
		Client:          &http.Client{Timeout: time.Duration(5 * time.Second)},
		BitcoinNodeAddr: fmt.Sprintf("http://%s:8332", btcMainDomain),
	}
}
