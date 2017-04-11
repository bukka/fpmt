package main

import (
	"fmt"
	"os"
	
	"github.com/bukka/fpmt/client"
)

func main() {
	fsClient := client.GetFlagSet()
	
	if len(os.Args) < 2 {
		fmt.Println("No argument set")
		os.Exit(1)
	}
	
	switch os.Args[1] {
		case "client":
			if err := fsClient.Parse(os.Args[2:]); err == nil {
				client.Run()
				os.Exit(0)
			}
			fmt.Println("Error when parsing client options")
			os.Exit(1)
		default:
			fmt.Println("Unknown command")
			os.Exit(1)
	}
}