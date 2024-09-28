package main

import (
	"Pinger/pkg/Parser"
	"fmt"
)

func main() {
	rangeIP := "127.0.0.254 - 127.0.1.3"
	fmt.Println(Parser.Parse(rangeIP))
}
