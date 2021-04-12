package blockchain

import (
	"reflect"
)

const START_DIFFICULTY = 4

type Chain struct {
	Blocks []Block
}

type BodyStructure struct {
	Flag    string `json:"flag"`
	Message string `json:"message"`
}

func CreateChain() Chain {
	chain := Chain{
		Blocks: []Block{},
	}

	chain.CreateGenesisBlock()
	return chain
}

func (chain *Chain) CreateGenesisBlock() {
	genesis := CreateBlock(
		0, []byte{0}, 0,
		START_DIFFICULTY,
		GENESIS, []byte{},
	)
	genesis.MineBlock()
	chain.Blocks = append(chain.Blocks, genesis)
}

func (chain Chain) GetLatestBlock() Block {
	return chain.Blocks[len(chain.Blocks)-1]
}

func (chain *Chain) AddBlock(newBlock Block) {
	newBlock.Head.PrevHash = chain.GetLatestBlock().Tail.CurrHash
	newBlock.MineBlock()

	chain.Blocks = append(chain.Blocks, newBlock)
}

func (chain *Chain) AddStringBlock(data string) {
	latestBlock := chain.GetLatestBlock()
	newBlock := CreateBlock(
		latestBlock.Head.Index+1,
		latestBlock.Tail.CurrHash, 0,
		latestBlock.Head.Difficulty,
		DATA, []byte(data),
	)

	chain.AddBlock(newBlock)
}

func (chain Chain) IsValid() bool {
	for index := range chain.Blocks {
		if index < 1 {
			continue
		}

		currBlock := chain.Blocks[index]
		prevBlock := chain.Blocks[index-1]

		if !reflect.DeepEqual(currBlock.Tail.CurrHash, currBlock.CaculateHash()) {
			return false
		}

		if !reflect.DeepEqual(currBlock.Head.PrevHash, prevBlock.Tail.CurrHash) {
			return false
		}
	}

	return true
}
