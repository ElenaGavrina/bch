package main

import "time"

type Block struct {
	PrevHash  string
	Hash      string
	Timestamp time.Time
	Data      string
}

type Blockchain struct {
	bch []*Block
} 

func NewBlock() *Block {
	return &Block{
		
	}
}