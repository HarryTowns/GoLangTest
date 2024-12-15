package main

import (
	"fmt"
	"simple-blockchain/mypackage"
)

func main() {
	// Initialize a new blockchain
	blockchain := mypackage.InitializeBlockchain()

	// Add blocks to the blockchain
	blockchain.AddBlock("Alice pays Bob 5 BTC")
	blockchain.AddBlock("Bob pays Charlie 3 BTC")
	blockchain.AddBlock("Charlie pays Dave 1 BTC")

	// Display the blockchain
	fmt.Println("Blockchain:")
	for _, block := range blockchain.Blocks {
		fmt.Printf("\nIndex: %d", block.Index)
		fmt.Printf("\nTimestamp: %s", block.Timestamp)
		fmt.Printf("\nData: %s", block.Data)
		fmt.Printf("\nPrevious Hash: %s", block.PreviousHash)
		fmt.Printf("\nHash: %s", block.Hash)
		fmt.Printf("\nNonce: %d", block.Nonce)
		fmt.Println("\n----------------------")
	}
}
