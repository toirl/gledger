package ledger

import "testing"
import "fmt"

func TestNewTransaction(t *testing.T) {
	tx := NewTransaction()
	if tx.Version != TRANSACTION_VERSION {
		t.Error(fmt.Errorf("Expected Version '%s', got %s ", TRANSACTION_VERSION, tx.Version))
	}
}
