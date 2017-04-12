package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/bukka/fpmt/client"
)

func clientFlagSet() (*flag.FlagSet, *client.Client) {
	c := &client.Client{}
	fsClient := flag.NewFlagSet("client", flag.ContinueOnError)
	fsClient.StringVar(&c.Host, "host", "127.0.0.1", "Server host")
	fsClient.UintVar(&c.Port, "port", 9800, "Server port")
	fsClient.StringVar(&c.Script, "script", "", "Script name")

	return fsClient, c
}

func main() {
	fsClient, c := clientFlagSet()

	if len(os.Args) < 2 {
		fmt.Println("No argument set")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "client":
		if err := fsClient.Parse(os.Args[2:]); err == nil {
			c.Run()
			os.Exit(0)
		}
		fmt.Println("Error when parsing client options")
		os.Exit(1)
	default:
		fmt.Println("Unknown command")
		os.Exit(1)
	}
}
