package blockchain

import (
	"bytes"
	"crypto/md5"
	"encoding/binary"
	"fmt"
	"math/big"
)

// Difficulty defines how hard it is to mine a new block
const Difficulty = 12

// ProofOfWork represents the proof of work algorithm for a block
type ProofOfWork struct {
	Block  *Block
	Target *big.Int // Target hash value to be achieved
}

// NewProofOfWork initializes a ProofOfWork for a given block
func NewProofOfWork(block *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-Difficulty))

	return &ProofOfWork{block, target}
}

// PrepareData combines block attributes and nonce to generate a hash input
func (pow *ProofOfWork) PrepareData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			[]byte(pow.Block.PrevHash),
			[]byte(pow.Block.Data),
			make([]byte, 8),
			make([]byte, 8),
		},
		[]byte{},
	)

	binary.BigEndian.PutUint64(data[len(data)-16:], uint64(nonce))
	binary.BigEndian.PutUint64(data[len(data)-8:], uint64(Difficulty))

	return data
}

// MineBlock performs the mining process to find a valid block hash
func (pow *ProofOfWork) MineBlock() (int, []byte) {
	var intHash big.Int
	var hash [16]byte
	nonce := 0

	for {
		data := pow.PrepareData(nonce)
		hash = md5.Sum(data)

		fmt.Printf("\r%x", hash)
		intHash.SetBytes(hash[:])

		if intHash.Cmp(pow.Target) == -1 {
			break
		} else {
			nonce++
		}
	}
	fmt.Println()

	return nonce, hash[:]
}

// Validate checks if the current block hash is valid according to the proof of work
func (pow *ProofOfWork) Validate() bool {
	var intHash big.Int

	data := pow.PrepareData(pow.Block.Nonce)
	hash := md5.Sum(data)
	intHash.SetBytes(hash[:])

	return intHash.Cmp(pow.Target) == -1
}
