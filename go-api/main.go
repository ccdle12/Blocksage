package main

// "github.com/ccdle12/Blocksage/go-api/dbclient"

func main() {
	//DBHanlder
	// TODO: Intialise an Object to pass as argument?
	// _, err := dbclient.New("172.25.0.4", "5432", "postgres-dev", "12345", "dev", "postgres")
	// if err != nil {
	// 	fmt.Println("There was an error")
	// }

	// blockTable, err := blocktable.New(dbClient)
	// if err != nil {
	// 	fmt.Println("There was an error create blocktable")
	// }

	// if err := blockTable.CreateTable(); err != nil {
	// 	fmt.Println("Error create blockTable")
	// }

	// TODO: Throwing crawler here, will be separated later
	// inj := injector.DependencyInjector{}

	// // go func() {
	// // 	bitcoinCrawler := inj.InjectBitcoinCrawler()
	// // 	bitcoinCrawler.Start()
	// // }()

	// API := inj.InjectMainnetAPI()
	// router := inj.InjectRouter()

	// router.HandleFunc("/api/network-info", API.NetworkInfo).Methods("GET")
	// router.HandleFunc("/api/blocks/{blockhash}", API.Blocks).Methods("GET")
	// router.HandleFunc("/api/txs/{txhash}", API.Transactions).Methods("GET")

	// fmt.Println("Serving on Port 8548...")
	// log.Fatal(http.ListenAndServe(":8548", router))
}
