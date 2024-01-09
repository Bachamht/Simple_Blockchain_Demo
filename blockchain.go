package main

const Tranamount = 2

// 区块链数组
type BlockChain struct {
	blocks              []*Block
	CurrentTransactions []Transaction
}

// 创建区块链（带创世块）
func NewBlockChain() *BlockChain {
	return &BlockChain{[]*Block{NewGesisBlock()}, []Transaction{}}
}

// 向链中加入新区块
func (bc *BlockChain) AddBlock() {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(bc, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

// 增加交易信息，当交易信息超过100条时候打包成块
func (bc *BlockChain) NewTransaction(sender, recipient string, amount float64) {
	transaction := Transaction{
		Sender:    sender,
		Recipient: recipient,
		Amount:    amount,
	}
	bc.CurrentTransactions = append(bc.CurrentTransactions, transaction)
	if len(bc.CurrentTransactions) >= Tranamount {
		bc.AddBlock()
		bc.CurrentTransactions = make([]Transaction, 0)
	}
}

// 返回最后一个区块
func (bc *BlockChain) LastBlock() *Block {
	return bc.blocks[len(bc.blocks)-1]
}
