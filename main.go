package main

import "time"

type Block struct {
	PrevHash  uint8
	Hash      uint8
	Timestamp time.Time
	Data      string
}

type Blockchain struct {
	bch []*Block
}