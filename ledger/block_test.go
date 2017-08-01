package ledger

import "testing"

func TestGensisBlock(t *testing.T) {
	block := GenesisBlock()
	if block.Payload != GENESIS_PAYLOAD {
		t.Error("Expected Payload '%s', got %s ", GENESIS_PAYLOAD, block.Payload)
	}
}

func TestNewBlock(t *testing.T) {
	var want_version string = "1"
	var want_payload string = "Foo"
	block := NewBlock(want_payload)
	if block.Version != want_version {
		t.Error("Expected Version %s, got %s", want_version, block.Version)
	}
	if block.Payload != want_payload {
		t.Error("Expected Payload '%s', got %s ", want_payload, block.Payload)
	}
}
