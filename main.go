package main

import (
	"fmt"

	blockchain "github.com/pmh-only/blockchain/structures"
)

func main() {
	chain := blockchain.CreateChain()
	block := chain.AddStringBlock("PMH is handsome")

	fmt.Printf("%x\n", block.SerializationWithTail())
	fmt.Printf("%v\n", chain.IsValid())

	chain.Blocks[1].Body = []byte("sans")

	fmt.Printf("%v\n", chain.IsValid())
}
