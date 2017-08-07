package ledger

import "crypto/rand"
import "crypto/rsa"

const NBYTES_SECRET = 2048

func NewKey() *rsa.PrivateKey {
	reader := rand.Reader
	bitSize := NBYTES_SECRET
	key, _ := rsa.GenerateKey(reader, bitSize)
	return key
}

// Wallet contains a collection of key pairs, each consisting of a private key
// and a public key.
type Wallet struct {
	Keys []rsa.PrivateKey
}

func NewWallet() Wallet {
	wallet := Wallet{}
	// Generate 5 Keys
	for i := 1; i <= 1; i++ {
		wallet.Keys = append(wallet.Keys, *NewKey())
	}
	return wallet
}
