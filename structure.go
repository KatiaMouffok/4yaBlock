package main //Import the main package
// Create the Block data structure
// A block contains this info:

type Block struct {
	Timestamp         int64
	PreviousBlockHash []byte
	MyBlockHash       []byte
	AllData           []byte
}

type Blockchain struct {
	Blocks []*Block
}
