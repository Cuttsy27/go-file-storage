package main

import (
	"log"

	"github.com/cuttsy27/go-file-storage/p2p"
)

func main() {
	tcpOpts := p2p.TCPTransportOpts{
		ListenAddress: ":4000",
		HandshakeFunc: p2p.NOPHandshakeFunc,
		Decoder:       p2p.GOBDecoder{},
	}
	tr := p2p.NewTCPTransport(tcpOpts)
	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}

	select {}
}
