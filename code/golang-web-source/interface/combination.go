package main

import (
	"fmt"
	"strings"
)

func main() {

	var person = []map[string]interface{}{
		{"name": "wick", "age": 23},
		{"name": "ethan", "age": 23},
		{"name": "bourne", "age": 22},
	}

	for _, each := range person { // looping the slice (each value)
		fmt.Println(each)
		fmt.Println(each["name"], "age is", each["age"])
	}

	fmt.Println(strings.Repeat("=", 30))

	var fruits = []interface{}{
		map[string]interface{}{
			"name":  "strawbery",
			"total": 10,
		},
		[]string{"manggo", "pineapple", "papay"},
		"orange",
	}

	for index, each := range fruits { // kenapa bisa diloop padahal sebuah interface, dikarenakan ditambahkan slice [] sehingga bisa di loop
		fmt.Println(index, each)

		if index == 1 {
			fmt.Println("ini di dalam loop slice :")
			for _, value := range each.([]string) {
				fmt.Println(value)

			}
		}
	}
}
