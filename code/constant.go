package main

import "fmt"

func main() {
	const firstname string = "Joan"
	const lastname = "Lopo"

	fmt.Println(firstname)
	fmt.Println(lastname)

	const (
		nickname = "ama"
		surname  = "lopo"
	)

	fmt.Println(nickname)
	fmt.Println(surname)
}
