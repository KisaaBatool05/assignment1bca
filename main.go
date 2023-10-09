package main

import (
	"assignment01bca" // Import your Blockchain package
	"fmt"
)

func main() {
	// Create a new blockchain
	blockchain := &assignment01bca.Blockchain{}

	// Add some blocks to the blockchain
	blockchain.NewBlock("Alice to Bob", 123, "0000000000")
	blockchain.NewBlock("Bob to Cara", 456, blockchain.Blocks[len(blockchain.Blocks)-1].CurrentHash)
	blockchain.NewBlock("Alice to Elve", 789, blockchain.Blocks[len(blockchain.Blocks)-1].CurrentHash)

	// Display all blocks in the blockchain
	blockchain.DisplayBlocks()

	// Change the transaction of a block (e.g., block 2)
	if err := blockchain.ChangeBlock(2, "David to Dave"); err != nil {
		fmt.Println("Error:", err)
	}

	// Verify the integrity of the blockchain
	isValid := blockchain.VerifyChain()
	if isValid {
		fmt.Println("Blockchain is valid.")
	} else {
		fmt.Println("Blockchain is not valid.")
	}

	// Calculate the hash of a block (e.g., the first block)
	hash := assignment01bca.calculateHash(blockchain.Blocks[0].Transaction)
	fmt.Printf("Hash of the first block: %s\n", hash)
}
