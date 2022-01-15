package main

import (
	"fmt"
	"github.com/t-kinomura/blockchain-go/blockchain"
	"time"
)

func main() {
	bc := blockchain.NewBlockchain()

	bc.AddTransaction("transaction1")
	bc.AddTransaction("transaction2")
	bc.CreateBlock()

	now := time.Now()
	bc.AddTransaction("transaction3")
	bc.AddTransaction("transaction4")
	bc.CreateBlock()
	ms := time.Since(now).Milliseconds()

	bc.Print()

	fmt.Printf("1ブロック生成にかかる時間: %vms\n", ms)
}


