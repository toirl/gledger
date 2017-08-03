package ledger

import "strconv"
import "strings"
import "time"
import "crypto/sha256"
import "encoding/hex"

const BLOCK_VERSION = "1"
const GENESIS_PAYLOAD = "20170801 Genesis Block. NY TIMES: In Mosul, Revealing the Last ISIS Stronghold"
const GENESIS_ADDRESS = "000000"

type Block struct {
	Version   string
	Payload   string
	Address   string
	Index     int
	timestamp time.Time
}

// GenerateAddress will generate the address for the current block. The
// address is build from a content of the block an the parent block in the
// blockchain.
func (block *Block) GenerateAddress(parent Block) {
	var address_source []string
	address_source = append(address_source, strconv.Itoa(block.Index))
	address_source = append(address_source, block.timestamp.String())
	address_source = append(address_source, parent.Address)
	// address_source.= address_source.append(address_source, strconv(block.Payload))
	block.Address = doubleSHA256([]byte(strings.Join(address_source, "")))
}

//Return a Sha256 Hash of given data
func hash256(data []byte) []byte {
	hash := sha256.New()
	hash.Write(data)
	return hash.Sum(nil)
}

//Returns a string representation of doubled-hashed data
func doubleSHA256(data []byte) string {
	hash := hash256(hash256(data))
	return hex.EncodeToString(hash)
}

func NewBlock(payload string) Block {
	block := Block{}
	block.Version = BLOCK_VERSION
	block.Payload = payload
	block.timestamp = time.Now()

	// The Address of the block is initially empty. It gets generated once
	// the block is finally added to the blockchain.
	block.Address = ""
	// The Index of the block is initially 0. It gets set once
	// the block is finally added to the blockchain.
	block.Index = 0

	return block
}

func GenesisBlock() Block {
	block := NewBlock(GENESIS_PAYLOAD)
	block.Address = GENESIS_ADDRESS
	block.Index = 1
	return block
}
