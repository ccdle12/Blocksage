package dependencyinjector

import (
	"github.com/ccdle12/Blocksage/go-api/api"
	"github.com/ccdle12/Blocksage/go-api/bitcoinclient"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

// DependencyInjector is a struct that will return requested objects
type DependencyInjector struct{}

var (
	singletonAPI = api.API{
		BitcoinClient: &bitcoinclient.BitcoinClient{
			Client:          &http.Client{Timeout: time.Duration(5 * time.Second)},
			BitcoinNodeAddr: "http://35.194.42.115:8332",
		},
		Router: mux.NewRouter(),
	}
)

// InjectAPI will return an initialised API struct
func (d *DependencyInjector) InjectAPI() *api.API {
	return &singletonAPI
}
