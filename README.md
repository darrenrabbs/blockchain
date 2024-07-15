# Simple Blockchain Implementation in Go

This repository contains a simple implementation of a blockchain written in Go. The primary goal is to understand how blockchains work, including creating blocks, adding transactions, and implementing a proof-of-work consensus algorithm.

## What is a Blockchain?

A blockchain is a distributed ledger that records transactions in a secure, immutable way. It is composed of a series of blocks, each containing data, a hash of the current block, and a hash of the previous block. This creates a chain of blocks that are cryptographically linked together, ensuring the integrity and security of the data.

### Key Components of a Blockchain:

1. **Block**: Represents a single unit in the blockchain containing data, the current block's hash, and the previous block's hash.
2. **Blockchain**: A chain of blocks linked together.
3. **Transaction**: Represents a data transfer within the blockchain.
4. **Wallet**: Contains cryptographic keys used to sign transactions.
5. **Proof of Work (PoW)**: A consensus algorithm used to validate transactions and add new blocks.

### Diagram of a Blockchain:

```plaintext
+---------+       +---------+       +---------+
| Block 0 |  -->  | Block 1 |  -->  | Block 2 | 
+---------+       +---------+       +---------+
|  Data   |       |  Data   |       |  Data   |
|  Hash0  |       |  Hash1  |       |  Hash2  |
|  Prev:0 |       |  Prev:1 |       |  Prev:2 |
+---------+       +---------+       +---------+
```

## Project Structure

- **blockchain/**: Contains the core blockchain implementation
  - `block.go`: Contains the `Block` struct and functions for creating blocks and the genesis block.
  - `blockchain.go`: Contains the `Blockchain` struct and functions for initializing and adding blocks to the blockchain.
  - `proof.go`: Implements the proof-of-work algorithm used to mine new blocks.
  - `wallet.go`: Manages user wallets, including RSA key generation, signing transactions, and verifying transaction signatures.
  - `block_test.go`: Unit tests for `block.go`.
  - `blockchain_test.go`: Unit tests for `blockchain.go`.
  - `proof_test.go`: Unit tests for `proof.go`.
  - `wallet_test.go`: Unit tests for `wallet.go`.
- `main.go`: Demonstrates the usage of the blockchain, including creating wallets, making transactions, and verifying them.
- `go.mod`: Go module file.
- `Makefile`: Makefile for building, running, and testing the application.

## Running the Project

To run the project, you will need Go installed on your machine. Follow the steps below:

1. Clone the repository:
    ```sh
    git clone https://github.com/darrenrabbs/blockchain.git
    cd blockchain
    ```

2. Run the application:
    ```sh
    make run
    ```

3. Run the tests:
    ```sh
    make test
    ```

## Building the Application

To build the application, use the following command:
```sh
make build
```

## Understanding the Code

### Blockchain

The Blockchain struct manages the chain of blocks. It includes methods to initialize the blockchain with a genesis block and to add new blocks to the chain.

```go
func main() {
    chain := blockchain.NewBlockchain()
    chain.AddBlock("Block 1", "Alice", []*blockchain.Transaction{tx})
}
```


Block 

The Block struct represents each block in the blockchain. It contains the block’s data, its hash, the hash of the previous block, a nonce, and a list of transactions.
```go
type Block struct {
    Hash         string           // Hash of the block
    Data         string           // Data inside the block
    PrevHash     string           // Hash of the previous block
    Nonce        int              // Random value to ensure uniqueness
    Transactions []*Transaction   // List of transactions in the block
}
```

Transaction

The Transaction struct represents a data transfer in the blockchain. It includes the sender’s and receiver’s addresses, the amount, and a flag indicating whether it is a coinbase transaction.

```go
type Transaction struct {
    Sender   string  // Sender's address
    Receiver string  // Receiver's address
    Amount   float64 // Amount to be transferred
    Coinbase bool    // Indicator if this is a coinbase transaction
}
```

Wallet
The Wallet struct manages the user’s cryptographic keys and provides methods to sign and verify transactions.

```go
wallet, err := blockchain.CreateWallet()
signature, err := wallet.SignTransaction(tx)
err = blockchain.VerifyTransaction(tx, wallet.PublicKey, signature)
```

Proof of Work

The ProofOfWork struct implements the proof-of-work algorithm used to mine new blocks. It includes methods to prepare data for hashing, mine a block, and validate a block.

```go
pow := blockchain.NewProofOfWork(block)
nonce, hash := pow.MineBlock()
valid := pow.Validate()
```

## Conclusion
This project provides a simple implementation of a blockchain to help understand the fundamental concepts. It demonstrates how blocks are created, linked together, and validated using a proof-of-work consensus algorithm. Wallets are used to manage cryptographic keys and sign transactions, ensuring secure and verifiable transfers of data within the blockchain.

Enjoy exploring and understanding the basics of blockchain technology with a simple Go implementation!


## Additional References
	•	Books:
	•	“Mastering Bitcoin” by Andreas M. Antonopoulos: An in-depth book that explains the technical workings of Bitcoin and blockchain.
	•	“Blockchain Basics: A Non-Technical Introduction in 25 Steps” by Daniel Drescher: A great book for beginners that explains blockchain technology in simple terms.
	•	Online Courses:
	•	Coursera: Blockchain Specialization: A series of courses offered by the University at Buffalo and SUNY that cover blockchain basics and advanced topics.
	•	edX: Blockchain Fundamentals: A professional certificate program from UC Berkeley that covers the fundamental concepts of blockchain technology.
	•	Websites and Articles:
	•	Bitcoin.org Developer Guide: A comprehensive guide for developers interested in the technical aspects of Bitcoin.
	•	Ethereum Whitepaper: The original whitepaper for Ethereum, detailing its architecture and functionality.
	•	IBM Blockchain 101: Quick-Start Guide: An introductory guide to blockchain technology from IBM.
	•	Videos:
	•	Andreas M. Antonopoulos YouTube Channel: A series of videos by Andreas M. Antonopoulos that cover various aspects of Bitcoin and blockchain technology.
	•	Blockchain at Berkeley YouTube Channel: Educational videos on blockchain from the Blockchain at Berkeley organization.