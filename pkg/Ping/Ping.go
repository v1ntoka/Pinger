package Ping

import (
	"errors"
	"fmt"
	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
	"net"
	"os"
	"time"
)

func Single(s *stat, pe time.Duration, to time.Duration) {
	c, err := icmp.ListenPacket("ip4:icmp", fmt.Sprintf("0.0.0.0"))
	defer c.Close()
	if err != nil {
		s.Err = err
		return
	}
	for {
		time.Sleep(pe)
		s.Sent++
		wm := icmp.Message{
			Type: ipv4.ICMPTypeEcho, Code: 0,
			Body: &icmp.Echo{
				ID:   os.Getpid() & 0xffff,
				Seq:  s.Index,
				Data: make([]byte, 32),
			},
		}
		wb, err := wm.Marshal(nil)
		if err != nil {
			s.Err = err
			return
		}
		_, err = c.WriteTo(wb, &net.IPAddr{IP: s.IP})
		if err != nil {
			s.Err = err
			return
		}
		err = c.SetReadDeadline(time.Now().Add(to))
		if err != nil {
			s.Err = err
			return
		}
		rb := make([]byte, 1500)
		n, _, err := c.ReadFrom(rb)
		if err != nil {
			if errors.Is(err, os.ErrDeadlineExceeded) {
				continue
			} else {
				s.Err = err
				return
			}
		}
		rm, err := icmp.ParseMessage(ipv4.ICMPTypeEchoReply.Protocol(), rb[:n])
		if err != nil {
			s.Err = err
			return
		}
		if rm.Type == ipv4.ICMPTypeEchoReply && rm.Body.(*icmp.Echo).Seq == s.Index {
			s.Received++
		}
		s.Percent = float32(s.Received) / float32(s.Sent) * 100
	}
}
