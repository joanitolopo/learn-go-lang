package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(10)

	fmt.Println("Random ke-1: ", rand.Int()) // 5221277731205826435
	fmt.Println("Random ke-2: ", rand.Int()) // 3852159813000522384
	fmt.Println("Random ke-3: ", rand.Int()) // 8532807521486154107

	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println("Random ke-1 using unique seed: ", rand.Int())
	fmt.Println("Random ke-2 using unique seed: ", rand.Int())
	fmt.Println("Random ke-3 using unique seed: ", rand.Int())

}
