package main

import (
	"fmt"
	"strings"
)

type student struct {
	name  string
	grade int
}

func main() {
	var s1 = student{"john wick", 21}
	s1.sayHello() // halo

	var name = s1.getNameAt(2)
	fmt.Println("nama panggilan: ", name) // wick

}

func (s student) sayHello() { // maksudnya adalah fungsi sayHello dideklarasikan sebagaia method milik struct student

	fmt.Println("halo")
}

func (s student) getNameAt(i int) string {
	return strings.Split(s.name, " ")[i-1]

}
