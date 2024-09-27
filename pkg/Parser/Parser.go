package Parser

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const (
	IPTemplate = `^(\d{1,3}\.){3}\d{1,3}$`
)

func Parse(s string) ([]IPAddress, error) {
	res := make([]IPAddress, 0)
	parts := strings.Split(s, "-")
	if len(parts) < 2 {
		return nil, fmt.Errorf("invalid IP range: %s", s)
	}
	ip, err := parseIP(parts[0])
	if err != nil {
		return nil, err
	}
	lastIP, err := parseIP(parts[len(parts)-1])
	if err != nil {
		return nil, err
	}
	if ip.Equal(lastIP) {
		res = append(res, *ip)
		return res, nil
	}
	for ip.Less(lastIP) {
		res = append(res, *ip)
		ip, err = ip.Increment()
		if err != nil {
			return nil, err
		}
	}
	res = append(res, *lastIP)
	return res, nil
}

//todo: В будущем планируется сделать возможность вводить как диапазон, так и отдельные адреса
//func parseRange(s string) *IPAddress {
//	return nil
//}

func parseIP(s string) (*IPAddress, error) {
	s = strings.TrimSpace(s)
	isCorrect, _ := regexp.MatchString(IPTemplate, s)
	if !isCorrect {
		return nil, fmt.Errorf("invalid IP address: %s", s)
	}
	parts := strings.Split(s, `.`)
	res := make([]byte, 4)
	for i, part := range parts {
		number, err := strconv.Atoi(part)
		if err != nil {
			return nil, err
		}
		if number > 255 {
			return nil, fmt.Errorf("invalid IP address: %s", s)
		}
		res[i] = byte(number)
	}
	return NewIP(res), nil
}
