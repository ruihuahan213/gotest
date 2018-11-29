package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	if len(os.Args) !=2 {
		fmt.Fprintf(os.Stderr, "Usage: %s dotted-ip-addr\n", os.Args[0])
		os.Exit(1)
	}

	var dotAddr string = os.Args[1]

	var  addr net.IP = net.ParseIP(dotAddr)

	if addr == nil {
		fmt.Printf("Invalid address")
		os.Exit(1)
	}

	var mask net.IPMask = addr.DefaultMask()
	var network net.IP = addr.Mask(mask)
	var ones, bits int = mask.Size()

	fmt.Println(
		"Address is ", addr.String(),
		" Default mas length is ", bits,
		" Leading ones count is ", ones,
		" Mask is (hex) ", mask.String(),
		" Network is ", network.String())
	
}