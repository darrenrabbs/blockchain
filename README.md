# Simple Blockchain Implementation in Go

This repository contains a simple implementation of a blockchain written in Go. The primary goal is to understand how blockchains work, including creating blocks, adding transactions, and implementing a proof-of-work consensus algorithm.

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
    git clone https://github.com/yourusername/blockChainScratch.git
    cd blockChainScratch
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