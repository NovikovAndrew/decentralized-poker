package main

import "decentralized-poker/cmd/internal/p2p/server"

func main() {
	//cards := deck.NewDeck()
	//for _, c := range cards {
	//	fmt.Println(c)
	//}

	cgf := server.ServerConfig{
		ListenAddr: ":999",
		Version:    "Poker v0.0.1",
	}
	s := server.NewServer(cgf)
	go s.Start()

	remoteCgf := server.ServerConfig{
		ListenAddr: ":333",
		Version:    "Poker v0.0.1",
	}
	remoteServer := server.NewServer(remoteCgf)
	remoteServer.Start()
}
