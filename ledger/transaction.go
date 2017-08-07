package ledger

const TRANSACTION_VERSION = "0001"

// Output is created by every bitcoin transaction. The are recorded on the
// bitcoin ledger. Almost all of these outputs create spendable chunks of
// bitcoin called unspent transaction outputs or UTXO, which are then
// recognized by the whole network and available for the owner to spend in a
// future transaction
type Output struct {
	// Amount of bitcoin, denominated in satoshis, the smallest bitcoin unit
	Amount int
	// Script (locking script), also known as an "encumbrance" that
	// "locks" this amount by specifying the conditions that must be met
	// to spend the output
	Script string
}

// Input is in simple terms, are pointers to UTXO. They point to a specific
// UTXO by reference to the transaction hash and sequence number where the
// UTXO is recorded in the blockchain. To spend UTXO, a transaction input also
// includes unlocking scripts that satisfy the spending conditions set by the
// UTXO. The unlocking script is usually a signature proving ownership of the
// bitcoin address that is in the locking script.
type Input struct {
	// Pointer to the transaction containing the UTXO to be spent
	TransactionHash string
	// The index number of the UTXO to be spent; first one is 0
	OutputIndex int
	// A script that fulfills the conditions of the UTXO locking script.
	Script string
}

// Transaction is a data structure that encodes a transfer of value from a
// source of funds, called an input, to a destination, called an output.
// Transaction inputs and outputs are not related to accounts or identities.
// Instead, you should think of them as bitcoin amounts—chunks of
// bitcoin—being locked with a specific secret that only the owner, or person
// who knows the secret, can unlock
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
