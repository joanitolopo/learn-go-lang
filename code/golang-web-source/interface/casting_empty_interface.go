package main

import (
	"fmt"
	"strings"
)

func main() {

	var secret = any(2)
	var number = secret.(int) * 10 // because of secret tipe data is interface, you need cast it to int before times with 10
	fmt.Println(secret, "multiplied by 10 is :", number)

	secret = []string{"apple", "manggo", "banana"}
	var gruits = strings.Join(secret.([]string), ", ") // value use slice to declare, but secret variable is a interface, so we need to cast to slice type data before do the operation
	fmt.Println(gruits, "is my favorite fruits")

}
