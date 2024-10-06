package main

import (
	"Pinger/pkg/Parser"
	"Pinger/pkg/Ping"
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	fmt.Println("Enter the IP range in format 0.0.0.0-0.0.0.1")
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	ips, err := Parser.Parse(s.Text())
	if err != nil {
		log.Fatal(err)
	}
	p := Ping.NewPinger(500*time.Millisecond, time.Second)
	p.AddIPs(ips)
	p.Run()
}
