package blockchain

type Block struct {
	Head BlockHead
	Body BlockBody
	Tail BlockTail
}

type BlockHead struct {
	Index      uint16
	CreatedAt  uint32
	PrevHash   []byte
	Nonce      uint32
	Difficulty uint8
}

type BlockBody struct {
	Flag    BodyFlags
	Message []byte
}

type BlockTail struct {
	CurrHash []byte
}

type BodyFlags byte

type Chain struct {
	Blocks []Block
}

type BodyStructure struct {
	Flag    string `json:"flag"`
	Message string `json:"message"`
}

type Peer struct {
	chain *Chain
	data  PeerData
}

type PeerData struct {
	Address string `json:"address"`
	Port    int    `json:"port"`
	Version string `json:"version"`
}
