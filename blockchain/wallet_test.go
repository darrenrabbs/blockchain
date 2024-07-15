package blockchain

import (
	"testing"
)

func TestGenerateKeys(t *testing.T) {
	privateKey, publicKey, err := GenerateKeys()
	if err != nil {
		t.Errorf("Failed to generate keys: %v", err)
	}
	if privateKey == nil || publicKey == nil {
		t.Errorf("Expected non-nil keys, got nil")
	}
}

func TestSignAndVerifyTransaction(t *testing.T) {
	wallet, err := CreateWallet()
	if err != nil {
		t.Errorf("Failed to create wallet: %v", err)
	}

	tx := &Transaction{
		Sender:   "Alice",
		Receiver: "Bob",
		Amount:   5.0,
	}

	signature, err := wallet.SignTransaction(tx)
	if err != nil {
		t.Errorf("Failed to sign transaction: %v", err)
	}

	err = VerifyTransaction(tx, wallet.PublicKey, signature)
	if err != nil {
		t.Errorf("Failed to verify transaction: %v", err)
	}
}
