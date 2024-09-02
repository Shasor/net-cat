package pkg

import (
	"bufio"
	"fmt"
	"log"
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

	fmt.Println(len(s.Members))

	if len(s.Members) >= maxConnections {
		conn.Write([]byte("Chat is full, try again later\n"))
		return
	}

	log.Printf("new client has connected: %s", conn.RemoteAddr().String())

	c := &Client{
		Conn:     conn,
		Username: "anonymous",
		Msgs:     s.Msgs,
	}

	s.Members[c.Conn.RemoteAddr()] = c

	c.Conn.Write([]byte(welcomeMsg))

	username, _ := bufio.NewReader(c.Conn).ReadString('\n')
	username = strings.Trim(username, "\r\n")
	c.Username = username

	if username != "anonymous" && username != "" {
		if history != "" {
			c.Conn.Write([]byte(history))
		}

		s.broadcast(c, fmt.Sprintf("%s has joined our chat...", c.Username))

		c.readInput()

		s.broadcast(c, fmt.Sprintf("%s has left our chat...", c.Username))
	} else {
		c.Conn.Write([]byte("entrez un pseudo valide\n"))
		delete(s.Members, c.Conn.RemoteAddr())
		return
	}

	delete(s.Members, c.Conn.RemoteAddr())

	s.Mu.Lock()
	close(c.Msgs)
	s.Mu.Unlock()
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
		if addr != sender.Conn.RemoteAddr() && m.Username != "anonymous" {
			m.msg(msg)
		}
	}
}
