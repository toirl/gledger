package ledger

import "testing"
import "strconv"

func TestNewChain(t *testing.T) {
	chain := NewChain()
	genesisBlock := chain.last()
	if genesisBlock.Payload != GENESIS_PAYLOAD {
		t.Error("Expected '%s' as Payload in GenesisBlock, got '%s'", GENESIS_PAYLOAD, genesisBlock.Payload)
	}
}

func TestBuildChain(t *testing.T) {
	chain := NewChain()
	for i := 1; i <= 1000; i++ {
		block := NewBlock(strconv.Itoa(i))
		chain.add(block)
	}
	if len(chain.blocks) != 1001 {
		t.Error("Expected len of chain to be 1001, got '%s'", len(chain.blocks))
	}
}
