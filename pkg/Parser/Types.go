package Parser

import (
	"fmt"
	"strconv"
)

type IPAddress [4]byte

func NewIP(n []byte) *IPAddress {
	res := &IPAddress{}
	for i, k := range n {
		res[i] = k
	}
	return res
}

func (ip *IPAddress) Less(ip2 *IPAddress) bool {
	for i := range ip {
		if ip[i] < ip2[i] {
			return true
		}
	}
	return false
}

func (ip *IPAddress) Equal(ip2 *IPAddress) bool {
	for i := range ip {
		if ip[i] != ip2[i] {
			return false
		}
	}
	return true
}

func (ip *IPAddress) String() string {
	res := ""
	for _, c := range ip {
		res += strconv.Itoa(int(c)) + "."
	}
	return res[:len(res)-1]
}

func (ip *IPAddress) Add(n int) (*IPAddress, error) {
	res := &IPAddress{}
	return res, nil
}

func (ip *IPAddress) Increment() (*IPAddress, error) {
	res := &IPAddress{}
	for i, n := range ip {
		res[i] = n
	}
	for i := 3; i >= 0; i-- {
		if res[i] < 0xff {
			res[i] = res[i] + 1
			return res, nil
		}
	}
	return nil, fmt.Errorf("end of ip range")
}
