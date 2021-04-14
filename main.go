package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"strconv"
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

	if subcmd == "serve" {
		if len(args) < 2 {
			println("Uncompleted Subcommand")
			return
		}

		port, _ := strconv.Atoi(args[1])
		peer := chain.CreatePeer(port)

		chain.AddPeerBlock(peer)
		chain.Save(DATA_DIRECTORY)

		peer.Open()
	}

	if subcmd == "client-test" {
		if len(args) < 3 {
			println("Uncompleted Subcommand")
			return
		}

		packet := make([]byte, 2048)

		indexraw, _ := strconv.Atoi(args[2])
		index := make([]byte, 2)

		binary.BigEndian.PutUint16(index, uint16(indexraw))

		conn, _ := net.Dial("udp", args[1])
		conn.Write(index)

		_, err := bufio.NewReader(conn).Read(packet)
		if err == nil {
			fmt.Printf("%s\n", packet)
		} else {
			fmt.Printf("Some error %v\n", err)
		}
		conn.Close()
	}
}
