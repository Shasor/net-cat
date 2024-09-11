package main

import (
	"fmt"
	"net"
	"os"
	"pkg"
)

func main() {
	if len(os.Args) > 2 {
		fmt.Println("[USAGE]: ./TCPChat $port")
		return
	}

	file := pkg.CreateFile()

	port := pkg.DefaultPort
	if len(os.Args) == 2 {
		port = os.Args[1]
	}

	s := pkg.NewServer()
	go s.Run(file)

	listener, err := net.Listen("tcp", ":"+port)
	pkg.ErrorsHandler(err, true)
	defer listener.Close()
	fmt.Printf("Listening on the port " + ":" + port + "\n")

	for {
		conn, err := listener.Accept()
		pkg.ErrorsHandler(err, false)
		go s.NewClient(conn, file)
	}
}
