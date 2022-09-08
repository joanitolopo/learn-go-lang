package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {

	var secret any = &person{name: "wick", age: 27} // the value is a object referency not a real value
	var name = secret.(*person).name                // to get the real value, we need to cast to pointer

	fmt.Println(name)
}
