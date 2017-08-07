package ledger

import "testing"

func TestNewWallet(t *testing.T) {
	NewWallet()
}

func TestNewPrivateKey(t *testing.T) {
	key := NewKey()
	t.Log("Generated Key:", key)
}
