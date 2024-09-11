package pkg

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
	"time"
)

type Server struct {
	Members map[net.Addr]*Client
	Msgs    chan Msg
	Mu      sync.Mutex
}

func NewServer() *Server {
	return &Server{
		Members: make(map[net.Addr]*Client),
		Msgs:    make(chan Msg),
	}
}

func (s *Server) NewClient(conn net.Conn, file *os.File) {
	defer conn.Close()

	if len(s.Members) >= maxConnections {
		conn.Write([]byte("Chat is full, try again later\n"))
		return
	}

	c := &Client{
		Conn:     conn,
		Username: "",
		Msgs:     s.Msgs,
	}

	c.Conn.Write([]byte(welcomeMsg))

	s.HandleClient(c, file)

	delete(s.Members, c.Conn.RemoteAddr())

	s.Mu.Lock()
	close(c.Msgs)
	s.Mu.Unlock()
}

func (s *Server) HandleClient(c *Client, file *os.File) {
	c.Conn.Write([]byte("[ENTER YOUR NAME]: "))
	name, _ := bufio.NewReader(c.Conn).ReadString('\n')
	name = strings.TrimSpace(name)
	c.Username = name

	if c.Username != "" {
		s.Members[c.Conn.RemoteAddr()] = c

		c.WriteHistory()

		s.broadcast(c, fmt.Sprintf("%s has joined our chat...", c.Username))

		c.readInput()

		s.broadcast(c, fmt.Sprintf("%s has left our chat...", c.Username))
	} else {
		c.Conn.Write([]byte("enter a valid name!\n"))
		go s.HandleClient(c, file)
	}
}

func (s *Server) Run(file *os.File) {
	s.Mu.Lock()
	defer s.Mu.Unlock()
	for msg := range s.Msgs {
		s.msg(msg.Client, msg.Msg, file)
	}
}

func (s *Server) msg(c *Client, msg string, file *os.File) {
	msg = fmt.Sprintf("[%s][%s]: %s", time.Now().Format("2006-01-02 15:04:05"), c.Username, msg)
	s.broadcast(c, msg)
	_, err := file.Write([]byte(msg + "\n"))
	ErrorsHandler(err, false)
}

func (s *Server) broadcast(sender *Client, msg string) {
	for addr, m := range s.Members {
		if addr != sender.Conn.RemoteAddr() {
			m.msg(msg)
		}
	}
}
