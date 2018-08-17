package bitcoincrawler

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/ccdle12/Blocksage/go-crawler/bitcoinclient"
	"github.com/ccdle12/Blocksage/go-crawler/dbhandler"
	"github.com/ccdle12/Blocksage/go-crawler/models"
	"github.com/ccdle12/Blocksage/go-crawler/tables"
)

// BitcoinCrawler is a struct that will crawl and retrieve data from
// the Blockchain
type BitcoinCrawler struct {
	BTCClient  *bitcoinclient.BitcoinClient
	DBHandler  *dbhandler.DBHandler
	BlockTable *blocktable.BlockTable
}

// BitcoinBlockResponse inherits from RPCBitcoinResponse, the block
// struct overwrites the Result field in RPCBitcoinResponse
type BitcoinBlockResponse struct {
	RPCBitcoinResponse bitcoinclient.RPCBitcoinResponse
	Result             *models.Block `json:"result"`
}

var (
	// The crawler will start crawling blocks from the genesis
	genesisBlockHash = "000000000019d6689c085ae165831e934ff763ae46a2a6c172b3f1b60a8ce26f"

	// TODO (ccdle12): The following blocks are used for debugging and development
	block533979 = "0824fec1580a505a2f363347777aeff54242914316c40ab93d6fee6b4ddbf48c"
	block536486 = "000000000000000000041302a5c5d322d6cd46f74cfb9e6d460b832bad920433"
	block537242 = "00000000000000000011153335dedf5d455c19fc87de71df3c9f10b791fd8dfb"

	// wg is the wait group used as an async lock
	wg sync.WaitGroup
)

// New returns an instance of the BitcoinCrawler
func New(bitcoinClient *bitcoinclient.BitcoinClient, dbHandler *dbhandler.DBHandler, blockTable *blocktable.BlockTable) *BitcoinCrawler {
	return &BitcoinCrawler{
		BTCClient:  bitcoinClient,
		DBHandler:  dbHandler,
		BlockTable: blockTable,
	}
}

// Start will run the crawler and request blocks from the BTC Node,
// once the crawler has reached the highest block, it will start
// polling the BTC Node for the next block every minute
func (b *BitcoinCrawler) Start() {
	latestBlock := b.crawl(block537242)
	fmt.Println("This is about to execute pollForBlocks(latestBlock)")
	wg.Add(1)
	go b.pollForBlocks(latestBlock)
	wg.Wait()
	fmt.Println("Start() is exiting")
}

// TODO (ccdle12): Should be returning (string, error)
func (b *BitcoinCrawler) crawl(blockHash string) string {
	nextBlockHash := blockHash
	for {
		res, err := b.BTCClient.GetBlock(nextBlockHash)
		if err != nil {
			// TODO (ccdle12): Should be returning string, err
			fmt.Println(err.Error())
		}
		// fmt.Println(res)
		bitcoinBlockResponse := b.createBlock(res)

		// TODO (ccdle12):
		// 1. Write the block to the DB
		if insertBlockError := b.BlockTable.InsertBlock(bitcoinBlockResponse.Result); insertBlockError != nil {
			fmt.Println("Insert Block Error: ", err.Error())
		}

		// 2. Request each tx in the block
		// 3. Write each tx to the DB
		// 4. Write each sender and receiver to the DB

		// fmt.Println("Block as struct: ", block)
		// fmt.Println("Next BlockHash: ", block.Result.NextBlockHash)

		// TODO (ccdle12):
		// Should not write the latest block?
		if len(bitcoinBlockResponse.Result.NextBlockHash) == 0 {
			fmt.Println("LATEST BLOCK: ", nextBlockHash)
			fmt.Println("NO MORE BLOCKS")
			return nextBlockHash
		}

		nextBlockHash = bitcoinBlockResponse.Result.NextBlockHash
	}
}

func (b *BitcoinCrawler) createBlock(res string) BitcoinBlockResponse {
	var resBlock BitcoinBlockResponse
	json.Unmarshal([]byte(res), &resBlock)

	return resBlock
}

func (b *BitcoinCrawler) pollForBlocks(blockHash string) {
	fmt.Println("Poll for blocks")
	ticker := time.NewTicker(1 * time.Minute)
	quit := make(chan struct{})

	latestBlock := blockHash
	for {
		select {
		case <-ticker.C:
			latestBlock := b.crawl(latestBlock)
			fmt.Println("LatestBlock from polling: ", latestBlock)
		case <-quit:
			ticker.Stop()
			wg.Done()
			return
		}
	}
}
