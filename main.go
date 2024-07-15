package main

import (
	"github/com/darrenrabbs/blockchain",
	"fmt"
	"strconv"
)

func main() {
	chain := blockchain.NewBlockchain()

	// Create wallets for users
	aliceWallet, err := blockchain.CreateWallet()
	if err != nil {
		fmt.Println("Error creating Alice's wallet:", err)
		return
	}
	fmt.Println("Alice's wallet successfully created")

	bobWallet, err := blockchain.CreateWallet()
	if err != nil {
		fmt.Println("Error creating Bob's wallet:", err)
		return
	}
	fmt.Println("Bob's wallet successfully created")

	// Create and sign a transaction from Alice to Bob
	tx := &blockchain.Transaction{
		Sender:   aliceWallet.PublicKey.N.String(),
		Receiver: bobWallet.PublicKey.N.String(),
		Amount:   5.0,
	}
	fmt.Println("Transaction from Alice to Bob created successfully")

	signature, err := aliceWallet.SignTransaction(tx)
	if err != nil {
		fmt.Println("Error signing the transaction:", err)
		return
	}

	err = blockchain.VerifyTransaction(tx, aliceWallet.PublicKey, signature)
	if err != nil {
		fmt.Println("Transaction verification failed:", err)
		return
	}
	fmt.Println("Transaction verified successfully")

	// Add the verified transaction to the blockchain
	chain.AddBlock("Block 1", "Alice", []*blockchain.Transaction{tx})
	fmt.Println()

	// Print the blockchain
	for _, block := range chain.Blocks {
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Block Data: %s\n", block.Data)
		fmt.Printf("Block Hash: %x\n", block.Hash)

		pow := blockchain.NewProofOfWork(block)
		fmt.Printf("IsValidPoW: %s\n",
			strconv.FormatBool(pow.Validate()))
		fmt.Println("Transactions:")
		for _, tx := range block.Transactions {
			fmt.Printf("Sender: %s\n", tx.Sender)
			fmt.Printf("Receiver: %s\n", tx.Receiver)
			fmt.Printf("Amount: %f\n", tx.Amount)
			fmt.Printf("Coinbase: %t\n", tx.Coinbase)
			fmt.Println()
		}
		fmt.Println()
	}
}
