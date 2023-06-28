package server

import (
	"fmt"
	"log"
	"net"
	"sync"
)

type Peer struct {
	conn net.Conn
}

func (p *Peer) Send(b []byte) error {
	_, err := p.conn.Write(b)
	return err
}

type ServerConfig struct {
	ListenAddr string
	Version    string
}

type Message struct {
	payload []byte
	From    net.Addr
}

type Server struct {
	ServerConfig

	rw       sync.RWMutex
	handler  Handler
	peers    map[net.Addr]*Peer
	listener net.Listener
	addPeer  chan *Peer
	delPeer  chan *Peer
	msgCh    chan *Message
}

func NewServer(sCfg ServerConfig) *Server {
	return &Server{
		handler:      NewHandler(),
		ServerConfig: sCfg,
		peers:        make(map[net.Addr]*Peer),
		addPeer:      make(chan *Peer),
		delPeer:      make(chan *Peer),
		msgCh:        make(chan *Message),
	}
}

func (s *Server) Start() {
	go s.loop()
	if err := s.listen(); err != nil {
		panic(err)
	}

	fmt.Printf("game server run on addr, %s\n", s.ServerConfig.ListenAddr)
	s.acceptLoop()
}

// Connect maybe construct new peer and handshake protocol after registration?
func (s *Server) Connect(addr string) error {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return err
	}
	peer := &Peer{conn: conn}
	s.addPeer <- peer

	return peer.Send([]byte(s.Version))
}

func (s *Server) acceptLoop() {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			panic(err)
		}

		peer := &Peer{
			conn: conn,
		}
		s.addPeer <- peer

		if err := peer.Send([]byte(s.Version)); err != nil {
			log.Fatal(err)
		}

		go s.handleConnection(peer)
	}
}

func (s *Server) handleConnection(peer *Peer) {
	buff := make([]byte, 1024)
	for {
		n, err := peer.conn.Read(buff)
		if err != nil {
			break
		}
		s.msgCh <- &Message{
			payload: buff[:n],
			From:    peer.conn.RemoteAddr(),
		}
		fmt.Printf("readed info: %s\n", string(buff[:n]))
	}

	s.delPeer <- peer
}

func (s *Server) listen() error {
	ln, err := net.Listen("tcp", s.ListenAddr)
	if err != nil {
		return err
	}

	s.listener = ln
	return nil
}

func (s *Server) loop() {
	for {
		select {
		case peer := <-s.delPeer:
			delete(s.peers, peer.conn.RemoteAddr())
			fmt.Println("player disconnected", peer.conn.LocalAddr(), peer.conn.RemoteAddr())
		case peer := <-s.addPeer:
			s.peers[peer.conn.RemoteAddr()] = peer
			fmt.Println("new player connected", peer.conn.LocalAddr(), peer.conn.RemoteAddr())
		case msg := <-s.msgCh:
			s.handler.HandleMessage(msg)
		}
	}
}
