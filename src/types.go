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

const (
	GENESIS BodyFlags = 0
	PEER    BodyFlags = 1
	DATA    BodyFlags = 2
)
