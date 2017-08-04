package ledger

import "testing"
import "fmt"

func TestGensisBlock(t *testing.T) {
	block := GenesisBlock()
	if block.Header.Version != BLOCK_VERSION {
		t.Error(fmt.Errorf("Expected Payload '%s', got %s ", BLOCK_VERSION, block.Header.Version))
	}
}

func TestNewBlock(t *testing.T) {
	block := NewBlock("TX")
	if block.Header.Version != BLOCK_VERSION {
		t.Error(fmt.Errorf("Expected Payload '%s', got %s ", BLOCK_VERSION, block.Header.Version))
	}
}

func TestNewBlockHeader(t *testing.T) {
	header := NewBlockHeader()
	if header.Version != BLOCK_VERSION {
		t.Error(fmt.Errorf("Expected Payload '%s', got %s ", BLOCK_VERSION, header.Version))
	}
}
