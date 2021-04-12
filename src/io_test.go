package blockchain_test

import (
	"reflect"
	"testing"

	blockchain "github.com/pmh-only/blockchain/src"
)

func TestIo(t *testing.T) {
	chain := blockchain.Load("./data_test")

	chain.AddStringBlock("hi")

	lastBlock := chain.GetLatestBlock()
	chain.Save("./data_test")

	chain2 := blockchain.Load("./data_test")
	lastBlock2 := chain2.GetLatestBlock()

	got := lastBlock2.SerializationWithTail()
	want := lastBlock.SerializationWithTail()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
