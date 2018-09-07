package testutils

import (
	"github.com/ccdle12/Blocksage/go-crawler/models"
)

var (
	Block0 = &models.Block{
		Hash:              "0000000000000000001ca03d9e1dd30d2cf49e44ba1569c8819e56cef88b67d4",
		Strippedsize:      194985,
		Size:              329911,
		Weight:            914866,
		Height:            538770,
		Version:           536870912,
		VersionHex:        "20000000",
		MerkleRoot:        "d4ffdec00fec5c29ceb948a70c250d07e2254c67ac93b31adc977a8dde9b046c",
		TX:                []string{},
		Time:              1535396467,
		MedianTime:        1535394491,
		Nonce:             2294488336,
		Bits:              "1729d72d",
		Difficulty:        6727225469722.534,
		Chainwork:         "000000000000000000000000000000000000000002eb51495ec06b0a5427f048",
		PreviousBlockHash: "000000000000000000025bfdceefbf04269011ce649538fa07dcbc189f5e3eeb",
		NextBlockHash:     "00000000000000000011187b5694d487815d32a4e3e4ac7b5fba28d2cc47df2a",
	}
)
