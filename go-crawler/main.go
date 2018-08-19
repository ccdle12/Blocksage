package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/ccdle12/Blocksage/go-crawler/bitcoinclient"
	"github.com/ccdle12/Blocksage/go-crawler/bitcoincrawler"
	"github.com/ccdle12/Blocksage/go-crawler/dbhandler"
)

// The following variables are retreived environment variables
var (
	btcMainDomain = os.Getenv("BTC_MAIN_DOMAIN")
)

func main() {
	// Create the connection to the DB
	dbHandler, err := dbhandler.New("172.25.0.4", "5432", "postgres-dev", "12345", "dev", "postgres")
	if err != nil {
		log.Fatal(err.Error())
	}

	bitcoinClient := &bitcoinclient.BitcoinClient{
		Client:          &http.Client{Timeout: time.Duration(5 * time.Second)},
		BitcoinNodeAddr: fmt.Sprintf("http://%s:8332", btcMainDomain),
	}

	bitcoinCrawler := bitcoincrawler.New(bitcoinClient, dbHandler)
	bitcoinCrawler.Start()
}
