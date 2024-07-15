package blockchain

import (
	"testing"
)

func TestNewBlockchain(t *testing.T) {
	chain := NewBlockchain()

	if len(chain.Blocks) != 1 {
		t.Errorf("Expected blockchain length to be 1, got %d", len(chain.Blocks))
	}
	if chain.Blocks[0].Data != "Genesis Block" {
		t.Errorf("Expected Genesis block data to be 'Genesis Block', got %s", chain.Blocks[0].Data)
	}
}

func TestAddBlock(t *testing.T) {
	chain := NewBlockchain()
	transactions := []*Transaction{
		{Sender: "Alice", Receiver: "Bob", Amount: 5.0},
	}
	chain.AddBlock("Block 1", "Alice", transactions)

	if len(chain.Blocks) != 2 {
		t.Errorf("Expected blockchain length to be 2, got %d", len(chain.Blocks))
	}
	if chain.Blocks[1].Data != "Block 1" {
		t.Errorf("Expected block data to be 'Block 1', got %s", chain.Blocks[1].Data)
	}
	if chain.Blocks[1].Transactions[1].Sender != "Alice" {
		t.Errorf("Expected transaction sender to be 'Alice', got %s", chain.Blocks[1].Transactions[1].Sender)
	}
}
