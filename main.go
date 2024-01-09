package main

import "fmt"

func main() {
	bc := NewBlockChain()
	bc.NewTransaction("0x00000ad0ad00001", "0xhttps://github.com/Bachamht", 100)
	bc.NewTransaction("0x00000ad0ad00002", "0xhttps://github.com/Bachamht", 100)
	bc.NewTransaction("0x00000ad0ad00003", "0xhttps://github.com/Bachamht", 100)
	bc.NewTransaction("0x00000ad0ad00004", "0xhttps://github.com/Bachamht", 100)
	bc.NewTransaction("0x00000ad0ad00005", "0xhttps://github.com/Bachamht", 100)
	bc.NewTransaction("0x00000ad0ad00006", "0xhttps://github.com/Bachamht", 100)

	//遍历每个区块，并打印出每个区块的数据
	for _, block := range bc.blocks {
		fmt.Println("Index:", block.Index)
		fmt.Println("Timestamp", block.Timestamp)
		fmt.Println("PrevBlock hash:", block.PrevBlockHash)
		fmt.Println("Transactions: ", block.Transactions)
		fmt.Println("Hash: ", block.Hash)
		fmt.Println("Nounce:", block.Nonce)
	}
}
