package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Joanito",
		"address": "Soe",
	}

	person["title"] = "Programmer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	var book map[string]string = make(map[string]string)
	book["title"] = "Belajar Go-Lang"
	book["author"] = "Joan"
	book["Ups"] = "Salah"
	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "Ups")

	fmt.Println(book)
	fmt.Println(len(book))

}
