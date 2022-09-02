package main

import (
	"fmt"
	"strings"
)

func main() {
	var names [4]string

	names[0] = "trafagar"
	names[1] = "d"
	names[2] = "water"
	names[3] = "law"

	fmt.Println(names[0], names[1], names[2], names[3])
	fmt.Println(strings.Repeat("=", 20))

	var fruits = [4]string{"apple", "grape", "banana", "melon"}
	fmt.Println("Jumlah element \t\t", len(fruits))
	fmt.Println("Isi semua element \t", fruits)
	fmt.Println(strings.Repeat("=", 50))

	var numbers = [...]int{2, 3, 2, 4, 3}

	fmt.Println("data array \t:", numbers)
	fmt.Println("jumlah elemen \t:", len(numbers))
	fmt.Println(strings.Repeat("=", 30))

	var number1 = [2][3]int{{3, 2, 3}, {3, 4, 5}}
	fmt.Printf("Array 1: %d --> Array 2: %d \n", number1[0], number1[1])

}
