package core

type Transaction struct {
	From      string
	To        string
	AssetID   string
	Amount    uint64
	Fee       uint64
	Nonce     uint64
	Signature string
}
