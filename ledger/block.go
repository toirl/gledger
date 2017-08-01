package ledger

const BLOCK_VERSION = "1"
const GENESIS_PAYLOAD = "20170801 Genesis Block. NY TIMES: In Mosul, Revealing the Last ISIS Stronghold"

type Block struct {
	Version string
	Payload string
}

func NewBlock(payload string) Block {
	block := Block{}
	block.Version = BLOCK_VERSION
	block.Payload = payload
	return block
}

func GenesisBlock() Block {
	block := NewBlock(GENESIS_PAYLOAD)
	return block
}
