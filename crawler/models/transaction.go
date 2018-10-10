package models

// Transaction is the struct which holds all the details for transactions from the
// Bitcoin Network.
type Transaction struct {
	Blockhash string              `json:"blockhash"`
	TXID      string              `json:"txid"`
	Hash      string              `json:"hash"`
	Version   int                 `json:"version"`
	Size      int                 `json:"size"`
	Vsize     int                 `json:"vsize"`
	Locktime  int                 `json:"locktime"`
	Vin       []TransactionInput  `json:"vin"`
	Vout      []TransactionOutput `json:"vout"`
}

// TransactionInput is the struct which holds all the details for inputs in a transaction.
type TransactionInput struct {
	Txid      string               `json:"txid"`
	Vout      int                  `json:"vout"`
	ScriptSig TransactionScriptSig `json:"scriptSig"`
	Sequence  int64                `json:"sequence"`
}

// TransactionScriptSig holds the details for the unlock script of a transaction.
type TransactionScriptSig struct {
	Asm string `json:"asm"`
	Hex string `json:"hex"`
}

// TransactionOutput holds all the output details for a transaction.
type TransactionOutput struct {
	Value        float64                 `json:"value"`
	N            int                     `json:"n"`
	ScriptPubKey TransactionScriptPubKey `json:"scriptPubKey"`
}

// TransactionScriptPubKey holds all the details for the locking script of a transaction.
type TransactionScriptPubKey struct {
	Asm       string   `json:"asm"`
	Hex       string   `json:"hex"`
	ReqSigs   int      `json:"reqSigs"`
	Type      string   `json:"type"`
	Addresses []string `json:"addresses"`
}
