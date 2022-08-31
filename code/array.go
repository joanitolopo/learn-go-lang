package main

import "fmt"

func main() {

	// sintax ribet
	var names [3]string

	names[0] = "Joanito"
	names[1] = "Agili"
	names[2] = "Lopo"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	// membuat array langsung
	var values = [3]int{
		90,
		80,
		95,
	}
	fmt.Println(values)
	fmt.Println(values[0], values[1], values[2])

	// array functions
	fmt.Println(len(values))
	fmt.Println(len(names))

	// ingat len bukan banyaknya data, tetapi panjangnya array
	var lagi [10]string
	fmt.Println(len(lagi))

}
