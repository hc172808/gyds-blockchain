package core

type Account struct {
	Address  string
	Nonce    uint64
	Balances map[string]uint64 // assetID -> amount
	Stake    uint64            // For PoS validators
}
