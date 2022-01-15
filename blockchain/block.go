package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strings"
)

const ZeroCount = 4

type Block struct {
	PreviousHash [32]byte `json:"previous_hash"` // 直前のブロックから計算されるハッシュ値.
	Transaction  []string `json:"transaction"`
	NONCE        int      `json:"nonce"`
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
	for {
		hash := b.CalcHash()
		if b.canClearRule(hash[:]) {
			break
		}

		b.NONCE++
	}
}

// canClearRule ルールをクリアできるかどうか.
// ハッシュ値の最初の〇〇(ZeroCountで指定する)桁が"0"ならクリア.
func (b *Block) canClearRule(hash []byte) bool {
	return hex.EncodeToString(hash)[0:ZeroCount] == strings.Repeat("0", ZeroCount)
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
	fmt.Printf("1つ前のブロックのハッシュ：         %x\n", b.PreviousHash)
	fmt.Printf("このブロックのトランザクション：    %s\n", b.Transaction)
	fmt.Printf("ナンス：                            %d\n", b.NONCE)
}
