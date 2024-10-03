package Parser

import (
	"fmt"
	"net"
	"strings"
)

const (
	IPTemplate = `^(\d{1,3}\.){3}\d{1,3}$`
)

func Parse(s string) (res []net.IP, err error) {
	res = make([]net.IP, 0)
	parts := strings.Split(s, "-")
	if len(parts) < 2 {
		return nil, fmt.Errorf("invalid IP range: %s", s)
	}
	ip := parseIP(parts[0])
	lastIP := parseIP(parts[len(parts)-1])
	if ip.Equal(lastIP) {
		res = append(res, ip)
		return res, nil
	}
	for Less(ip, lastIP) {
		res = append(res, ip)
		ip, err = Increment(ip)
		if err != nil {
			return nil, err
		}
	}
	res = append(res, lastIP)
	return res, nil
}

//todo: В будущем планируется сделать возможность вводить как диапазон, так и отдельные адреса
//func parseRange(s string) *net.IPAddr {
//	return nil
//}

func parseIP(s string) net.IP {
	s = strings.TrimSpace(s)
	res := net.ParseIP(s)
	return res.To4()
}
