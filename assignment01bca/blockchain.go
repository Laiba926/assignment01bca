package assignment01bca

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"time"
)

type Block struct {
	Index        int
	Timestamp    string
	Transaction  string
	Nonce        int
	PreviousHash string
	Hash         string
}

var Blockchain []Block

func CalculateHash(data string) string {
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}

func NewBlock(transaction string, nonce int, previousHash string) *Block {

	newBlock := Block{}

	newBlock.Index = len(Blockchain)
	newBlock.Timestamp = time.Now().String()
	newBlock.Transaction = transaction
	newBlock.Nonce = nonce
	newBlock.PreviousHash = previousHash

	data := transaction +
		strconv.Itoa(nonce) +
		previousHash +
		strconv.Itoa(newBlock.Index) +
		newBlock.Timestamp

	newBlock.Hash = CalculateHash(data)

	Blockchain = append(Blockchain, newBlock)

	return &newBlock
}

func ListBlocks() {

	for _, block := range Blockchain {

		println("Index:", block.Index)
		println("Timestamp:", block.Timestamp)
		println("Transaction:", block.Transaction)
		println("Nonce:", block.Nonce)
		println("PreviousHash:", block.PreviousHash)
		println("Hash:", block.Hash)
		println("----------------------------")
	}
}
func ChangeBlock(index int, newTransaction string) {

	if index < len(Blockchain) {
		Blockchain[index].Transaction = newTransaction
	}

}
func VerifyChain() {

	for i := 1; i < len(Blockchain); i++ {

		current := Blockchain[i]
		previous := Blockchain[i-1]

		data := current.Transaction +
			strconv.Itoa(current.Nonce) +
			current.PreviousHash +
			strconv.Itoa(current.Index) +
			current.Timestamp

		calculatedHash := CalculateHash(data)

		if calculatedHash != current.Hash {
			println("Blockchain Tampered at Block", i)
			return
		}

		if current.PreviousHash != previous.Hash {
			println("Blockchain Broken at Block", i)
			return
		}
	}

	println("Blockchain Verified Successfully")
}
