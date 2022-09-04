package main

import (
	"fmt"
	"strings"
)

func main() {
	// Map adalah tipe data asosiatif yang berbentuk key-value pai

	// without default value
	chicken := map[string]int{}

	chicken["Januari"] = 50
	chicken["Februari"] = 40

	fmt.Println("Januari", chicken["Januari"]) // Januari 50
	fmt.Println("Mei", chicken["Mei"])         // Mei 0

	// with default value
	chicken1 := map[string]int{
		"Januari":  50,
		"Februari": 40}
	fmt.Println("Map with default value: ", chicken1) // map[Februari:40 Januari:50]

	// 3 ways to create a map
	// chicken3 := map[string]int{}
	// chicken4 := make(map[string]int)
	// chicken5 := *new(map[string]int)
	fmt.Println(strings.Repeat("=", 50))

	// iterasi map using for range
	chicken3 := map[string]int{
		"Januari":  50,
		"Februari": 40,
		"Maret":    34,
		"April":    67,
	}

	for key, val := range chicken3 {
		fmt.Println(key, " \t", val)
	}

	// delete item in map
	fmt.Println(strings.Repeat("=", 50))

	fmt.Println("Chicken3 before delete item: \n", chicken3) // map[April:67 Februari:40 Januari:50 Maret:34]
	delete(chicken3, "Januari")
	fmt.Println("Chicken3 after delete item: \n", chicken3) // map[April:67 Februari:40 Maret:34]

}
