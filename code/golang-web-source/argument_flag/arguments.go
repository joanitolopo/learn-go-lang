package main

import (
	"fmt"
	"os"
)

func main() {
	var argsRaw = os.Args
	fmt.Printf("-> %#v\n", argsRaw)
	// -> []string{"/tmp/go-build3082151662/b001/exe/arguments", "banana", "potato", "ice cream"}

	var args = argsRaw[1:]
	fmt.Printf("-> %#v\n", args)
	// -> []string{"banana", "potato", "ice cream"}
}
