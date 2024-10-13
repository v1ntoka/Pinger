package main

import (
	"Pinger/pkg/Parser"
	"Pinger/pkg/Ping"
	"bufio"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"time"
)

func main() {
	go func() {
		http.ListenAndServe("localhost:8080", nil)
	}()
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)
	go func() {
		for _ = range c {
			fmt.Println("\r- Ctrl+C pressed in Terminal")
			os.Exit(0)
		}
	}()
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
