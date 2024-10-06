package main

import (
	"Pinger/pkg/Parser"
	"Pinger/pkg/Ping"
	"bufio"
	"log"
	"os"
	"time"
)

func main() {
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
