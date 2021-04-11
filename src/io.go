package blockchain

import (
	"encoding/binary"

	"github.com/dgraph-io/badger/v3"
)

func Load(path string) Chain {
	chain := Chain{}
	chain.Blocks = []Block{}

	db, err := badger.Open(badger.DefaultOptions(path))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_ = db.View(func(txn *badger.Txn) error {
		for indexCnt := 0; ; indexCnt++ {
			key := make([]byte, 2)
			binary.BigEndian.PutUint16(key, uint16(indexCnt))

			item, err := txn.Get(key)
			if err != nil {
				break
			}

			serial := []byte{}
			_ = item.Value(func(val []byte) error {
				serial = append(serial, val...)
				return nil
			})

			block := Deserialization(serial)
			chain.Blocks = append(chain.Blocks, block)
		}
		return nil
	})

	if len(chain.Blocks) < 1 {
		chain.CreateGenesisBlock()
	}

	return chain
}
