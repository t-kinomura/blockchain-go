package blockchain

import (
	"fmt"
)

type Blockchain struct {
	TransactionPool []string // トランザクションを一時的にプールするフィールド
	Block           []*Block
}

func NewBlockchain() *Blockchain {
	bc := &Blockchain{}
	bc.CreateBlock() // 最初のBlockの作成

	return bc
}

func (bc *Blockchain) CreateBlock() {
	if len(bc.Block) == 0 {
		b := NewBlock([32]byte{}, nil)
		bc.Block = append(bc.Block, b)
	} else {
		previousHash := bc.Block[len(bc.Block)-1].CalcHash()

		b := NewBlock(previousHash, bc.TransactionPool)
		bc.Block = append(bc.Block, b)
	}

	bc.TransactionPool = nil
}

func (bc *Blockchain) AddTransaction(transaction string) {
	bc.TransactionPool = append(bc.TransactionPool, transaction)
}

func (bc *Blockchain) Print() {
	for i, b := range bc.Block {
		fmt.Printf("%d個目のブロック \n", i)
		b.Print()
		fmt.Println()
	}
}
