package main

import (
	"Pnigger/pkg/Parser"
	"fmt"
)

func main() {
	test, _ := Parser.NewIP(127, 0, 0, 1)
	fmt.Println(test)
	test2, _ := Parser.NewIP(127, 0, 0, 2)
	fmt.Println(test2.Less(test))
}
