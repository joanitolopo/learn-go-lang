package main

import (
	"fmt"
	"strings"
)

func main() {
	var names [4]string

	names[0] = "trafagar"
	names[1] = "d"
	names[2] = "water"
	names[3] = "law"

	fmt.Println(names[0], names[1], names[2], names[3])
	fmt.Println(strings.Repeat("=", 20))
}
