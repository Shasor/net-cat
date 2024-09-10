package pkg

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"sync"
	"time"
)

type Server struct {
	Members map[net.Addr]*Client
	Msgs    chan Msg
	Mu      sync.Mutex
}

func (s *Server) NewClient(conn net.Conn) {
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

	s.HandleClient(c)

	delete(s.Members, c.Conn.RemoteAddr())

	s.Mu.Lock()
	close(c.Msgs)
	s.Mu.Unlock()
}

func (s *Server) HandleClient(c *Client) {
	c.Conn.Write([]byte("[ENTER YOUR NAME]: "))
	username, _ := bufio.NewReader(c.Conn).ReadString('\n')
	username = strings.TrimSpace(username)
	c.Username = username

	if c.Username != "" {
		s.Members[c.Conn.RemoteAddr()] = c

		if history != "" {
			c.Conn.Write([]byte(history))
		}

		s.broadcast(c, fmt.Sprintf("%s has joined our chat...", c.Username))

		c.readInput()

		s.broadcast(c, fmt.Sprintf("%s has left our chat...", c.Username))
	} else {
		c.Conn.Write([]byte("enter a valid username!\n"))
		go s.HandleClient(c)
	}
}

func (s *Server) Run() {
	s.Mu.Lock()
	defer s.Mu.Unlock()
	for msg := range s.Msgs {
		s.msg(msg.Client, msg.Msg)
	}
}

func (s *Server) msg(c *Client, msg string) {
	msg = fmt.Sprintf("[%s][%s]: %s", time.Now().Format("2006-01-02 15:04:05"), c.Username, msg)
	s.broadcast(c, msg)
	history += msg + "\n"
}

func (s *Server) broadcast(sender *Client, msg string) {
	for addr, m := range s.Members {
		if addr != sender.Conn.RemoteAddr() {
			m.msg(msg)
		}
	}
}
