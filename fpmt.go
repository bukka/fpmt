package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/bukka/fpmt/client"
	"github.com/bukka/fpmt/server"
)

func clientFlagSet() (*flag.FlagSet, *client.Client) {
	c := &client.Client{}
	fsClient := flag.NewFlagSet("client", flag.ContinueOnError)
	fsClient.StringVar(&c.Host, "h", "127.0.0.1", "Server host")
	fsClient.StringVar(&c.Port, "p", "9012", "Server port")
	fsClient.StringVar(&c.Script, "s", "", "Script name")
	fsClient.StringVar(&c.Body, "b", "", "Body")
	fsClient.StringVar(&c.BodyType, "t", "", "Body type")
	fsClient.StringVar(&c.Uri, "u", "", "Request URI")

	return fsClient, c
}

func serverFlagSet() (*flag.FlagSet, *server.Server) {
	s := &server.Server{}
	fsServer := flag.NewFlagSet("server", flag.ContinueOnError)
	fsServer.StringVar(&s.FpmBinary, "f", "/usr/local/sbin/php-fpm", "FPM binary")
	fsServer.StringVar(&s.FpmConfig, "c", "/usr/local/etc/php-fpm.conf'", "FPM config")
	fsServer.StringVar(&s.FpmConfigTemplate, "t", "", "FPM config template")

	return fsServer, s
}

func main() {
	fsClient, c := clientFlagSet()
	fsServer, s := serverFlagSet()

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
	case "server":
		if err := fsServer.Parse(os.Args[3:]); err != nil {
			fmt.Println("Error when parsing server options")
			os.Exit(1)
		}

		if err := s.Run(action); err != nil {
			fmt.Println("Error ", err)
			os.Exit(1)
		}
		os.Exit(0)
	default:
		fmt.Println("Unknown command")
		os.Exit(1)
	}
}
