package main

import "fmt"

func main() {
	// fungsi copy digunakan untuk memcopy elemen slice pada src ke dst --> copy(dst, src)
	// Jumlah elemen yang dicopy dari src adalah sejumlah lebar yang dipersiapkan oleh src atau dikatakan sebagai len(dst)
	dst := make([]string, 3)
	src := []string{"watermelon", "pinnaple", "apple", "orange"}
	n := copy(dst, src)

	fmt.Println("Value in dst slice after copy: ", dst)
	fmt.Println("Value src: ", src)
	fmt.Println("n: ", n)
}
