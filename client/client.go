package client

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/bukka/fpmt/client/fastcgi"
)

type Client struct {
	Host     string
	Port     string
	Script   string
	Body     string
	BodyType string
}

func (c *Client) String() string {
	return fmt.Sprintf("{host: %s, port: %s, script: '%s'}",
		c.Host, c.Port, c.Script)
}

func (c *Client) dial() (fcgi *fastcgi.FCGIClient, err error) {
	addr := net.JoinHostPort(c.Host, string(c.Port))

	return fastcgi.DialTimeout("tcp", addr, 0)
}

func (c *Client) log(fcgiParams map[string]string, response *http.Response) {
	fmt.Println("REQUEST:")
	fmt.Println("  FastCGI parameters:")
	for k, v := range fcgiParams {
		fmt.Printf("    %s: %v\n", k, v)
	}
	fmt.Println("RESPONSE:")
	fmt.Println("  StatusCode:", response.StatusCode)
	fmt.Println("  ContentLength:", response.ContentLength)
	fmt.Println("  Headers:")
	for k, v := range response.Header {
		fmt.Printf("    %s: %v\n", k, v)
	}
	fmt.Println("  Body:")
	fmt.Println("----------------")
	defer response.Body.Close()
	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err == nil {
		fmt.Println(string(bodyBytes))
	}
	fmt.Println("----------------")
}

func (c *Client) prepareRequest() (*fastcgi.FCGIClient, map[string]string, error) {
	fcgiParams := make(map[string]string)
	fcgiParams["SERVER_PROTOCOL"] = "HTTP/1.1"
	fcgiParams["SCRIPT_FILENAME"], _ = filepath.Abs(c.Script)

	// connect
	fcgi, err := c.dial()

	return fcgi, fcgiParams, err
}

func (c *Client) doGet() error {
	fcgi, fcgiParams, err := c.prepareRequest()
	if err != nil {
		return err
	}
	// send request
	response, err := fcgi.Get(fcgiParams)
	if err != nil {
		return err
	}
	// log response
	c.log(fcgiParams, response)

	return nil
}

func (c *Client) doPost(method string) error {
	fcgi, fcgiParams, err := c.prepareRequest()
	if err != nil {
		return err
	}

	// send request
	response, err := fcgi.Post(fcgiParams, method,
		c.BodyType, strings.NewReader(c.Body), int64(len(c.Body)))
	if err != nil {
		return err
	}
	// log response
	c.log(fcgiParams, response)

	return nil
}

// Run the client
func (c *Client) Run(action string) error {
	switch action {
	case "get":
		return c.doGet()
	case "post":
		return c.doPost("POST")
	case "put":
		return c.doPost("PUT")
	default:
		return fmt.Errorf("Unknown action %s", action)
	}
}
