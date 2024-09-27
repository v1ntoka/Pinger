package Parser

import "strconv"

type IPAddress [4]byte

func NewIP(n1, n2, n3, n4 byte) (*IPAddress, error) {
	return &IPAddress{n1, n2, n3, n4}, nil
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
