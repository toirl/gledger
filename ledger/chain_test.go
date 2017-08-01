package ledger

import "testing"

func TestNewChain(t *testing.T) {
	chain := NewChain()
	genesisBlock := chain.last()
	if genesisBlock.Payload != GENESIS_PAYLOAD {
		t.Error("Expected '%s' as Payload in GenesisBlock, got '%s'", GENESIS_PAYLOAD, genesisBlock.Payload)
	}
}
