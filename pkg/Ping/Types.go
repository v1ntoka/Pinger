package Ping

import (
	"fmt"
	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
	"net"
	"os"
	"sync"
	"time"
)

func ip2string(ip net.IPAddr) string {
	res := ""
	for _, c := range ip.String() {
		if c != ' ' && c != ':' {
			res += string(c)
		} else {
			return res
		}
	}
	return res
}

type Pinger struct {
	sync.RWMutex
	sync.WaitGroup
	pool map[string]*stat
}

type stat struct {
	Sent     int64
	Received int64
	Percent  float32
	Err      error
}

func NewPinger() *Pinger {
	return &Pinger{
		pool: make(map[string]*stat),
	}
}

func (p *Pinger) Run() {
	for s := range p.pool {
		go p.ping(s)
	}
	go func() {
		for {
			for ip, res := range p.pool {
				fmt.Println(ip, res.Sent, res.Received, res.Percent, res.Err)
			}
		}
	}()

}

func (p *Pinger) AddIPs(addr []net.IPAddr) {
	for _, ip := range addr {
		p.pool[ip2string(ip)] = &stat{}
	}
}

func (p *Pinger) ping(ip string) {
	defer p.Done()
	c, err := icmp.ListenPacket("ip4:icmp", "0.0.0.0")
	if err != nil {
		p.pool[ip].Err = err
		return
	}
	defer c.Close()
	for {

		wm := icmp.Message{
			Type: ipv4.ICMPTypeEcho, Code: 0,
			Body: &icmp.Echo{
				ID: os.Getpid() & 0xffff, Seq: 1,
				Data: make([]byte, 32),
			},
		}

		wb, err := wm.Marshal(nil)
		if err != nil {
			p.pool[ip].Err = err
			return
		}
		if _, err := c.WriteTo(wb, &net.IPAddr{IP: net.ParseIP(ip)}); err != nil {
			p.pool[ip].Err = err
			return
		}

		err = c.SetReadDeadline(time.Now().Add(time.Second * 1))
		if err != nil {
			fmt.Printf("Error on SetReadDeadline %v", err)
			panic(err)
		}

		rb := make([]byte, 1500)
		n, _, err := c.ReadFrom(rb)
		if err != nil {
			p.pool[ip].Err = err
			return
		}
		rm, err := icmp.ParseMessage(ipv4.ICMPTypeEchoReply.Protocol(), rb[:n])
		if err != nil {
			p.pool[ip].Err = err
			return
		}
		p.pool[ip].Sent++
		if rm.Type == ipv4.ICMPTypeEchoReply {
			p.pool[ip].Received++
		}
		p.pool[ip].Percent = float32(p.pool[ip].Sent) * 100 / float32(p.pool[ip].Received)
	}
}
