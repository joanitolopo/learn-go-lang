package main

import (
	"fmt"
	"strings"
)

func main() {
	// 3 index adalah teknik slicing elemen yang sekaligus menentukan kapasitasnya
	var fruits = []string{"apple", "grape", "banana"}
	var aFruits = fruits[0:2]
	var bFruits = fruits[0:2:2]

	fmt.Println("Values in fruits slice: ", fruits)
	fmt.Println("Total values (len) in fruits: ", len(fruits))
	fmt.Println("Max capacites (cap) in fruits: ", cap(fruits))

	fmt.Println(strings.Repeat("=", 40))

	fmt.Println("Values in aFruits slice: ", aFruits)
	fmt.Println("Total values (len) in aFruits: ", len(aFruits))
	fmt.Println("Max capacites (cap) in aFruits: ", cap(aFruits))

	fmt.Println(strings.Repeat("=", 40))

	fmt.Println("Values in fruits bFruits: ", bFruits)
	fmt.Println("Total values (len) in bFruits: ", len(bFruits))
	fmt.Println("Max capacites (cap) in bFruits: ", cap(bFruits))

	fmt.Println(strings.Repeat("=", 40))

}
