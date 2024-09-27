package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("Hello")
	testIP := net.IPAddr{
		net.IPv4(127, 0, 0, 1),
		"",
	}
}
