package main

import "fmt"

func main() {
	// Tipe satu
	var name string

	name = "Eko Kurniawan"
	fmt.Println(name)

	name = "Eko Khannedy"
	fmt.Println(name)

	// Tipe dua
	var friendname = "Joanito agili lopo"
	fmt.Println(friendname)

	var age int8 = 30
	fmt.Println(age)

	// Tipe tiga
	lastname := "Agili Lopo"
	fmt.Println(lastname)

	// Tipe empat --> ultimeate
	var (
		firstname = "Joanito"
		endname   = "Agili"
	)
	fmt.Println(firstname)
	fmt.Println(endname)
}
