package core

type Block struct {
	Height       int
	PreviousHash string
	Hash         string
	Transactions []*Transaction
	Validator    string
}

// NewGenesisBlock creates the first block
func NewGenesisBlock() *Block {
	return &Block{
		Height:       0,
		PreviousHash: "",
		Hash:         "GENESIS_HASH",
		Transactions: []*Transaction{},
		Validator:    "GENESIS",
	}
}
