package main

import "os"
import "path"
import "fmt"
import "flag"
import "github.com/toirl/gledger/ledger"

const PATH_CONFIG = ".gledger"
const PATH_WALLET = "wallet.dat"

var wallet = flag.String("wallet", path.Join(os.Getenv("HOME"), PATH_CONFIG, PATH_WALLET), "Path to the wallet file.")

func main() {
	flag.Parse()
	fmt.Println("Welcome to Gledger! Nothing functional yet here.")
	wallet := ledger.GetWallet(*wallet)
	fmt.Println("Keys:", len(wallet.Keys))
}
