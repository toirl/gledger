package ledger

import "strconv"
import "strings"
import "time"
import "crypto/sha256"
import "encoding/hex"

const BLOCK_VERSION = "0001"
const GENESIS_PAYLOAD = "20170801 Genesis Block. NY TIMES: In Mosul, Revealing the Last ISIS Stronghold"
const GENESIS_ADDRESS = "000000"

// BlockHeader consists of three sets of block metadata. First, there is
// a reference to a previous block hash, which connects this block to the
// previous block in the blockchain. The second set of metadata, namely the
// difficulty, timestamp, and nonce, relate to the mining competition. The
// third piece of metadata is the merkle tree root, a data structure used to
// efficiently summarize all the transactions in the block.
type BlockHeader struct {
	Version    string
	Parent     string
	Merkle     string
	Timestamp  time.Time
	Difficulty int
	Nonce      int
}

// Block is a container data structure that aggregates transactions for
// inclusion in the public ledger, the blockchain. The block is made of a header,
// containing metadata, followed by a long list of transactions that make up the
// bulk of its size. The block header is 80 bytes, whereas the average
// transaction is at least 250 bytes and the average block contains more than 500
// transactions. A complete block, with all transactions, is therefore 1,000
// times larger than the block header.
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

func NewBlockHeader() BlockHeader {
	header := BlockHeader{}
	header.Version = BLOCK_VERSION
	header.Timestamp = time.Now()
	header.Parent = ""
	header.Merkle = ""
	header.Difficulty = 0
	header.Nonce = 0
	return header
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
