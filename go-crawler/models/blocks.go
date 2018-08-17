package models

type Block struct {
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
	Difficulty    float64  `json:"difficulty"`
	Chainwork     string   `json:"chainwork"`
	NextBlockHash string   `json:"nextblockhash"`
}
