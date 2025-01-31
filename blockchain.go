package main

func (blockchain *Blockchain) AddBlock(data string, difficulty int64) {
	PreviousBlock := blockchain.Blocks[len(blockchain.Blocks)-1]      // the previous block is needed, so let's get it
	newBlock := NewBlock(data, PreviousBlock.MyBlockHash, difficulty) // create a new block containing the data and the hash of the previous block
	blockchain.Blocks = append(blockchain.Blocks, newBlock)           // add that block to the chain to create a chain of blocks
}

func NewBlockchain() *Blockchain { // the function is created
	return &Blockchain{[]*Block{NewGenesisBlock(0)}} // the genesis block is added first to the chain
}
