package blockchain_test

import (
	"strings"
	"testing"

	blockchain "github.com/pmh-only/blockchain/src"
)

func TestCreateChain(t *testing.T) {
	chain := blockchain.CreateChain()
	lastBlock := chain.GetLatestBlock()

	got := lastBlock.IsValid()
	want := true

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

	got2 := lastBlock.Body.Flag
	want2 := blockchain.GENESIS

	if got2 != want2 {
		t.Errorf("got %v, wanted %v", got2, want2)
	}
}

func TestAddBlock(t *testing.T) {
	chain := blockchain.CreateChain()
	genesis := chain.GetLatestBlock()

	newBlock := blockchain.CreateBlock(
		1, genesis.Tail.CurrHash,
		0, 4, blockchain.DATA, []byte("alice"),
	)

	chain.AddBlock(newBlock)

	got := chain.IsValid()
	want := true

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestAddStringBlock(t *testing.T) {
	chain := blockchain.CreateChain()
	chain.AddStringBlock("alice")

	got := chain.IsValid()
	want := true

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

	lastBlock := chain.GetLatestBlock()

	got2 := string(lastBlock.Body.Message)
	want2 := "alice"

	if strings.Compare(got2, want2) != 0 {
		t.Errorf("got %v, wanted %v", got2, want2)
	}
}
