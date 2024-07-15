package blockchain

import (
	"testing"
)

func TestNewBlock(t *testing.T) {
	transactions := []*Transaction{
		{Sender: "Alice", Receiver: "Bob", Amount: 10.0},
	}
	block := NewBlock("Test Data", "PreviousHash", transactions)

	if block.Data != "Test Data" {
		t.Errorf("Expected Data to be 'Test Data', got %s", block.Data)
	}
	if block.PrevHash != "PreviousHash" {
		t.Errorf("Expected PrevHash to be 'PreviousHash', got %s", block.PrevHash)
	}
	if block.Transactions[0].Sender != "Alice" {
		t.Errorf("Expected Sender to be 'Alice', got %s", block.Transactions[0].Sender)
	}
}

func TestGenesisBlock(t *testing.T) {
	block := GenesisBlock()

	if block.Data != "Genesis Block" {
		t.Errorf("Expected Data to be 'Genesis Block', got %s", block.Data)
	}
	if block.PrevHash != "" {
		t.Errorf("Expected PrevHash to be empty, got %s", block.PrevHash)
	}
	if !block.Transactions[0].Coinbase {
		t.Errorf("Expected Coinbase to be true, got %t", block.Transactions[0].Coinbase)
	}
}
