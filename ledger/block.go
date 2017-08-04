package ledger

import "time"

const BLOCK_VERSION = "0001"
const GENESIS_PAYLOAD = "20170801 Genesis Block. NY TIMES: In Mosul, Revealing the Last ISIS Stronghold"
const GENESIS_ADDRESS = "000000"

///////////////////////////////////////////////////////////////////////////////
//                                Blockheader                                //
///////////////////////////////////////////////////////////////////////////////

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

///////////////////////////////////////////////////////////////////////////////
//                                   Block                                   //
///////////////////////////////////////////////////////////////////////////////

// Block is a container data structure that aggregates transactions for
// inclusion in the public ledger, the blockchain. The block is made of a header,
// containing metadata, followed by a long list of transactions that make up the
// bulk of its size. The block header is 80 bytes, whereas the average
// transaction is at least 250 bytes and the average block contains more than 500
// transactions. A complete block, with all transactions, is therefore 1,000
// times larger than the block header.
type Block struct {
	Header    BlockHeader
	TxCounter int
	Tx        []int
}

func NewBlock(payload string) Block {
	block := Block{}
	block.Header = NewBlockHeader()
	return block
}

// GenesisBlock returns the first Block in the blockchain.
// It was created in 2017. It is the common ancestor of all the blocks in the
// blockchain, meaning that if you start at any block and follow the chain
// backward in time, you will eventually arrive at the genesis block.
// Data of this block are statically encoded into the code.
func GenesisBlock() Block {
	block := NewBlock(GENESIS_PAYLOAD)
	return block
}
