package main

import (
	f "fmt" // we make alias for fmt package
	"strings"

	// . "public_private/library" if we use dot, then we dont need to specify the package name
	"public_private/library"
)

func main() {
	var s1 = library.Mahasiswa{"ethan", 21}
	f.Println("name ", s1.Name)
	f.Println("grade ", s1.Grade)

	f.Println(strings.Repeat("=", 15))

	// sayHello("ethan") if we want to execute here we need to call all of the function --> go run *.go

	f.Printf("Name : %s\n", library.Student.Name)
	f.Printf("Grade : %d\n", library.Student.Grade)
}
