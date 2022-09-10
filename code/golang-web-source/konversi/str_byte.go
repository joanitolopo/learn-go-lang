package main

import "fmt"

func main() {
	// sebuah string di go adalah sebuah lice/array byte
	// cara mendapatkan slice byte dari sebuah data string adalah dengan cast ke tipe []byte

	// data bertipe []byte dicari bentuk stringnya
	var byte1 = []byte{104, 97, 108, 111}
	var s = string(byte1)
	fmt.Printf("%s \n", s) // halo
}
