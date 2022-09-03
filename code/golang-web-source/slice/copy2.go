package main

import "fmt"

func main() {

	dst := []string{"Potato", "Potato", "Potato"}
	src := []string{"Watermelon", "Pinnaple"}
	n := copy(dst, src)

	fmt.Println("Value in dst after copy: ", dst)
	fmt.Println("Value in src: ", src)
	fmt.Println("n: ", n)
}
