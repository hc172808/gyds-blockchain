package main

import (
	"fmt"
	"gyds-blockchain/core"
	"gyds-blockchain/node"
)

func main() {
	fmt.Println("Starting GYDS blockchain node...")

	// Initialize blockchain
	chain := core.NewBlockchain()

	// Start node network listener
	node.StartFullNodeListener(chain)

	fmt.Println("GYDS node running!")
}
