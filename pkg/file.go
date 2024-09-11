package pkg

import (
	"bufio"
	"os"
)

func CreateFile() {
	file, _ = os.Create("log.txt")
}

func (c *Client) WriteHistory() {
	file, err := os.Open("log.txt")
	ErrorsHandler(err, false)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		c.Conn.Write([]byte(scanner.Text() + "\n"))
	}
}

func (c *Client) SaveHistory(msg string) {
	_, err := file.Write([]byte(msg + "\n"))
	ErrorsHandler(err, false)
}
