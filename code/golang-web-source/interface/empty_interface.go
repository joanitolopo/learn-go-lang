package main

import (
	"fmt"
	"strings"
)

func main() {

	// interface kosong adalah tipe data yang spesial karena dapat menampung segala jenis data
	// tipe data ini biasanya disebut dengan dynamic typing

	var secret interface{}

	secret = "ethan hunt"
	fmt.Println(secret)

	secret = []string{"apple", "manggo", "banana"}
	fmt.Println(secret)

	secret = 12.4
	fmt.Println(secret)

	fmt.Println(strings.Repeat("=", 30))

	var data map[string]interface{}

	data = map[string]interface{}{
		"name":      "ethan hunt",
		"grade":     2,
		"breakfast": []string{"apple", "manggo", "banana"},
	}

	fmt.Println(data)

	fmt.Println(strings.Repeat("=", 50))

	var data2 map[string]any

	data2 = map[string]any{
		"nama": "reynal",
		"pendidikan": map[string]string{
			"SD":  "SD NEGERI 6",
			"SMP": "SMP NEGERI 7",
			"SMA": "SMA NEGERI 8",
		},
	}

	fmt.Println(data2)

	fmt.Println(strings.Repeat("=", 50))

}
