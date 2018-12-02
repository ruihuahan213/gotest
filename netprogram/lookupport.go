package main

import (
	"net"
	"os"
	"fmt"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "usage: %s network-type service-name",
								os.Args[0])
		os.Exit(1)
	}

	networkType := os.Args[1]
	service := os.Args[2]

	port, err := net.LookupPort(networkType, service)

	if err != nil {
		fmt.Println("Error: ", err.Error())
		os.Exit(1)
	}

	fmt.Println("servcie port: ", port)
	os.Exit(0)
}