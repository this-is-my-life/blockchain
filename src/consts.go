package blockchain

const (
	GENESIS BodyFlags = 0
	PEER    BodyFlags = 1
	DATA    BodyFlags = 2
)

const SIZE_OF_HEAD = 2 + 4 + 64 + 4 + 1

const START_DIFFICULTY = 8

const DIFFICULTY_INCREASE_STEP = 1024

const PEER_NETWORK_VERSION = "v0.1"
