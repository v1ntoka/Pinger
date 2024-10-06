package Ping

import (
	"net"
	"sync"
	"time"
)

type Pinger struct {
	Pool      []*stat
	Timeout   time.Duration
	PingEvery time.Duration
}

type stat struct {
	sync.Mutex
	IP       net.IP
	Index    int
	Sent     int64
	Received int64
	Percent  float32
	Err      error
}

func NewPinger(t time.Duration, pe time.Duration) *Pinger {
	return &Pinger{
		Pool:      make([]*stat, 0),
		Timeout:   t,
		PingEvery: pe,
	}
}

func (p *Pinger) AddIPs(addr []net.IP) {
	for i, ip := range addr {
		p.Pool = append(p.Pool, &stat{
			IP:    ip,
			Index: i,
		})
	}
}

func (p *Pinger) Run() {
	for _, s := range p.Pool {
		go Single(s, p.PingEvery, p.Timeout)
	}
	Declare(p)
}
