package main

import (
	"fmt"

	chain "github.com/pmh-only/blockchain/structures"
)

func main() {
	block := chain.CreateBlock(0, []byte{0}, 3, 4, []byte("hi"))
	println(block.Head.CreatedAt)
	fmt.Printf("%08b", block.SerializationWithTail())
}
