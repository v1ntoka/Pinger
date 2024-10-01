package main

import (
	"Pinger/pkg/Parser"
	"Pinger/pkg/Ping"
	"log"
)

func main() {
	rng := "8.8.8.8 - 8.8.8.8"
	ips, err := Parser.Parse(rng)
	if err != nil {
		log.Fatal(err)
	}
	p := Ping.NewPinger()
	p.AddIPs(ips)
	p.Run()
}
