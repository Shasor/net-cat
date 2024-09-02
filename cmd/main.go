package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"pkg"
)

func newServer() *pkg.Server {
	return &pkg.Server{
		Members: make(map[net.Addr]*pkg.Client),
		Msgs:    make(chan pkg.Msg),
	}
}

func main() {
	if len(os.Args) > 2 {
		fmt.Println("[USAGE]: ./TCPChat $port")
		return
	}

	port := pkg.DefaultPort
	if len(os.Args) == 2 {
		port = os.Args[1]
	}

	s := newServer()
	go s.Run()

	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("unable to start server: %s", err.Error())
	}
	defer listener.Close()
	fmt.Printf("Listening on the port " + ":" + port + "\n")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("unable to accept connection: %s", err.Error())
			continue
		}
		go s.NewClient(conn)
	}
}
