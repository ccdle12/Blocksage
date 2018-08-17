package injector

// "github.com/ccdle12/Blocksage/go-api/bitcoinclient"
// "github.com/ccdle12/Blocksage/go-api/bitcoincrawler"

// DependencyInjector is a struct that will return requested objects
type DependencyInjector struct{}

// var (
// 	btcMainDomain = os.Getenv("BTC_MAIN_DOMAIN")
// 	mainnetAPI    = api.API{}
// 	router        = mux.NewRouter()
// )

// InjectMainnetAPI will return an initialised API struct
// func (d *DependencyInjector) InjectMainnetAPI() *api.API {
// 	mainnetAPI.BitcoinClient = d.InjectBitcoinClient()
// 	return &mainnetAPI
// }

// // InjectRouter will return the mux Router
// func (d *DependencyInjector) InjectRouter() *mux.Router {
// 	return router
// }

// InjectBitcoinClient will create and return a BitcoinClient struct
// func (d *DependencyInjector) InjectBitcoinClient() *bitcoinclient.BitcoinClient {
// 	return &bitcoinclient.BitcoinClient{
// 		Client:          &http.Client{Timeout: time.Duration(5 * time.Second)},
// 		BitcoinNodeAddr: fmt.Sprintf("http://%s:8332", btcMainDomain),
// 	}
// }

// InjectBitcoinCrawler will create and return a BitcoinCrawler struct
// func (d *DependencyInjector) InjectBitcoinCrawler() *bitcoincrawler.BitcoinCrawler {
// 	return &bitcoincrawler.BitcoinCrawler{
// 		BitcoinClient: d.InjectBitcoinClient(),
// 	}
// }
