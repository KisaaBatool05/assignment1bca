package assignment01bca

import (
	"crypto/sha256"
	"fmt"
)

// Block represents a block in the blockchain.
type Block struct {
	Transaction  string
	Nonce        int
	PreviousHash string
	CurrentHash  string
}

// Blockchain represents a blockchain.
type Blockchain struct {
	Blocks []*Block
}

// NewBlock creates a new block and adds it to the blockchain.
func (bc *Blockchain) NewBlock(transaction string, nonce int, previousHash string) {
	block := &Block{
		Transaction:  transaction,
		Nonce:        nonce,
		PreviousHash: previousHash,
		CurrentHash:  "",
	}

	block.CurrentHash = calculateHash(block)
	bc.Blocks = append(bc.Blocks, block)
}

// DisplayBlocks prints all blocks in the blockchain.
func (bc *Blockchain) DisplayBlocks() {
	for i, block := range bc.Blocks {
		fmt.Printf("Block %d:\n", i)
		fmt.Printf("  Transaction: %s\n", block.Transaction)
		fmt.Printf("  Nonce: %d\n", block.Nonce)
		fmt.Printf("  Previous Hash: %s\n", block.PreviousHash)
		fmt.Printf("  Current Hash: %s\n", block.CurrentHash)
	}
}

// ChangeBlock changes the transaction of the given block reference.
func (bc *Blockchain) ChangeBlock(blockIndex int, newTransaction string) error {
	if blockIndex < 0 || blockIndex >= len(bc.Blocks) {
		return fmt.Errorf("Block index out of range")
	}

	block := bc.Blocks[blockIndex]
	block.Transaction = newTransaction
	block.CurrentHash = calculateHash(block)
	return nil
}

// VerifyChain verifies the integrity of the blockchain.
func (bc *Blockchain) VerifyChain() bool {
	for i := 1; i < len(bc.Blocks); i++ {
		currentBlock := bc.Blocks[i]
		previousBlock := bc.Blocks[i-1]

		// Verify that the current block's PreviousHash matches the hash of the previous block.
		if currentBlock.PreviousHash != calculateHash(previousBlock) {
			return false
		}
	}

	return true
}

// calculateHash calculates the hash of a block.
func calculateHash(block *Block) string {
	data := fmt.Sprintf("%s%d%s", block.Transaction, block.Nonce, block.PreviousHash)
	hash := sha256.Sum256([]byte(data))
	return fmt.Sprintf("%x", hash)
}
