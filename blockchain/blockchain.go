package blockchain

// Blockchain is a series of validated Blocks
type Blockchain struct {
	Blocks []*Block // List of blocks in the chain
}

// Transaction represents a data transfer in the blockchain
type Transaction struct {
	Sender   string  // Sender's address
	Receiver string  // Receiver's address
	Amount   float64 // Amount to be transferred
	Coinbase bool    // Indicator if this is a coinbase transaction
}

// NewBlockchain creates a new Blockchain with a genesis block
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{GenesisBlock()}}
}

// AddBlock adds a new block to the blockchain with the provided data and transactions
func (chain *Blockchain) AddBlock(data string, coinbaseReceiver string, transactions []*Transaction) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]

	coinbaseTx := &Transaction{
		Sender:   "Coinbase",
		Receiver: coinbaseReceiver,
		Amount:   10.0,
		Coinbase: true,
	}

	newBlock := NewBlock(data, prevBlock.Hash, append([]*Transaction{coinbaseTx}, transactions...))
	chain.Blocks = append(chain.Blocks, newBlock)
}
