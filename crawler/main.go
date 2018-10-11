package main

import (
	"github.com/ccdle12/Blocksage/crawler/db-client"
	"github.com/ccdle12/Blocksage/crawler/indexer"
	"github.com/ccdle12/Blocksage/crawler/injector"
	"github.com/ccdle12/Blocksage/crawler/node-client"
	"log"
)

func main() {
	// Create db client controller.
	var db dbclient.Controller
	db, err := dbclient.New(
		dbclient.DBPort(injector.PostgresPort()),
		dbclient.DBName(injector.PostgresDBName()),
		dbclient.DBUser(injector.PostgresUserName()),
		dbclient.DBHost(injector.PostgresDomain()),
		dbclient.DBPassword(injector.PostgresPassword()),
		dbclient.PostgresClient())
	if err != nil {
		log.Fatalf(err.Error())
	}

	// Create connection to the DB.
	if err = db.Connect(); err != nil {
		log.Fatalf(err.Error())
	}
	defer db.Close()

	// Create a node Controller to communicate with a Blockchain Node.
	var node nodeclient.Controller
	node = nodeclient.New(
		injector.DefaultHTTPClient(),
		injector.BTCDomain(),
		injector.BTCUsername(),
		injector.BTCPassword())

	// Create the indexer.
	indexer := indexer.New(node, db)
	// TODO (ccdle12): "Hard Code" Genesis Block in the DB.
	if err = indexer.Crawl(injector.Block1Hash); err != nil {
		log.Fatalf(err.Error())
	}
}
