package main

import (
	"flag"
	"fmt"
	"net"
)

func main() {
	// parse host
	host := flag.String("host", "hex2dec.com", "the host name for dns lookup")
	flag.Parse()

	// dns lookup
	addrs, err := net.LookupHost(*host)
	if err != nil {
		fmt.Printf("Got error: %v\n", err)
		return
	}

	// print dns info
	fmt.Printf("The host: %s\n", *host)
	for _, addr := range addrs {
		fmt.Println(addr)
	}
}
