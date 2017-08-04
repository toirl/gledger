package ledger

type Chain struct {
	version string
	blocks  []Block
}

// Add will add the given block to the blockchain.
func (chain *Chain) Add(block Block) {
	chain.blocks = append(chain.blocks, block)
}

func (chain Chain) last() Block {
	return chain.blocks[len(chain.blocks)-1]
}

// NewChain creates a new blockchain with the first intial GenesisBlock
func NewChain() Chain {
	chain := Chain{}
	genesisBlock := GenesisBlock()
	chain.Add(genesisBlock)
	return chain
}
