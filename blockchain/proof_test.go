package blockchain

import (
	"testing"
)

func TestProofOfWork_Validate(t *testing.T) {
	block := NewBlock("Test Data", "PreviousHash", nil)
	pow := NewProofOfWork(block)
	nonce, hash := pow.MineBlock()

	block.Nonce = nonce
	block.Hash = string(hash[:])

	if !pow.Validate() {
		t.Errorf("Expected ProofOfWork to be valid")
	}
}
