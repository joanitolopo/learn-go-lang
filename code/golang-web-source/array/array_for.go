package main

import (
	"fmt"
	"strings"
)

func main() {

	fruits := [4]string{"apple", "grape", "banana", "melon"}

	// Using for
	for i := 0; i < len(fruits); i++ {
		fmt.Printf("elemen %d : %s\n", i, fruits[i])
	}

	fmt.Println(strings.Repeat("=", 20))

	// using for range
	for i, fruit := range fruits {
		fmt.Printf("elemen %d : %s\n", i, fruit)
	}

}
