package main

import (
	"Pnigger/pkg/Parser"
	"fmt"
)

func main() {
	rangeIP := "127.0.0.1 - 127.0.0.3"
	fmt.Println(Parser.Parse(rangeIP))
}
