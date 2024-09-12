package pkg

type commandID int

const (
	CMD_NAME commandID = iota
)

type Command struct {
	Id     commandID
	Client *Client
	Args   []string
}
