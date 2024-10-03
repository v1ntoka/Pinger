package Parser

import (
	"fmt"
	"log"
	"net"
)

func Less(ip, ip2 net.IP) bool {
	if len(ip) != len(ip2) {
		log.Fatalf("incorrect ip's to compare: %v and %v", ip, ip2)
	}
	for i := range ip {
		if ip[i] < ip2[i] {
			return true
		}
	}
	return false
}

func Increment(ip net.IP) (net.IP, error) {
	res := make(net.IP, len(ip))
	for i, n := range ip {
		res[i] = n
	}
	for i := len(res) - 1; i >= 0; i-- {
		if res[i] < 0xff {
			res[i] = res[i] + 0x01
			return res, nil
		} else {
			res[i] = 0
			continue
		}
	}
	return nil, fmt.Errorf("the end of ip range")
}
