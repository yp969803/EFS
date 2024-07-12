package p2p

import (
	
	"fmt"
	"net"
	"sync"
)

// TCPeer represents the remote node over a  TCP established connection
type TCPeer struct {
	conn net.Conn

	// if we dial a connection => outbound == true
	// if we accept a connection => outbound == false
	outbound bool
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPeer {
       return &TCPeer{
		conn: conn,
		outbound: outbound,
	   }
}

type TCPTransport struct {
	listenAddress string 
	listener       net.Listener
	handshakeFunc  HandshakeFunc
	decoder        Decoder
	mu            sync.RWMutex
	peers        map[net.Addr]Peer

}

func NewTCPTrancport(listenAddr string) *TCPTransport{
	return &TCPTransport{
		handshakeFunc: NOPHandshakeFunc,
		listenAddress: listenAddr,
	}
}

func (t *TCPTransport) ListenAndAcccept() error {
	var err error 
	t.listener, err = net.Listen("tcp", t.listenAddress)
	if err!=nil {
        return err 
	}
	go t.startAcceptLoop()
	return nil
    
}

func (t *TCPTransport) startAcceptLoop() {
	for {
		conn, err  := t.listener.Accept()
		if err!=nil {
			fmt.Printf("TCP accept error: %s \n", err)
		}
		go t.handleConn(conn)

	}
}

type Temp struct {}

func (t *TCPTransport) handleConn(conn net.Conn){
	peer := NewTCPPeer(conn, true)

	if err := t.handshakeFunc(peer); err!= nil {

	}

	msg := &Temp{}

	// Read loop
	for {
        if err := t.decoder.Decode(conn, msg); err!= nil {
           fmt.Println("TCP error: %s \n", err)
		   continue;
		}
	}
	fmt.Printf("new incoming connection %+v \n", peer)
}


