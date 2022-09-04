package main

import (
	"fmt"
	"strings"
)

func main() {

	chickens := []map[string]string{
		{"name": "chicken blue", "gender": "male"},
		{"name": "chicken red", "gender": "male"},
		{"name": "chicken yellow", "gender": "female"},
	}

	for _, chicken := range chickens {
		fmt.Println(chicken["gender"], chicken["name"])
	}

	// multi keys
	data := []map[string]string{
		{"name": "chicken blue", "gender": "male", "color": "brown"},
		{"addres": "mangga stree", "id": "k001"},
		{"community": "chicken lovers"},
	}
	fmt.Println(strings.Repeat("=", 20))
	for _, datas := range data {
		fmt.Println(datas["addres"], datas["community"])
	}

}
