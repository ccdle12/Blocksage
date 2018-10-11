package testutils

import (
	"github.com/ccdle12/Blocksage/crawler/models"
)

var (
	// Block538770 is a real world Bitcoin block.
	Block538770 = &models.Block{
		Hash:              "0000000000000000001ca03d9e1dd30d2cf49e44ba1569c8819e56cef88b67d4",
		Strippedsize:      194985,
		Size:              329911,
		Weight:            914866,
		Height:            538770,
		Version:           536870912,
		VersionHex:        "20000000",
		MerkleRoot:        "d4ffdec00fec5c29ceb948a70c250d07e2254c67ac93b31adc977a8dde9b046c",
		TX:                TXBlock538770,
		Time:              1535396467,
		MedianTime:        1535394491,
		Nonce:             2294488336,
		Bits:              "1729d72d",
		Difficulty:        6727225469722.534,
		Chainwork:         "000000000000000000000000000000000000000002eb51495ec06b0a5427f048",
		PreviousBlockHash: "000000000000000000025bfdceefbf04269011ce649538fa07dcbc189f5e3eeb",
		NextBlockHash:     "00000000000000000011187b5694d487815d32a4e3e4ac7b5fba28d2cc47df2a",
	}

	// ReducedBlock538770 is a block with a reduced number of transactions.
	ReducedBlock538770 = &models.Block{
		Hash:              "0000000000000000001ca03d9e1dd30d2cf49e44ba1569c8819e56cef88b67d4",
		Strippedsize:      194985,
		Size:              329911,
		Weight:            914866,
		Height:            538770,
		Version:           536870912,
		VersionHex:        "20000000",
		MerkleRoot:        "d4ffdec00fec5c29ceb948a70c250d07e2254c67ac93b31adc977a8dde9b046c",
		TX:                ReducedTXBlock538770,
		Time:              1535396467,
		MedianTime:        1535394491,
		Nonce:             2294488336,
		Bits:              "1729d72d",
		Difficulty:        6727225469722.534,
		Chainwork:         "000000000000000000000000000000000000000002eb51495ec06b0a5427f048",
		PreviousBlockHash: "000000000000000000025bfdceefbf04269011ce649538fa07dcbc189f5e3eeb",
		NextBlockHash:     "00000000000000000011187b5694d487815d32a4e3e4ac7b5fba28d2cc47df2a",
	}

	Block506664 = &models.Block{
		Hash:              "00000000000000000033096861183add56d76b9fa06d09e84a75ffaeacbcdc1c",
		Strippedsize:      983197,
		Size:              1048468,
		Weight:            3998059,
		Height:            506664,
		Version:           536870912,
		VersionHex:        "20000000",
		MerkleRoot:        "6e9d4a9c1dad90575c9c38c26a45c2ef29b4b4c9a6d226f8eddb1eeed7c485e8",
		TX:                TXBlock506664,
		Time:              1517238427,
		MedianTime:        1517236840,
		Nonce:             3682140350,
		Bits:              "176c2146",
		Difficulty:        2603077300218.593,
		Chainwork:         "000000000000000000000000000000000000000000fea1c09af0ebd2126ba98f",
		PreviousBlockHash: "0000000000000000001db5495639a344ff152a53def7496ea95977b3b72fd43a",
		NextBlockHash:     "000000000000000000443d5a5d0cdf525d9f8e6406567783e85ce2d1b0c5f2f4",
	}
)
