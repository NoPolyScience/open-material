package network

import (
	"net"
	//"github.com/sirupsen/logrus"
)

type MessageType uint8

// Possible message types:
// Proposal creation for new scientific data addition to the db
// Proposal to add new institutions to the network.
type Message struct {
	From        string
	MessageType MessageType
	Payload     []byte
}

const (
	NewMaterial MessageType = iota
	NewPeer
)

type Peer struct {
	conn net.Conn
}

type ServerConfig struct {
	Version    string
	ListenAddr string
}

type Server struct {
	ServerConfig

	ln      net.Listener
	peers   map[net.Addr]*Peer
	addPeer chan *Peer
	delPeer chan *Peer
	msgch   chan *Message
}

func NewServer(cfg ServerConfig) *Server {
	return &Server{
		ServerConfig: cfg,
		peers:        make(map[net.Addr]*Peer),
		addPeer:      make(chan *Peer),
		delPeer:      make(chan *Peer),
		msgch:        make(chan *Message),
	}
}

func (s *Server) Start() {

}
