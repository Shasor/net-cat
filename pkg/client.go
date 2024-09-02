package pkg

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
)

type Msg struct {
	Client *Client
	Msg    string
}

type Client struct {
	Conn     net.Conn
	Username string
	Msgs     chan<- Msg
}

func (c *Client) readInput() {
	for {
		c.Conn.Write([]byte(fmt.Sprintf("[%s][%s]: ", time.Now().Format("2006-01-02 15:04:05"), c.Username)))
		input, err := bufio.NewReader(c.Conn).ReadString('\n')
		if err != nil {
			return
		}

		input = strings.TrimSpace(input)

		if input != "" {
			c.Msgs <- Msg{
				Client: c,
				Msg:    input,
			}
		}
	}
}

func (c *Client) msg(msg string) {
	c.Conn.Write([]byte(fmt.Sprintf("\n%s\n[%s][%s]: ", msg, time.Now().Format("2006-01-02 15:04:05"), c.Username)))
}
