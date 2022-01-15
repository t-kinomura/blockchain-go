package main

import "github.com/t-kinomura/blockchain-go/blockchain"

func main() {
	bc := blockchain.NewBlockchain()

	bc.AddTransaction("transaction1")
	bc.AddTransaction("transaction2")
	bc.CreateBlock()

	bc.AddTransaction("transaction3")
	bc.AddTransaction("transaction4")
	bc.CreateBlock()

	bc.Print()
}


