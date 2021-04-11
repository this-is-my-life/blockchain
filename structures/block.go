package structures

import (
	"encoding/binary"
	"time"

	"github.com/pmh-only/blockchain/utils"
	"golang.org/x/crypto/sha3"
)

type Block struct {
	Head BlockHead
	Body BlockBody
	Tail BlockTail
}

const SIZE_OF_HEAD = 2 + 4 + 64 + 4 + 1

type BlockHead struct {
	Index      uint16
	CreatedAt  uint32
	PrevHash   []byte
	Nonce      uint32
	Difficulty uint8
}

type BlockBody []byte

type BlockTail struct {
	CurrHash []byte
}

func CreateBlock(index uint16, prevHash []byte, nonce uint32, difficulty uint8, body []byte) Block {
	block := Block{}

	block.Head.Index = index
	block.Head.CreatedAt = uint32(time.Now().Unix())
	block.Head.PrevHash = prevHash
	block.Head.Nonce = nonce
	block.Head.Difficulty = difficulty

	block.Body = body

	block.Tail.CurrHash = block.CaculateHash()

	return block
}

func (block Block) CaculateHash() []byte {
	serial := block.SerializationWithoutTail()
	hash := make([]byte, 64)
	sha3.ShakeSum256(hash, serial)
	return hash
}

func (block Block) SerializationWithoutTail() []byte {
	head := make([]byte, 0, SIZE_OF_HEAD)
	body := []byte(string(block.Body))

	index := make([]byte, 2)
	createdAt := make([]byte, 4)
	nonce := make([]byte, 4)

	prevHash := utils.PadOrTrim(block.Head.PrevHash, 64)
	difficulty := block.Head.Difficulty

	binary.BigEndian.PutUint16(index, block.Head.Index)
	binary.BigEndian.PutUint32(createdAt, block.Head.CreatedAt)
	binary.BigEndian.PutUint32(nonce, block.Head.Nonce)

	head = append(head, index...)
	head = append(head, createdAt...)
	head = append(head, prevHash...)
	head = append(head, nonce...)
	head = append(head, difficulty)

	return append(head, body...)
}

func (block Block) SerializationWithTail() []byte {
	return append(block.SerializationWithoutTail(), block.CaculateHash()...)
}
