package main

import (
	"time"
)

// 交易数据结构
type Transaction struct {
	Sender    string  `json:"sender"`
	Recipient string  `json:"recipient"`
	Amount    float64 `json:"amount"`
}

// 区块结构
type Block struct {
	Index         int
	Timestamp     int64
	Transactions  []Transaction
	PrevBlockHash []byte
	Hash          []byte
	//Data          []byte
	Nonce int
}

// 创建新区块
func NewBlock(bc *BlockChain, preBlockHash []byte) *Block {
	block := &Block{
		Index:         len(bc.blocks) + 1,
		Timestamp:     time.Now().Unix(),
		PrevBlockHash: preBlockHash,
		Hash:          []byte{},
		Transactions:  bc.CurrentTransactions,
		Nonce:         0}
	pow := NewProofOfWork(block)
	nounce, hash := pow.Run()
	block.Hash = hash[:]
	block.Nonce = nounce
	return block
}

// 创世块创建
func NewGesisBlock() *Block {
	block := &Block{
		Index:         0,
		Timestamp:     time.Now().Unix(),
		PrevBlockHash: nil,
		Hash:          []byte{},
		Transactions:  []Transaction{{Sender: "0x0000000000000001", Recipient: "0x0000000000000000", Amount: 100}},
		Nonce:         0,
	}
	return block
}
