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
	fsClient.StringVar(&c.Port, "p", "9012", "Server port")
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
		if err := fsClient.Parse(os.Args[3:]); err != nil {
			fmt.Println("Error when parsing client options")
			os.Exit(1)
		}

		if err := c.Run(action); err != nil {
			fmt.Println("Error ", err)
			os.Exit(1)
		}
		os.Exit(0)
	default:
		fmt.Println("Unknown command")
		os.Exit(1)
	}
}
