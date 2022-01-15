package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
)

type Block struct {
	PreviousHash [32]byte `json:"previous_hash"` // 直前のブロックから計算されるハッシュ値.
	Transaction  []string `json:"transaction"`
	NONCE        int64    `json:"nonce"`
}

func NewBlock(previousHash [32]byte, transaction []string) *Block {
	b := &Block{}

	b.PreviousHash = previousHash
	b.Transaction = transaction
	b.mining()

	return b
}

// mining NANCEの値を求める.
func (b *Block) mining() {
	b.NONCE = 0
	//var nance int64 = 0
	for {
		hash := b.CalcHash()
		if hex.EncodeToString(hash[:])[0:4] == "0000" {
			break
		}

		b.NONCE++
	}
}

// CalcHash Hash ブロックのハッシュを計算する.
func (b *Block) CalcHash() [32]byte {
	m, err := json.Marshal(b)
	if err != nil {
		panic(err)
	}

	return sha256.Sum256(m)
}

func (b *Block) Print() {
	fmt.Printf("PreviousHash    %x\n", b.PreviousHash)
	fmt.Printf("Transaction     %s\n", b.Transaction)
	fmt.Printf("NANCE           %d\n", b.NONCE)
}
