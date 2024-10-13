package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTCPTransport(t *testing.T) {
	tcpOpts := TCPTransportOpts{
		ListenAddress: ":4000",
		HandshakeFunc: NOPHandshakeFunc,
	}
	tr := NewTCPTransport(tcpOpts)
	assert.Equal(t, tcpOpts.ListenAddress, tr.ListenAddress)

	assert.Nil(t, tr.ListenAndAccept())

	select {}
}
