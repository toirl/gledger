package ledger

import "os"
import "encoding/gob"
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

func (wallet *Wallet) Save(path string) error {
	file, err := os.Create(path)
	if err == nil {
		encoder := gob.NewEncoder(file)
		encoder.Encode(wallet)
	}
	file.Close()
	return err
}

func (wallet *Wallet) Load(path string) error {
	file, err := os.Open(path)
	if err == nil {
		decoder := gob.NewDecoder(file)
		err = decoder.Decode(wallet)
	}
	file.Close()
	return err
}

func GetWallet(path string) Wallet {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// path to config does not exist
		wallet := NewWallet()
		wallet.Save(path)
		return wallet
	} else {
		wallet := new(Wallet)
		wallet.Load(path)
		return *wallet
	}
}

func NewWallet() Wallet {
	wallet := Wallet{}
	// Generate 5 Keys
	for i := 1; i <= 1; i++ {
		wallet.Keys = append(wallet.Keys, *NewKey())
	}
	return wallet
}
