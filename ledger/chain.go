package ledger

import "fmt"

type Chain struct {
	version string
	blocks  []Block
}

func (chain *Chain) add(block Block) {
	chain.blocks = append(chain.blocks, block)
	fmt.Println(len(chain.blocks))
}

func (chain Chain) last() Block {
	fmt.Println(len(chain.blocks))
	return chain.blocks[len(chain.blocks)-1]
}

// NewChain creates a new blockchain with the first intial GenesisBlock
func NewChain() Chain {
	chain := Chain{}
	genesisBlock := GenesisBlock()
	chain.add(genesisBlock)
	return chain
}
