package blockchain

import (
	"math/rand"
	"time"
)

// Block represents each 'item' in the blockchain
type Block struct {
	Hash         string         // Hash of the block
	Data         string         // Data inside the block
	PrevHash     string         // Hash of the previous block
	Nonce        int            // A random value to ensure uniqueness
	Transactions []*Transaction // List of transactions in the block
}

// NewBlock creates a new block with given data and previous block's hash
func NewBlock(data string, prevHash string, transactions []*Transaction) *Block {
	rand.Seed(time.Now().UnixNano())
	initialNonce := rand.Intn(10000)

	block := &Block{"", data, prevHash, initialNonce, transactions}

	pow := NewProofOfWork(block)
	nonce, hash := pow.MineBlock()

	block.Hash = string(hash[:])
	block.Nonce = nonce

	return block
}

// GenesisBlock creates the first block in the blockchain
func GenesisBlock() *Block {
	coinbaseTx := &Transaction{
		Sender:   "Coinbase",
		Receiver: "Genesis",
		Amount:   0.0,
		Coinbase: true,
	}

	return NewBlock("Genesis Block", "", []*Transaction{coinbaseTx})
}
