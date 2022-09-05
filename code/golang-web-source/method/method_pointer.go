package main

import "fmt"

type student struct {
	name  string
	grade int
}

func (s student) changeName1(name string) {
	fmt.Println("---> on changeName1, name changed to", name)
	s.name = name
}

func (s *student) changeName2(name string) {
	fmt.Println("---> on changeName2, name changed to", name)
	s.name = name
}

func main() {

	// Kelebihan method jenis ini adalah ketika kita melakukan manipulasi nilai pada property lain yang masih satu struct,
	// nilai pada property tersebut akan di rubah pada reference nya

	s1 := student{"john wick", 21}
	fmt.Println("s1 before", s1.name) // john wick

	s1.changeName1("jason bourne")
	fmt.Println("s1 after changeName1", s1.name) // john wick

	s1.changeName2("ethan hunt")
	fmt.Println("s1 after changeName2", s1.name) //ethan hunt

}
