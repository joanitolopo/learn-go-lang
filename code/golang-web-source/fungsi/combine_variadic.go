package main

import (
	"fmt"
	"strings"
)

func main() {
	hobbies := []string{"sleeping", "eating"}
	yourHobbies("wick", hobbies...)
}

func yourHobbies(name string, hobbies ...string) {
	// Parameter variadic bisa dikombinasikan dengan parameter biasa, dengan syarat parameter variadic-nya harus diposisikan di akhir
	hobbiesAsString := strings.Join(hobbies, ", ")

	fmt.Printf("Hello, my name is: %s\n", name)
	fmt.Printf("My Hobbies are: %s\n", hobbiesAsString)

}
