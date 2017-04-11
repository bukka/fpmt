package client

import (
	"flag"
	"fmt"
)

var host string
var port uint
var script string

func GetFlagSet() *flag.FlagSet {
	fsClient := flag.NewFlagSet("client", flag.ContinueOnError)
	fsClient.StringVar(&host, "host", "127.0.0.1", "Server host")
	fsClient.UintVar(&port, "port", 9800, "Server port")
	fsClient.StringVar(&script, "script", "", "Script name")

	return fsClient
}

func Run() {
	fmt.Println("run client")
}
