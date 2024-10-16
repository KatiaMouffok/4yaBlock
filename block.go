package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

func CreateBlock(Hearder, Body string) {
	fmt.Println(Hearder, "\n", Body)
}

func (block *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(block.Timestamp, 10))
	headers := bytes.Join([][]byte{timestamp, block.PreviousBlockHash, block.AllData}, []byte{})
	hash := sha256.Sum256(headers)
	block.MyBlockHash = hash[:]
}

func NewBlock(data string, prevBlockHash []byte, difficulty int64) *Block {
	block := &Block{time.Now().Unix(), prevBlockHash, []byte{}, []byte(data), 0, difficulty}
	block.SetHash()
	return block
}
func NewGenesisBlock(difficulty int64) *Block {
	return NewBlock("Genesis Block", []byte{}, difficulty)
}
