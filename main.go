package main

import (
	"fmt"

	blockchain "github.com/pmh-only/blockchain/src"
)

func main() {
	// chain := blockchain.CreateChain()

	// chain.AddStringBlock("Wa Sans")
	// chain.AddStringBlock("This is pmh-only/blockchain")

	// lastBlock := chain.GetLatestBlock()

	// fmt.Printf("nonce: %x\n", lastBlock.Head.Nonce)
	// fmt.Printf("serial: %x\n", lastBlock.SerializationWithTail())

	// serial := lastBlock.SerializationWithTail()
	// cloneBlock := blockchain.Deserialization(serial)

	// fmt.Printf("is same? %v\n", string(cloneBlock.SerializationWithTail()) == string(serial))
	// fmt.Printf("is valid? %v %v", lastBlock.IsValid(), cloneBlock.IsValid())

	chain := blockchain.Load("./data")

	fmt.Printf("%v", chain)
	println(chain.IsValid())
}
