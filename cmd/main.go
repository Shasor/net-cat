package main

import (
	"fmt"
	"net"
	"net-cat/pkg"
	"os"
)

func main() {
	if len(os.Args) > 2 {
		fmt.Println("[USAGE]: ./TCPChat $port")
		return
	}

	pkg.CreateFile()

	port := pkg.DefaultPort
	if len(os.Args) == 2 {
		port = os.Args[1]
	}

	s := pkg.NewServer()
	go s.Run()

	listener, err := net.Listen("tcp", ":"+port)
	pkg.ErrorsHandler(err, true)
	defer listener.Close()
	fmt.Printf("Listening on the port " + ":" + port + "\n")

	for {
		conn, err := listener.Accept()
		pkg.ErrorsHandler(err, false)
		go s.NewClient(conn)
	}
}
