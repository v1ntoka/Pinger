package Parser

import "fmt"

type IPAddress [4]byte

func New(n1, n2, n3, n4 byte) (*IPAddress, error) {
	if n1 > 255 || n2 > 255 || n3 > 255 || n4 > 255 {
		return nil, fmt.Errorf("invalid IP address: values cannot be greater than 255")
	}
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
	return fmt.Sprintf("%v.%v.%v.%v", ip)
}
