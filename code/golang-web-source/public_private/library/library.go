package library

import "fmt"

type Mahasiswa struct {
	Name  string
	Grade int
}

var Student = struct {
	Name  string
	Grade int
}{}

func init() {
	Student.Name = "John wick"
	Student.Grade = 2

	fmt.Println("---> library/library.go imported")

}
