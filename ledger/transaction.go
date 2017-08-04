package ledger

const TRANSACTION_VERSION = "0001"

type Transaction struct {
	Version       string
	InputCounter  int
	Inputs        string
	OutputCounter int
	Outputs       string
}

func NewTransaction() Transaction {
	tx := Transaction{}
	tx.Version = TRANSACTION_VERSION
	return tx
}
