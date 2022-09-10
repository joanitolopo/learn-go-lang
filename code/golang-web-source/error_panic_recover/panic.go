package main

import (
	"errors"
	"fmt"
	"strings"
)

func validate(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("Cannot be empty")
	}
	return true, nil

}

func main() {
	// sama seperi erro tetapi informasi errornya lebih mendetail dan heboh :)

	defer fmt.Println("end")

	var name string

	fmt.Print("Type your name: ")
	fmt.Scanln(&name)

	if valid, err := validate(name); valid {
		fmt.Println("halo", name)
	} else {
		panic(err.Error())
	}

}
