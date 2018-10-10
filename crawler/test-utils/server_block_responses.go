package testutils

// The following variables are used as mock responses for Block requests from the Node.
var (
	NodeResCorrectBlockNoTx0 = `
		{"result": {
			"hash": "0000000000000000001ca03d9e1dd30d2cf49e44ba1569c8819e56cef88b67d4",
			"confirmations": 3,
			"strippedsize": 194985,
			"size": 329911,
			"weight": 914866,
			"height": 538770,
			"version": 536870912,
			"versionHex": "20000000",
			"merkleroot": "d4ffdec00fec5c29ceb948a70c250d07e2254c67ac93b31adc977a8dde9b046c",
			"tx": [],
			"time": 1535396467,
			"mediantime": 1535394491,
			"nonce": 2294488336,
			"bits": "1729d72d",
			"difficulty": 6727225469722.534,
			"chainwork": "000000000000000000000000000000000000000002eb51495ec06b0a5427f048",
			"previousblockhash": "000000000000000000025bfdceefbf04269011ce649538fa07dcbc189f5e3eeb",
			"nextblockhash": "00000000000000000011187b5694d487815d32a4e3e4ac7b5fba28d2cc47df2a"
			},
		"error": null,
		"id": null
		}
	`

	NodeResCorrectBlockNoTx1 = `
	{"result": {
		{"hash": "0000000000000000001ca03d9e1dd30d2cf49e44ba1569c8819e56cef88b67d4",
			"confirmations": 3,
			"strippedsize": 194985,
			"size": 329911,
			"weight": 914866,
			"height": 538770,
			"version": 536870912,
			"versionHex": "20000000",
			"merkleroot": "d4ffdec00fec5c29ceb948a70c250d07e2254c67ac93b31adc977a8dde9b046c",
			"tx": [],
			"time": 1535396467,
			"mediantime": 1535394491,
			"nonce": 2294488336,
			"bits": "1729d72d",
			"difficulty": 6727225469722.534,
			"chainwork": "000000000000000000000000000000000000000002eb51495ec06b0a5427f048",
			"previousblockhash": "000000000000000000025bfdceefbf04269011ce649538fa07dcbc189f5e3eeb",
			"nextblockhash": "00000000000000000011187b5694d487815d32a4e3e4ac7b5fba28d2cc47df2a"
			},
		"error": null,
		"id": null
		}
	`

	MalformedBlockNoTx0 = `
		"hashh""0000000000000000001ca03d9e1dd30d2cf49e44ba1569c8819e56cef88b67d4",
		"confirmations": 3,
		"strippedsize": 194985,
		"size": 329911,
		"weight": 914866,
		"height": 538770,
		"version": 536870912,
		"versionHex": "20000000",
		"merkleroot": "d4ffdec00fec5c29ceb948a70c250d07e2254c67ac93b31adc977a8dde9b046c",
		"tx": [],
		"time": 1535396467,
		"mediantime": 1535394491,
		"nonce": 2294488336,
		"bits": "1729d72d",
		"difficulty": 6727225469722.534,
		"chainwork": "000000000000000000000000000000000000000002eb51495ec06b0a5427f048",
		"previousblockhash": "000000000000000000025bfdceefbf04269011ce649538fa07dcbc189f5e3eeb",
		"nextblockhash": "00000000000000000011187b5694d487815d32a4e3e4ac7b5fba28d2cc47df2a"}`
)
