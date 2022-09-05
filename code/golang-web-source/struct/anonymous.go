package main

import (
	"fmt"
	"strings"
)

type person struct {
	name string
	age  int
}

func main() {

	// anonyomous struct tanpa pengisian property
	s1 := struct {
		person //embedde struct
		grade  int
	}{} // anonymous struct

	s1.person = person{"wick", 21}
	s1.grade = 2

	fmt.Println("name: ", s1.person.name) // wick
	fmt.Println("age: ", s1.person.age)   // 21
	fmt.Println("grade: ", s1.grade)      // 2
	fmt.Println(strings.Repeat("=", 20))

	// anonymous struct dengan pengisian property
	var s2 = struct {
		person
		grade int
	}{
		person: person{"wick", 21},
		grade:  2,
	}

	fmt.Println("name2: ", s2.person.name) // wick
	fmt.Println("age2: ", s2.person.age)   // 21
	fmt.Println("grade2: ", s2.grade)      // 2

}
