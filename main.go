package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

type Block struct {
	PrevHash  []byte
	Hash      []byte
	Data      []byte
}

type Blockchain struct {
	blocks []*Block
} 

func(b *Block) DeriveHash() {
	sum := bytes.Join([][]byte{[]byte(b.Data), []byte(b.PrevHash)}, []byte{})
	hash :=sha256.Sum256(sum)
	b.Hash = hash[:]
}

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{
		PrevHash: prevHash,
		Data: []byte(data),
	}

	block.DeriveHash()
	return block
}

func(chain *Blockchain) AddBlock (data string) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	new := CreateBlock(data,prevBlock.Hash)
	chain.blocks = append(chain.blocks, new)
}

func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

func InitBlockChain() *Blockchain {
	return &Blockchain{[] *Block{Genesis()}}
}

func main() {
	chain := InitBlockChain()
	chain.AddBlock("First block")
	chain.AddBlock("Second block")
	chain.AddBlock("Third block")

	for _, block := range chain.blocks {
		fmt.Printf("Previous hash: %x\n",block.PrevHash)
		fmt.Printf("Hash: %x\n",block.Hash)
		fmt.Printf("Dats: %s\n",block.Data)
	}
}