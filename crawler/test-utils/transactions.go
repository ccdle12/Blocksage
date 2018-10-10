package testutils

import (
	"github.com/ccdle12/Blocksage/crawler/models"
)

var (
	SampleTX = &models.Transaction{
		Blockhash: "",
		TXID:      "1",
		Hash:      "123",
		Version:   1,
		Size:      1,
		Vsize:     1,
		Locktime:  1,
		Vin:       nil,
		Vout:      nil,
	}
)
