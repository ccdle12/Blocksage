package bitcoincrawler

import (
	"encoding/json"
	"fmt"
	"github.com/ccdle12/Blocksage/go-api/bitcoinclient"
	"net/http"
	"os"
	"time"
)

// BitcoinCrawler is a struct that will crawl retrieve data from the Blockchain
type BitcoinCrawler struct{}

// BitcoinBlockResponse inherits from RPCBitcoinResponse, the block struct overwrites the Result field in RPCBitcoinResponse
// TODO: Separate the structs into a model package?
type BitcoinBlockResponse struct {
	RPCBitcoinResponse bitcoinclient.RPCBitcoinResponse
	Result             block `json:"result"`
}

type block struct {
	Hash          string   `json:"hash"`
	Confirmations int      `json:"confirmations"`
	Strippedsize  int      `json:"strippedsize"`
	Size          int      `json:"size"`
	Weight        int      `json:"weight"`
	Height        int      `json:"height"`
	Version       int      `json:"version"`
	VersionHex    string   `json:"versionHex"`
	MerkleRoot    string   `json:"merkleroot"`
	TX            []string `json:"tx"`
	Time          int      `json:"time"`
	MedianTime    int      `json:"mediantime"`
	Nonce         int      `json:"nonce"`
	Bits          string   `json:"bits"`
	Difficulty    int      `json:"difficulty"`
	Chainwork     string   `json:"chainwork"`
	NextBlockHash string   `json:"nextblockhash"`
}

var (
	btcMainDomain = os.Getenv("BTC_MAIN_DOMAIN")
	bitcoinClient = &bitcoinclient.BitcoinClient{
		Client:          &http.Client{Timeout: time.Duration(5 * time.Second)},
		BitcoinNodeAddr: fmt.Sprintf("http://%s:8332", btcMainDomain),
	}
	genesisBlockHash = "000000000019d6689c085ae165831e934ff763ae46a2a6c172b3f1b60a8ce26f"
	recentBlock      = "000000000000000000071629bba3984b668c31b3b599b3ad3c4de2341fb2c874"
)

// Start will run the crawler
func (b *BitcoinCrawler) Start() {
	latestBlock := b.crawl(recentBlock)
	b.pollForBlocks(latestBlock)
}

func (b *BitcoinCrawler) crawl(blockHash string) string {
	nextBlockHash := blockHash
	for {
		res, err := bitcoinClient.GetBlock(nextBlockHash)
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Printf("Response on: %s \n", res)

		block := b.createBlock(res)

		fmt.Println("Block as struct: ", block)
		fmt.Println("Next BlockHash: ", block.Result.NextBlockHash)

		if len(block.Result.NextBlockHash) == 0 {
			fmt.Println("LATEST BLOCK: ", nextBlockHash)
			return nextBlockHash
		}
		nextBlockHash = block.Result.NextBlockHash
	}
}

func (b *BitcoinCrawler) createBlock(res string) BitcoinBlockResponse {
	var resBlock BitcoinBlockResponse
	json.Unmarshal([]byte(res), &resBlock)

	return resBlock
}

func (b *BitcoinCrawler) pollForBlocks(blockHash string) {
	ticker := time.NewTicker(1 * time.Minute)
	quit := make(chan struct{})
	go func() {
		latestBlock := blockHash
		for {
			select {
			case <-ticker.C:
				latestBlock := b.crawl(latestBlock)
				fmt.Println("LatestBlock from polling: ", latestBlock)
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}
