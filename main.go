package main

import (
	"fmt"

	blockchain "github.com/pmh-only/blockchain/structures"
)

func main() {
	chain := blockchain.CreateChain()
	chain.AddStringBlock("PMH is handsome")

	fmt.Printf("%v\n", chain.GetLatestBlock())
	fmt.Printf("%x\n", chain.GetLatestBlock().Head.PrevHash)
	fmt.Printf("%08b\n", chain.GetLatestBlock().SerializationWithTail())
}
