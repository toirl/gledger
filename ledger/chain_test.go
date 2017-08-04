package ledger

import "testing"
import "fmt"
import "strconv"

func TestNewChain(t *testing.T) {
	chain := NewChain()
	genesisBlock := chain.last()
	if genesisBlock.Header.Version != BLOCK_VERSION {
		t.Error("Expected Payload '%s', got %s ", BLOCK_VERSION, genesisBlock.Header.Version)
	}
}

func TestBuildChain(t *testing.T) {
	chain := NewChain()
	for i := 1; i <= 1000; i++ {
		block := NewBlock(strconv.Itoa(i))
		chain.Add(block)
	}
	if len(chain.blocks) != 1001 {
		t.Error(fmt.Errorf("Expected len of chain to be 1001, got %d", len(chain.blocks)))
	}
}
