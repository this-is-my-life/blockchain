package blockchain

import (
	"encoding/binary"
	"fmt"
	"net"

	"github.com/pmh-only/blockchain/utils"
)

func (chain *Chain) CreatePeer(port int) Peer {
	peer := Peer{
		chain: chain,
		data: PeerData{
			Address: utils.GetPublicIp(),
			Port:    port,
			Version: PEER_NETWORK_VERSION,
		},
	}

	return peer
}

func (peer *Peer) Open() {
	fmt.Printf("Peer Server is now on UDP://0.0.0.0:%d\nq", peer.data.Port)

	packet := make([]byte, 1024)
	addr := net.UDPAddr{
		Port: peer.data.Port,
		IP:   net.ParseIP("0.0.0.0"),
	}

	server, err := net.ListenUDP("udp", &addr)
	if err != nil {
		panic(err)
	}

	for {
		_, remote, err := server.ReadFromUDP(packet)

		if err != nil {
			continue
		}

		index := int(binary.BigEndian.Uint16(packet))
		if len(peer.chain.Blocks) <= index {
			continue
		}

		fmt.Printf("[net::%v] %d\n", remote, index)

		blocklock := peer.chain.Blocks[index]
		serial := blocklock.SerializationWithTail()

		server.WriteToUDP(serial, remote)
	}
}
