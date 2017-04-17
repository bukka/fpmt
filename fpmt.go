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
	fsClient.StringVar(&c.Host, "h", "127.0.0.1", "Server host")
	fsClient.UintVar(&c.Port, "p", 9800, "Server port")
	fsClient.StringVar(&c.Script, "s", "", "Script name")

	return fsClient, c
}

func main() {
	fsClient, c := clientFlagSet()

	if len(os.Args) < 2 {
		fmt.Println("No argument set")
		os.Exit(1)
	}
	if len(os.Args) == 2 {
		fmt.Println("No action set")
		os.Exit(1)
	}
	action := os.Args[2]

	switch os.Args[1] {
	case "client":
		if err := fsClient.Parse(os.Args[3:]); err == nil {
			c.Run(action)
			os.Exit(0)
		}
		fmt.Println("Error when parsing client options")
		os.Exit(1)
	default:
		fmt.Println("Unknown command")
		os.Exit(1)
	}
}
