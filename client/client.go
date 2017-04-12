package client

import (
	"fmt"
)

type Client struct {
	Host   string
	Port   uint
	Script string
}

func (c *Client) String() string {
	return fmt.Sprintf("{host: %s, port: %d, script: '%s'}",
		c.Host, c.Port, c.Script)
}

func (c *Client) Run() {
	fmt.Println("run client:", c)
}
