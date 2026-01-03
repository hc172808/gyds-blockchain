package core

type Blockchain struct {
	Blocks []*Block
}

// NewBlockchain initializes a new chain with genesis block
func NewBlockchain() *Blockchain {
	return &Blockchain{
		Blocks: []*Block{NewGenesisBlock()},
	}
}
