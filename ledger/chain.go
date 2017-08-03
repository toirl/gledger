package ledger

type Chain struct {
	version string
	blocks  []Block
}

// Add will add the given block to the blockchain.
func (chain *Chain) Add(block Block) {
	// If the new block added to the chain is not the genesisblock (chain
	// is empty) then take the last block and use it to generate the new
	// blocks address.
	if len(chain.blocks) > 0 {
		last := chain.last()
		block.Index = len(chain.blocks) + 1
		block.GenerateAddress(last)
	}
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
