package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	allStudents := []person{
		{name: "Wick", age: 23},
		{name: "Ethan", age: 23},
		{name: "Bourne", age: 23},
	}

	for _, student := range allStudents {
		fmt.Println(student.name, "age is", student.age)
	}

	// output:
	// Wick age is 23
	// Ethan age is 23
	// Bourne age is 23

}
