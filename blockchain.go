package main

import (
	"fmt"
	"strings"
)

type Blockchain struct {
	transactionPool []string // トランザクションを一時的にプールするフィールド
	block           []*Block
}

func NewBlockchain() *Blockchain {
	bc := &Blockchain{}
	bc.CreateBlock() // Genesis Blockの作成

	return bc
}

func (bc *Blockchain) CreateBlock() {
	if len(bc.block) == 0 {
		b := NewBlock([32]byte{}, nil)
		bc.block = append(bc.block, b)
	} else {
		previousHash := bc.block[len(bc.block)-1].CalcHash()

		b := NewBlock(previousHash, bc.transactionPool)
		bc.block = append(bc.block, b)
	}

	bc.transactionPool = nil
}

func (bc *Blockchain) AddTransaction(transaction string) {
	bc.transactionPool = append(bc.transactionPool, transaction)
}

func (bc *Blockchain) Print() {
	for i, b := range bc.block {
		fmt.Printf("%s Block %d %s\n", strings.Repeat("=", 15), i, strings.Repeat("=", 15))
		b.Print()
	}
}
