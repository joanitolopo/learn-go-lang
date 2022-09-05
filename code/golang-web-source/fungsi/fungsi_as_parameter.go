package main

import (
	"fmt"
	"strings"
)

func main() {
	data := []string{"wick", "jason", "ethan"}
	dataContainsO := filter(data, func(each string) bool {
		return strings.Contains(each, "o")
	})
	dataLenght5 := filter(data, func(each string) bool {
		return len(each) == 5
	})

	fmt.Println("data asli \t\t:", data)                     // [wick jason ethan]
	fmt.Println("filter ada huruf \"o\"\t:", dataContainsO)  // [jason]
	fmt.Println("filter jumlah huruf \"5\"\t:", dataLenght5) // [jason ethan]
}

type FilterCallback func(string) bool

func filter(data []string, callback FilterCallback) []string {
	var result []string

	for _, each := range data {
		if filtered := callback(each); filtered {
			result = append(result, each)
		}
	}
	return result

}
