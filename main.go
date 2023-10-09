package main

import (
	a1 "github.com/KisaaBatool05/assignment01bca_project"
)

func main() {
	// Create a new blockchain
	blockchain := &a1.Blockchain{}

	// Add blocks to the blockchain
	blockchain.NewBlock("Bob to Alice", 123, "Genesis Block")
	blockchain.NewBlock("Alice to Carol", 456, blockchain.Blocks[len(blockchain.Blocks)-1].CurrentHash)

	// Display all blocks in the blockchain
	blockchain.DisplayBlocks()
}
