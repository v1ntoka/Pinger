package Parser

import (
	"fmt"
	"log"
	"net"
)

func NewIP(n []byte) *net.IPAddr {
	res := &net.IPAddr{
		IP: make([]byte, 16),
	}
	for i, k := range n {
		res.IP[i] = k
	}
	return res
}

func Less(ip, ip2 *net.IPAddr) bool {
	if len(ip.IP) != len(ip2.IP) {
		log.Fatalf("incorrect ip's to compare: %v and %v", ip, ip2)
	}
	for i := range ip.IP {
		if ip.IP[i] < ip2.IP[i] {
			return true
		}
	}
	return false
}

func Equal(ip, ip2 *net.IPAddr) bool {
	if len(ip.IP) != len(ip2.IP) {
		return false
	}
	for i := range ip.IP {
		if ip.IP[i] != ip2.IP[i] {
			return false
		}
	}
	return true
}

func Increment(ip *net.IPAddr) (*net.IPAddr, error) {
	res := &net.IPAddr{
		IP: make([]byte, 16),
	}
	for i, n := range ip.IP {
		res.IP[i] = n
	}
	for i := len(res.IP) - 1; i >= 0; i-- {
		if res.IP[i] < 0xff {
			res.IP[i] = res.IP[i] + 0x01
			return res, nil
		} else {
			res.IP[i] = 0
			continue
		}
	}
	return nil, fmt.Errorf("the end of ip range")
}
