package ledger

import "testing"
import "fmt"

func TestGensisBlock(t *testing.T) {
	block := GenesisBlock()
	if block.Payload != GENESIS_PAYLOAD {
		t.Error("Expected Payload '%s', got %s ", GENESIS_PAYLOAD, block.Payload)
	}
}

func TestNewBlock(t *testing.T) {
	var want_version string = "0001"
	var want_payload string = "Foo"
	block := NewBlock(want_payload)
	if block.Version != want_version {
		t.Error(fmt.Errorf("Expected Version %s, got %s", want_version, block.Version))
	}
	if block.Payload != want_payload {
		t.Error(fmt.Errorf("Expected Version %s, got %s", want_payload, block.Payload))
	}
}

func TestNewBlockHeader(t *testing.T) {
	header := NewBlockHeader()
	want_version := "0001"
	if header.Version != want_version {
		t.Error(fmt.Errorf("Expected Version %s, got %s", want_version, header.Version))
	}
}
