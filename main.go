package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

type Block struct {
	Pos       int
	Hash      string
	PrevHash  string
	Timestamp string
	Data      BookCheckout
}

type Blockchain struct {
	blocks []*Block
}

type Book struct {
	ID          string 
	ISBN        string 
	Title       string 
	Author      string 
	PublishDate string 
}

type BookCheckout struct {
	BookID       string 
	User         string
	CheckoutDate string 
	IsGenesis    bool   
}

func (b *Block) generateHash() {
	bytes, _ := json.Marshal(b.Data)

	data := strconv.Itoa(b.Pos) + b.Timestamp + string(bytes)

	hash := sha256.New()
	hash.Write([]byte(data))
	b.Hash = hex.EncodeToString(hash.Sum(nil))
}

func (b *Block) validateHash(hash string) bool {
	b.generateHash()
	return b.Hash == hash
}

func (bc *Blockchain) AddBlock(data BookCheckout) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	block := CreateBlock(prevBlock, data)
	if validBlock(block, prevBlock) {
		bc.blocks = append(bc.blocks, block)
	}
}

func CreateBlock(prevBlock *Block, data BookCheckout) *Block {
	block := new(Block)
	block.Pos = prevBlock.Pos + 1
	block.Data = data
	block.Timestamp = time.Now().String()
	block.PrevHash = prevBlock.Hash
	block.generateHash()

	return block
}

func validBlock(block *Block, prevBlock *Block) bool {
	if block.PrevHash != prevBlock.Hash {
		return false
	}
	if !block.validateHash(block.Hash) {
		return false
	}
	if prevBlock.Pos+1 != block.Pos {
		return false
	}
	return true
}

func GenesisBlock() *Block {
	return CreateBlock(new(Block), BookCheckout{IsGenesis: true})
}

func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{GenesisBlock()}}
}

func main() {
	blockchain := NewBlockchain()
	for _, block := range blockchain.blocks {
		fmt.Printf("Prev.hash: %x\n", block.PrevHash)
		bytes, _ := json.MarshalIndent(block.Data, "", " ")
		fmt.Printf("Data:%v\n", string(bytes))
		fmt.Printf("Hash:%x\n", block.Hash)
		fmt.Println()
	}
}
