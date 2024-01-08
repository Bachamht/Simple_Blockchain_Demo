package main

import "fmt"

func main() {
	bc := NewBlockChain()
	bc.NewTransaction("0x00000ad0ad00006", "0x000000000000000", 100)
	bc.NewTransaction("0x00000ad0ad00006", "0x000000000000000", 100)
	bc.NewTransaction("0x00000ad0ad00006", "0x000000000000000", 100)
	bc.NewTransaction("0x00000ad0ad00006", "0x000000000000000", 100)
	bc.NewTransaction("0x00000ad0ad00006", "0x000000000000000", 100)
	bc.NewTransaction("0x00000ad0ad00006", "0x000000000000000", 100)

	for _, block := range bc.blocks {
		fmt.Println("Prev hash:", block.PrevBlockHash)
		fmt.Println("Data: ", block.Transactions)
		fmt.Println("Hash: ", block.Hash)
	}

}
