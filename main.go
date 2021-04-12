package main

import (
	"fmt"
	"os"
	"strings"

	blockchain "github.com/pmh-only/blockchain/src"
)

func main() {
	DATA_DIRECTORY := os.Getenv("DATA_DIRECTORY")
	if len(DATA_DIRECTORY) < 1 {
		DATA_DIRECTORY = "./data"
	}

	args := os.Args[1:]
	if len(args) < 1 {
		println("Invalid Subcommnd\n")
		return
	}

	subcmd := args[0]

	chain := blockchain.Load(DATA_DIRECTORY)

	fmt.Printf("Chain Valid: %v\n", chain.IsValid())
	fmt.Printf("Chain Length: %d blocks\n", len(chain.Blocks))

	fmt.Println()

	if subcmd == "list" {
		for _, block := range chain.Blocks {
			fmt.Printf("%d. (%d) %s\n", block.Head.Index, block.Body.Flag, block.Body.Message)
		}

		return
	}

	if subcmd == "list-detail" {
		for _, block := range chain.Blocks {
			fmt.Printf("\nIndex: %d\nCreatedAt: %d\nPrevHash: %x\nNonce: %d\nDifficulty: %d\nBody: (%d) %s\nCurrHash: %x\n",
				block.Head.Index,
				block.Head.CreatedAt,
				block.Head.PrevHash,
				block.Head.Nonce,
				block.Head.Difficulty,
				block.Body.Flag,
				block.Body.Message,
				block.Tail.CurrHash,
			)
		}

		return
	}

	if subcmd == "add" {
		if len(args) < 2 {
			println("Uncompleted Subcommand")
			return
		}

		data := strings.Join(args[1:], " ")
		chain.AddStringBlock(data)
		chain.Save(DATA_DIRECTORY)

		lastCreated := chain.GetLatestBlock()

		fmt.Printf("Block Created:\n%x", lastCreated.SerializationWithTail())
		return
	}

	if subcmd == "drop" {
		chain.Drop(DATA_DIRECTORY)

		println("Dropped")
	}
}
