package main

import "fmt"

func main() {

	chicken := (map[string]int{
		"Januari":  50,
		"Februari": 40,
	})

	var value, isExist = chicken["Mei"]

	fmt.Println("the value: ", value)          // 0
	fmt.Println("the isExist value:", isExist) // false

	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("Items is not exists")
	}

}
