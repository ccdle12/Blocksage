package testutils

// The following variables are used as mock responses for errors returned from the Node.
var (
	NodeResErrorMethodNotFound = `{
    "result": null,
    "error": {
        "code": -32601,
        "message": "Method not found"
    },
    "id": null
	}`

	NodeResNoBlockParams = `{
	"result": null,
	"error": {
		"code": -1,
		"message": "getblock \"blockhash\" ( verbosity ) \n\nIf verbosity is 0, returns a string that is serialized, hex-encoded data for block 'hash'.\nIf verbosity is 1, returns an Object with information about block <hash>.\nIf verbosity is 2, returns an Object with information about block <hash> and information about each transaction. \n\nArguments:\n1. \"blockhash\"          (string, required) The block hash\n2. verbosity              (numeric, optional, default=1) 0 for hex encoded data, 1 for a json object, and 2 for json object with transaction data\n\nResult (for verbosity = 0):\n\"data\"             (string) A string that is serialized, hex-encoded data for block 'hash'.\n\nResult (for verbosity = 1):\n{\n  \"hash\" : \"hash\",     (string) the block hash (same as provided)\n  \"confirmations\" : n,   (numeric) The number of confirmations, or -1 if the block is not on the main chain\n  \"size\" : n,            (numeric) The block size\n  \"strippedsize\" : n,    (numeric) The block size excluding witness data\n  \"weight\" : n           (numeric) The block weight as defined in BIP 141\n  \"height\" : n,          (numeric) The block height or index\n  \"version\" : n,         (numeric) The block version\n  \"versionHex\" : \"00000000\", (string) The block version formatted in hexadecimal\n  \"merkleroot\" : \"xxxx\", (string) The merkle root\n  \"tx\" : [               (array of string) The transaction ids\n     \"transactionid\"     (string) The transaction id\n     ,...\n  ],\n  \"time\" : ttt,          (numeric) The block time in seconds since epoch (Jan 1 1970 GMT)\n  \"mediantime\" : ttt,    (numeric) The median block time in seconds since epoch (Jan 1 1970 GMT)\n  \"nonce\" : n,           (numeric) The nonce\n  \"bits\" : \"1d00ffff\", (string) The bits\n  \"difficulty\" : x.xxx,  (numeric) The difficulty\n  \"chainwork\" : \"xxxx\",  (string) Expected number of hashes required to produce the chain up to this block (in hex)\n  \"previousblockhash\" : \"hash\",  (string) The hash of the previous block\n  \"nextblockhash\" : \"hash\"       (string) The hash of the next block\n}\n\nResult (for verbosity = 2):\n{\n  ...,                     Same output as verbosity = 1.\n  \"tx\" : [               (array of Objects) The transactions in the format of the getrawtransaction RPC. Different from verbosity = 1 \"tx\" result.\n         ,...\n  ],\n  ,...                     Same output as verbosity = 1.\n}\n\nExamples:\n> bitcoin-cli getblock \"00000000c937983704a73af28acdec37b049d214adbda81d7e2a3dd146f6ed09\"\n> curl --user myusername --data-binary '{\"jsonrpc\": \"1.0\", \"id\":\"curltest\", \"method\": \"getblock\", \"params\": [\"00000000c937983704a73af28acdec37b049d214adbda81d7e2a3dd146f6ed09\"] }' -H 'content-type: text/plain;' http://127.0.0.1:8332/\n"
	},
	"id": null
	}`
)
