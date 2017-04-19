package client

import (
	"errors"
	"fmt"
	"net"

	"github.com/bukka/fpmt/client/fastcgi"
)

type Client struct {
	Host   string
	Port   string
	Script string
}

func (c *Client) String() string {
	return fmt.Sprintf("{host: %s, port: %d, script: '%s'}",
		c.Host, c.Port, c.Script)
}

func (c *Client) dial() (fcgi *fastcgi.FCGIClient, err error) {
	return fastcgi.DialTimeout("tcp", net.JoinHostPort(c.Host, string(c.Port)), 0)
}

func (c *Client) doGet() error {
	fcgiParams := make(map[string]string)
	fcgiParams["REQUEST_METHOD"] = "GET"
	fcgiParams["SERVER_PROTOCOL"] = "HTTP/1.1"
	fcgiParams["SCRIPT_FILENAME"] = c.Script

	// connect
	fcgi, err := c.dial()
	if err != nil {
		return err
	}
	// send request
	resp, err := fcgi.Get(fcgiParams)
	if err != nil {
		return err
	}

	fmt.Println(resp)

	return nil
}

func (c *Client) Run(action string) error {
	switch action {
	case "get":
		return c.doGet()
	default:
		return errors.New(fmt.Sprintf("Unknown action %s", action))
	}
}
