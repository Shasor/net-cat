package pkg

import (
	"bufio"
	"os"
)

func CreateFile() *os.File {
	file, err := os.Create("log.txt")
	ErrorsHandler(err, false)
	return file
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
