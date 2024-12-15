package mypackage

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"
)

// Block represents a single block in the blockchain
type Block struct {
	Index        int    // Position of the block in the chain
	Timestamp    string // Time of block creation
	Data         string // Transaction data
	PreviousHash string // Hash of the previous block
	Hash         string // Hash of the current block
	Nonce        int    // Nonce for proof of work
}

// Blockchain is a series of connected blocks
type Blockchain struct {
	Blocks []Block
}

// CreateBlock creates a new block and returns it
func CreateBlock(index int, data string, previousHash string) Block {
	block := Block{
		Index:        index,
		Timestamp:    time.Now().String(),
		Data:         data,
		PreviousHash: previousHash,
		Nonce:        0,
	}

	block.Hash = GenerateProofOfWork(&block)
	return block
}

// GenerateProofOfWork applies Proof of Work by finding a hash with a specific pattern
func GenerateProofOfWork(block *Block) string {
	var hash string
	var hashBytes [32]byte

	for {
		data := blockToBytes(block)
		hashBytes = sha256.Sum256(data)
		hash = hex.EncodeToString(hashBytes[:])

		// Simple PoW condition: hash must start with "0000"
		if hash[:4] == "0000" {
			break
		}

		block.Nonce++
	}

	return hash
}

// blockToBytes converts a block's fields to a byte array for hashing
func blockToBytes(block *Block) []byte {
	data := bytes.Join([][]byte{
		[]byte(strconv.Itoa(block.Index)),
		[]byte(block.Timestamp),
		[]byte(block.Data),
		[]byte(block.PreviousHash),
		[]byte(strconv.Itoa(block.Nonce)),
	}, []byte{})

	return data
}

// AddBlock adds a new block to the blockchain
func (bc *Blockchain) AddBlock(data string) {
	previousBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := CreateBlock(previousBlock.Index+1, data, previousBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
}

// CreateGenesisBlock creates the first block in the blockchain
func CreateGenesisBlock() Block {
	return CreateBlock(0, "Genesis Block", "0")
}

// InitializeBlockchain initializes the blockchain with the genesis block
func InitializeBlockchain() *Blockchain {
	genesisBlock := CreateGenesisBlock()
	return &Blockchain{Blocks: []Block{genesisBlock}}
}
