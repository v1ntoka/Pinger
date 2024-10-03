package Ping

import (
	"golang.org/x/net/icmp"
	"golang.org/x/net/ipv4"
	"net"
	"os"
	"time"
)

func Single(s *stat) {
	c, err := icmp.ListenPacket("ip4:icmp", "0.0.0.0")
	defer func() {
		s.Err = err
	}()
	if err != nil {
		return
	}
	defer c.Close()

	for {
		s.Sent++
		wm := icmp.Message{
			Type: ipv4.ICMPTypeEcho, Code: 0,
			Body: &icmp.Echo{
				ID:   os.Getpid() & 0xffff,
				Data: make([]byte, 32),
			},
		}
		wb, err := wm.Marshal(nil)
		if err != nil {
			return
		}
		_, err = c.WriteTo(wb, &net.IPAddr{IP: s.IP})
		if err != nil {
			return
		}
		err = c.SetReadDeadline(time.Now().Add(s.Timeout))
		if err != nil {
			return
		}

		rb := make([]byte, 1500)
		n, _, err := c.ReadFrom(rb)
		if err != nil {
			return
		}

		rm, err := icmp.ParseMessage(ipv4.ICMPTypeEchoReply.Protocol(), rb[:n])
		if err != nil {
			return
		}
		if rm.Type == ipv4.ICMPTypeEchoReply {
			s.Received++
		}
		s.Percent = float32(s.Received) / float32(s.Sent) * 100
	}
}
