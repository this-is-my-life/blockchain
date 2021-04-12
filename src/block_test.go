package blockchain_test

import (
	"encoding/hex"
	"reflect"
	"testing"

	blockchain "github.com/pmh-only/blockchain/src"
)

func TestCreateBlock(t *testing.T) {
	block := blockchain.CreateBlock(0, []byte{}, blockchain.DATA, []byte{})
	block.MineBlock()

	got := block.IsValid()
	want := true

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestCaculateHash(t *testing.T) {
	prevHash, _ := hex.DecodeString("00008665778b92fa4d130089d5e02dddb75f4c4d853bf7a96f43f40a275c42687dbf010ea69446a17c3bd1962ef914655248afe9b4f77a51fe097e07c3e35444")

	block := blockchain.Block{
		Head: blockchain.BlockHead{
			Index:      1,
			CreatedAt:  1618186647,
			PrevHash:   prevHash,
			Nonce:      17709,
			Difficulty: 4,
		},
		Body: blockchain.BlockBody{
			Flag:    blockchain.DATA,
			Message: []byte("hi"),
		},
	}

	got := block.CaculateHash()
	want, _ := hex.DecodeString("0000e14e777cd3890638fa6260e9c7453395965eca076fa7679e25ba0c65206efe5f94d7bfd078786d4d4d1a58e873c4614bd743a2cfb33c8e14d86b66f8a4c4")

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestSerialization(t *testing.T) {
	prevHash, _ := hex.DecodeString("00008665778b92fa4d130089d5e02dddb75f4c4d853bf7a96f43f40a275c42687dbf010ea69446a17c3bd1962ef914655248afe9b4f77a51fe097e07c3e35444")
	currHash, _ := hex.DecodeString("0000e14e777cd3890638fa6260e9c7453395965eca076fa7679e25ba0c65206efe5f94d7bfd078786d4d4d1a58e873c4614bd743a2cfb33c8e14d86b66f8a4c4")

	block := blockchain.Block{
		Head: blockchain.BlockHead{
			Index:      1,
			CreatedAt:  1618186647,
			PrevHash:   prevHash,
			Nonce:      17709,
			Difficulty: 4,
		},
		Body: blockchain.BlockBody{
			Flag:    blockchain.DATA,
			Message: []byte("hi"),
		},
		Tail: blockchain.BlockTail{
			CurrHash: currHash,
		},
	}

	got := block.SerializationWithoutTail()
	want, _ := hex.DecodeString("00016073919700008665778b92fa4d130089d5e02dddb75f4c4d853bf7a96f43f40a275c42687dbf010ea69446a17c3bd1962ef914655248afe9b4f77a51fe097e07c3e354440000452d04026869")

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}

	got = block.SerializationWithTail()
	want, _ = hex.DecodeString("00016073919700008665778b92fa4d130089d5e02dddb75f4c4d853bf7a96f43f40a275c42687dbf010ea69446a17c3bd1962ef914655248afe9b4f77a51fe097e07c3e354440000452d040268690000e14e777cd3890638fa6260e9c7453395965eca076fa7679e25ba0c65206efe5f94d7bfd078786d4d4d1a58e873c4614bd743a2cfb33c8e14d86b66f8a4c4")

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestIsValid(t *testing.T) {
	block := blockchain.CreateBlock(0, []byte{}, blockchain.DATA, []byte("alice"))

	block.MineBlock()
	got := block.IsValid()
	want := true

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

	block.Body.Message = []byte("eve")

	got = block.IsValid()
	want = false

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestMining(t *testing.T) {
	block := blockchain.CreateBlock(1024, []byte{}, blockchain.DATA, []byte("alice"))

	got := block.IsMined()
	want := false

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

	block.MineBlock()

	got = block.IsMined()
	want = true

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}

func TestDeserialization(t *testing.T) {
	serial, _ := hex.DecodeString("00016073919700008665778b92fa4d130089d5e02dddb75f4c4d853bf7a96f43f40a275c42687dbf010ea69446a17c3bd1962ef914655248afe9b4f77a51fe097e07c3e354440000452d040268690000e14e777cd3890638fa6260e9c7453395965eca076fa7679e25ba0c65206efe5f94d7bfd078786d4d4d1a58e873c4614bd743a2cfb33c8e14d86b66f8a4c4")
	block := blockchain.Deserialization(serial)

	got := block.Body.Message
	want := []byte("hi")

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
}
