package main

import (
	"fmt"
	"strings"
)

type person struct {
	name string
	age  int
}

type student struct {
	grade  int
	person //embed
}

func main() {
	// Embedded struct adalah mekanisme untuk menempelkan sebuah struct sebagai properti struct lain

	s1 := student{}
	s1.name = "wick"
	s1.age = 21
	s1.grade = 2

	fmt.Println("name: ", s1.name) // wick
	fmt.Println("age: ", s1.age)   // 21
	s1.person.age = 22
	fmt.Println("age: ", s1.person.age) // 22
	fmt.Println("grade: ", s1.grade)    // 2

	fmt.Println(strings.Repeat("=", 20))

	// pengisian nilai sub-struct
	a1 := person{name: "joan", age: 21}
	b1 := student{person: a1, grade: 2}

	fmt.Println("name: ", b1.name)
	fmt.Println("age: ", b1.age)
	fmt.Println("grade; ", b1.grade)

}
