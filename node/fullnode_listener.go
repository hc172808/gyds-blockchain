package node

import (
	"fmt"
	"gyds-blockchain/core"
)

func StartFullNodeListener(chain *core.Blockchain) {
	fmt.Println("Full node listener started...")
	// TODO: implement P2P listener for incoming block/tx propagation
}
