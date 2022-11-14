package assignment02

import (
	"crypto/sha256"
	"fmt"
	"math/rand"
	"strconv"
)

type Transaction struct {
	TransactionID string
	Sender        string
	Receiver      string
	Amount        int
}

type Block struct {
	Nonce       int
	BlockData   []Transaction
	PrevPointer *Block
	PrevHash    string
	CurrentHash string
}

type Blockchain struct {
	ChainHead *Block
}

func GenerateNonce(blockData []Transaction) int {
	min := 1
	max := 69
	a := rand.Intn(max-min) + min
	return a
}

func CalculateHash(blockData []Transaction, nonce int) string {
	dataString := ""
	for i := 0; i < len(blockData); i++ {
		dataString += (blockData[i].TransactionID + blockData[i].Sender +
			blockData[i].Receiver + strconv.Itoa(blockData[i].Amount)) + strconv.Itoa(nonce)
	}
	return fmt.Sprintf("%x", sha256.Sum256([]byte(dataString)))
}

func NewBlock(blockData []Transaction, chainHead *Block) *Block {
	Nonce:= GenerateNonce(blockData)
	if chainHead == nil {
		b := new(Block)
		b.PrevPointer = nil
		b.BlockData = blockData
		b.Nonce = Nonce
		b.CurrentHash = CalculateHash(blockData, Nonce) 
		return b

	} else {
		
		bk := new(Block)
		bk.BlockData = blockData
		bk.PrevPointer = chainHead
		bk.Nonce = Nonce
		bk.CurrentHash = CalculateHash(blockData, Nonce) 
		bk.PrevHash = bk.PrevPointer.CurrentHash
		chainHead = bk
		return bk

	}

}

func ListBlocks(chainHead *Block) {
	a := chainHead
	for a != nil {
		fmt.Println(*a)
		DisplayTransactions(a.BlockData)
		a = a.PrevPointer
	}

}

func DisplayTransactions(blockData []Transaction) {
	for a := 0; a < len(blockData); a++ {
		fmt.Println("TransactionID = ", blockData[a].TransactionID)
		fmt.Println("Sender = ", blockData[a].Sender)
		fmt.Println("Receiver = ", blockData[a].Receiver)
		fmt.Println("Amount = ", blockData[a].Amount)
	}
}

func NewTransaction(sender string, receiver string, amount int) Transaction {
	var trans string
	t := new(Transaction)
	t.TransactionID = trans
	t.Sender = sender
	t.Receiver = receiver
	t.Amount = amount
	return *t
}
