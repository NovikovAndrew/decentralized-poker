package main

import (
	"fmt"
	"time"

	"decentralized-poker/cmd/internal/deck"
	"decentralized-poker/cmd/internal/p2p/server"
)

func main() {
	cgf := server.ServerConfig{
		ListenAddr: ":999",
		Version:    "Poker v0.0.1",
	}
	s := server.NewServer(cgf)
	go s.Start()
	// we start us connection and remote connect to us
	// :999 us tcp connection and start new tcp connection on :333
	// after connect :999 to :333
	// and we need sleep for goroutine we need time for dial connection
	time.Sleep(1 * time.Second)

	remoteCgf := server.ServerConfig{
		ListenAddr: ":333",
		Version:    "Poker v0.0.1",
	}

	// seems like remote node
	remoteServer := server.NewServer(remoteCgf)
	go remoteServer.Start()
	if err := remoteServer.Connect(":999"); err != nil {
		fmt.Println(err)
	}

	cards := deck.NewDeck()
	for _, card := range cards {
		fmt.Println(card)
	}

	select {}
}
