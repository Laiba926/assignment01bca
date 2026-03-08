package main

import (
	"Assignment1/assignment01bca"
)

func main() {

	// Genesis Block (Roll Number Included)
	genesis := assignment01bca.NewBlock("Genesis Block - 22I4947", 28, "")

	// Other Blocks (include last 3 digits 947)
	b1 := assignment01bca.NewBlock("Alice to Bob 947", 5, genesis.Hash)
	assignment01bca.NewBlock("Bob to Charlie", 7, b1.Hash)

	println("Blockchain:")
	assignment01bca.ListBlocks()

	println("Verifying Blockchain:")
	assignment01bca.VerifyChain()

	println("Tampering Block 1...")

	assignment01bca.ChangeBlock(1, "Hacker Transaction")

	println("Verifying Blockchain Again:")
	assignment01bca.VerifyChain()
}
