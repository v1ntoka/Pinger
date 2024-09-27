package Parser

import "strings"

const (
	IPTemplate = `^(\d{1,3}\.){3}\d{1,3}$`
)

func Parse(input string) []*IPAddress {
	parts := make([]string, 2)
	if strings.Contains(input, "-") {

	}

}

func parseRange(input string) *IPAddress {

}
