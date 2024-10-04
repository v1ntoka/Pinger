package Ping

import (
	"net"
	"sync"
	"time"
)

type Pinger struct {
	sync.RWMutex
	sync.WaitGroup
	Pool    []*stat
	Timeout time.Duration
}

type stat struct {
	sync.RWMutex
	IP       net.IP
	Index    int
	Timeout  time.Duration
	Sent     int64
	Received int64
	Percent  float32
	Err      error
}

func NewPinger(t time.Duration) *Pinger {
	return &Pinger{
		Pool:    make([]*stat, 0),
		Timeout: t,
	}
}

func (p *Pinger) AddIPs(addr []net.IP) {
	for i, ip := range addr {
		p.Pool = append(p.Pool, &stat{
			IP:       ip,
			Index:    i,
			Timeout:  p.Timeout,
			Sent:     0,
			Received: 0,
			Percent:  0,
			Err:      nil,
		})
	}
}

func (p *Pinger) Run() {
	for _, s := range p.Pool {
		go Single(s)
	}
	Declare(p)
}
